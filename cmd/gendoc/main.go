package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	apipb "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	authzv1 "github.com/terrabase-dev/terrabase/specs/terrabase/authz/v1"
)

type messageDoc struct {
	Name        string
	DisplayName string
	FullName    string
	Anchor      string
	Description string
	Fields      []fieldDoc
}

type fieldDoc struct {
	Name        string
	TypeFull    string
	Label       string
	Description string
	Required    bool
}

type fileEntry struct {
	fd       *descriptorpb.FileDescriptorProto
	comments map[string]string
}

type packageData struct {
	files      []fileEntry
	msgs       []messageDoc
	hasService bool
}

func main() {
	const descriptorPath = "docs/descriptor.bin"
	ignorePackages := []string{"google.api", "google.protobuf"}

	data, err := os.ReadFile(descriptorPath)
	if err != nil {
		fatalf("read descriptor set: %v", err)
	}

	var fds descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(data, &fds); err != nil {
		fatalf("unmarshal descriptor set: %v", err)
	}

	if err := os.MkdirAll("docs", 0o755); err != nil {
		fatalf("create docs dir: %v", err)
	}

	var (
		b             strings.Builder
		packages      = make(map[string]*packageData)
		messageByName = make(map[string]*messageDoc)
		pkgs          []string
	)

	for _, fd := range fds.File {
		pkg := fd.GetPackage()

		if slices.Contains(ignorePackages, pkg) {
			continue
		}

		pkgData := packages[pkg]

		if pkgData == nil {
			pkgData = &packageData{}
			packages[pkg] = pkgData
		}

		info := fd.GetSourceCodeInfo()
		comments := make(map[string]string)

		if info == nil {
			return
		}

		for _, loc := range info.Location {
			text := strings.TrimSpace(loc.GetLeadingComments())
			if text == "" {
				text = strings.TrimSpace(loc.GetTrailingComments())
			}
			if text == "" {
				continue
			}

			comments[pathKey(loc.Path)] = text
		}

		pkgData.files = append(pkgData.files, fileEntry{fd: fd, comments: comments})

		docs := collectMessages(pkg, fd.GetMessageType(), "", comments, []int32{4})
		start := len(pkgData.msgs)
		pkgData.msgs = append(pkgData.msgs, docs...)

		for i := range docs {
			md := &pkgData.msgs[start+i]
			messageByName[md.FullName] = md
		}

		if len(fd.GetService()) > 0 {
			pkgData.hasService = true
		}
	}

	for pkg, data := range packages {
		if data.hasService || len(data.msgs) > 0 {
			pkgs = append(pkgs, pkg)
		}
	}

	slices.Sort(pkgs)

	fmt.Fprint(&b, "# Terrabase API Reference\n\n")

	// Determine which message types are actually referenced (requests/responses/fields).
	neededMessages := make(map[string]bool)

	// Start by including all non-google.protobuf messages so top-level types are documented even if not referenced.
	for name := range messageByName {
		if strings.HasPrefix(name, "google.protobuf.") {
			continue
		}

		neededMessages[name] = true
	}

	var markType func(string)
	markType = func(typeName string) {
		if neededMessages[typeName] {
			return
		}

		md, ok := messageByName[typeName]
		if !ok {
			return
		}

		neededMessages[typeName] = true

		for _, f := range md.Fields {
			// only recurse into message types
			if messageByName[f.TypeFull] != nil {
				markType(f.TypeFull)
			}
		}
	}

	renderedMsgPkg := make(map[string]bool)

	for _, pkg := range pkgs {
		fmt.Fprintf(&b, "## %s\n\n", pkg)

		data := packages[pkg]

		for _, file := range data.files {
			for sIdx, svc := range file.fd.GetService() {
				svcHeading := fmt.Sprintf("%s (%s)", svc.GetName(), pkgLabel(pkg))

				fmt.Fprintf(&b, "### %s\n\n", svcHeading)

				if desc := commentFor(file.comments, 6, int32(sIdx)); desc != "" {
					b.WriteString(desc + "\n\n")
				}

				b.WriteString("| Name | Request | Response | Authentication Required | Required Scopes | Description |\n")
				b.WriteString("| --- | --- | --- | --- | --- | --- |\n")

				for mIdx, m := range svc.GetMethod() {
					opts := m.GetOptions()

					if opts == nil {
						return
					}

					var (
						authRequired bool
						adminOrSelf  bool
						scopes       []string
					)

					if proto.HasExtension(opts, authzv1.E_AuthRequired) {
						if v, ok := proto.GetExtension(opts, authzv1.E_AuthRequired).(bool); ok {
							authRequired = v
						}
					}

					if proto.HasExtension(opts, authzv1.E_RequiredScopes) {
						if vals, ok := proto.GetExtension(opts, authzv1.E_RequiredScopes).([]authzv1.Scope); ok {
							for _, s := range vals {
								scopes = append(scopes, fmt.Sprintf("`%s`", s.String()))
							}
						}
					}

					if proto.HasExtension(opts, authzv1.E_AdminOrSelf) {
						if v, ok := proto.GetExtension(opts, authzv1.E_AdminOrSelf).(bool); ok {
							adminOrSelf = v
						}
					}

					reqFull := strings.TrimPrefix(m.GetInputType(), ".")
					respFull := strings.TrimPrefix(m.GetOutputType(), ".")
					request := formatType(reqFull, messageByName)
					response := formatType(respFull, messageByName)

					var requiredScopes string

					if adminOrSelf {
						requiredScopes = "Admin or self"
					} else {
						requiredScopes = strings.Join(scopes, ", ")
					}

					markType(reqFull)
					markType(respFull)

					fmt.Fprintf(&b, "| `%s` | %s | %s | `%t` | %s | %s |\n", m.GetName(), request, response, authRequired, requiredScopes, commentFor(file.comments, 6, int32(sIdx), 2, int32(mIdx)))
				}

				b.WriteString("\n")
			}

			if len(data.msgs) > 0 {
				for _, msg := range data.msgs {
					if !neededMessages[msg.FullName] {
						continue
					}

					fmt.Fprintf(&b, "### %s\n\n", msg.DisplayName)

					if msg.Description != "" {
						b.WriteString(msg.Description + "\n\n")
					}

					if len(msg.Fields) == 0 {
						b.WriteString("- (no fields)\n\n")
						continue
					}

					b.WriteString("| Name | Type | Label | Required | Description |\n")
					b.WriteString("| --- | --- | --- | --- | --- |\n")

					for _, f := range msg.Fields {
						fmt.Fprintf(&b, "| `%s` | %s | %s | `%t` | %s |\n", f.Name, formatType(f.TypeFull, messageByName), f.Label, f.Required, f.Description)
					}

					b.WriteString("\n")
				}

				renderedMsgPkg[pkg] = true
			}
		}
	}

	outPath := filepath.Join("docs", "index.md")
	content := strings.TrimRight(b.String(), "\n") + "\n"

	if err := os.WriteFile(outPath, []byte(content), 0o644); err != nil {
		fatalf("write docs: %v", err)
	}
}

func collectMessages(pkg string, msgs []*descriptorpb.DescriptorProto, prefix string, comments map[string]string, pathPrefix []int32) []messageDoc {
	var docs []messageDoc

	for idx, m := range msgs {
		var segments []string
		if prefix != "" {
			segments = append(segments, prefix)
		}
		segments = append(segments, m.GetName())
		fullName := pkg + "." + strings.Join(segments, ".")
		displayName := m.GetName()
		if lbl := pkgLabel(pkg); lbl != "" {
			displayName = fmt.Sprintf("%s (%s)", m.GetName(), lbl)
		}

		var b strings.Builder
		prevHyphen := false
		for _, r := range fullName {
			switch {
			case r >= 'A' && r <= 'Z':
				b.WriteRune(r + ('a' - 'A'))
				prevHyphen = false
			case r >= 'a' && r <= 'z', r >= '0' && r <= '9', r == '_':
				b.WriteRune(r)
				prevHyphen = false
			case r == ' ' || r == '-':
				if !prevHyphen {
					b.WriteRune('-')
					prevHyphen = true
				}
			// other punctuation (including ".") is dropped to mirror GitHub-style heading IDs
			default:
				// skip
			}
		}
		anchor := strings.Trim(b.String(), "-")
		msgPath := append(pathPrefix, int32(idx))

		var fields []fieldDoc
		for i, f := range m.GetField() {
			fullType := strings.ToLower(strings.TrimPrefix(f.GetType().String(), "TYPE_"))

			if f.TypeName != nil && f.GetTypeName() != "" {
				fullType = strings.TrimPrefix(f.GetTypeName(), ".")
			}

			label := ""
			if f.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
				label = "repeated"
			}
			required := false
			if proto.HasExtension(f.Options, apipb.E_FieldBehavior) {
				if vals, ok := proto.GetExtension(f.Options, apipb.E_FieldBehavior).([]apipb.FieldBehavior); ok {
					for _, v := range vals {
						if v == apipb.FieldBehavior_REQUIRED {
							required = true
						}
					}
				}
			}
			fields = append(fields, fieldDoc{
				Name:        f.GetName(),
				TypeFull:    fullType,
				Label:       label,
				Required:    required,
				Description: commentForPath(comments, append(msgPath, 2, int32(i))),
			})
		}

		docs = append(docs, messageDoc{
			Name:        m.GetName(),
			DisplayName: displayName,
			FullName:    fullName,
			Anchor:      anchor,
			Description: commentForPath(comments, msgPath),
			Fields:      fields,
		})

		// Recurse into nested messages
		if nested := m.GetNestedType(); len(nested) > 0 {
			childPrefix := strings.Join(segments, ".")
			childDocs := collectMessages(pkg, nested, childPrefix, comments, append(msgPath, 3))
			docs = append(docs, childDocs...)
		}
	}
	return docs
}

func formatType(typeFull string, messages map[string]*messageDoc) string {
	display := typeFull

	if idx := strings.LastIndex(typeFull, "."); idx != -1 {
		display = typeFull[idx+1:]
	}

	if strings.HasPrefix(typeFull, "google.protobuf.") {
		return fmt.Sprintf("`%s`", display)
	}
	if md, ok := messages[typeFull]; ok {
		return fmt.Sprintf("[%s](#%s)", display, md.Anchor)
	}
	return display
}

func pkgLabel(pkg string) string {
	parts := strings.Split(pkg, ".")
	if len(parts) >= 2 {
		return strings.Join(parts[len(parts)-2:], ".")
	}
	if len(parts) == 1 {
		return parts[0]
	}
	return ""
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func commentFor(m map[string]string, path ...int32) string {
	return m[pathKey(path)]
}

func commentForPath(m map[string]string, path []int32) string {
	return m[pathKey(path)]
}

func pathKey(path []int32) string {
	if len(path) == 0 {
		return ""
	}
	var sb strings.Builder
	for i, p := range path {
		if i > 0 {
			sb.WriteString(".")
		}
		fmt.Fprintf(&sb, "%d", p)
	}
	return sb.String()
}
