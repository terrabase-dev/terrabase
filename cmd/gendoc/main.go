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

type enumDoc struct {
	Name        string
	DisplayName string
	FullName    string
	Anchor      string
	Description string
	Values      []enumValueDoc
}

type enumValueDoc struct {
	Name        string
	Number      int32
	Description string
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
	enums      []enumDoc
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
		enumByName    = make(map[string]*enumDoc)
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
			continue
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

		enumStart := len(pkgData.enums)
		enumDocs := collectEnums(pkg, fd.GetEnumType(), "", comments, []int32{5})
		pkgData.enums = append(pkgData.enums, enumDocs...)

		for i := range enumDocs {
			ed := &pkgData.enums[enumStart+i]
			enumByName[ed.FullName] = ed
		}

		nestedEnumStart := len(pkgData.enums)
		docs := collectMessages(pkg, fd.GetMessageType(), "", comments, []int32{4}, &pkgData.enums)
		start := len(pkgData.msgs)
		pkgData.msgs = append(pkgData.msgs, docs...)

		for i := range docs {
			md := &pkgData.msgs[start+i]
			messageByName[md.FullName] = md
		}

		if len(pkgData.enums) > nestedEnumStart {
			for i := nestedEnumStart; i < len(pkgData.enums); i++ {
				ed := &pkgData.enums[i]
				enumByName[ed.FullName] = ed
			}
		}

		if len(fd.GetService()) > 0 {
			pkgData.hasService = true
		}
	}

	for pkg, data := range packages {
		if data.hasService || len(data.msgs) > 0 || len(data.enums) > 0 {
			pkgs = append(pkgs, pkg)
		}
	}

	slices.Sort(pkgs)

	fmt.Fprint(&b, "# Terrabase API Reference\n\n")

	// Determine which message types are actually referenced (requests/responses/fields).
	neededMessages := make(map[string]bool)
	neededEnums := make(map[string]bool)

	// Start by including all non-google.protobuf messages and enums so top-level types are documented even if not referenced.
	// Dependency traversal ensures nested message/enum references are also marked as needed.

	var markMessage func(string)
	markEnum := func(enumName string) {
		if neededEnums[enumName] {
			return
		}

		ed, ok := enumByName[enumName]
		if !ok {
			return
		}

		neededEnums[ed.FullName] = true
	}

	markMessage = func(typeName string) {
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
				markMessage(f.TypeFull)
			}

			if enumByName[f.TypeFull] != nil {
				markEnum(f.TypeFull)
			}
		}
	}

	for name := range messageByName {
		if strings.HasPrefix(name, "google.protobuf.") {
			continue
		}

		markMessage(name)
	}

	for name := range enumByName {
		if strings.HasPrefix(name, "google.protobuf.") {
			continue
		}

		markEnum(name)
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
					authRequired, adminOrSelf, scopes := readMethodOptions(m.GetOptions())

					reqFull := strings.TrimPrefix(m.GetInputType(), ".")
					respFull := strings.TrimPrefix(m.GetOutputType(), ".")
					request := formatType(reqFull, messageByName, enumByName)
					response := formatType(respFull, messageByName, enumByName)

					var requiredScopes string

					if adminOrSelf {
						requiredScopes = "Admin or self"
					} else {
						requiredScopes = strings.Join(scopes, ", ")
					}

					markMessage(reqFull)
					markMessage(respFull)

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
						fmt.Fprintf(&b, "| `%s` | %s | %s | `%t` | %s |\n", f.Name, formatType(f.TypeFull, messageByName, enumByName), f.Label, f.Required, f.Description)
					}

					b.WriteString("\n")
				}

				renderedMsgPkg[pkg] = true
			}

			if len(data.enums) > 0 {
				for _, enum := range data.enums {
					if !neededEnums[enum.FullName] {
						continue
					}

					fmt.Fprintf(&b, "### %s\n\n", enum.DisplayName)

					if enum.Description != "" {
						b.WriteString(enum.Description + "\n\n")
					}

					if len(enum.Values) == 0 {
						b.WriteString("- (no values)\n\n")
						continue
					}

					b.WriteString("| Name | Number | Description |\n")
					b.WriteString("| --- | --- | --- |\n")

					for _, v := range enum.Values {
						fmt.Fprintf(&b, "| `%s` | `%d` | %s |\n", v.Name, v.Number, v.Description)
					}

					b.WriteString("\n")
				}
			}
		}
	}

	outPath := filepath.Join("docs", "index.md")
	content := strings.TrimRight(b.String(), "\n") + "\n"

	if err := os.WriteFile(outPath, []byte(content), 0o644); err != nil {
		fatalf("write docs: %v", err)
	}
}

func collectMessages(pkg string, msgs []*descriptorpb.DescriptorProto, prefix string, comments map[string]string, pathPrefix []int32, enums *[]enumDoc) []messageDoc {
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

		anchor := anchorFor(displayName)
		msgPath := append(pathPrefix, int32(idx))

		if enumTypes := m.GetEnumType(); len(enumTypes) > 0 {
			childPrefix := strings.Join(segments, ".")
			enumDocs := collectEnums(pkg, enumTypes, childPrefix, comments, append(msgPath, 4))
			*enums = append(*enums, enumDocs...)
		}

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
			childDocs := collectMessages(pkg, nested, childPrefix, comments, append(msgPath, 3), enums)
			docs = append(docs, childDocs...)
		}
	}
	return docs
}

func collectEnums(pkg string, enums []*descriptorpb.EnumDescriptorProto, prefix string, comments map[string]string, pathPrefix []int32) []enumDoc {
	var docs []enumDoc

	for idx, e := range enums {
		var segments []string
		if prefix != "" {
			segments = append(segments, prefix)
		}
		segments = append(segments, e.GetName())
		fullName := pkg + "." + strings.Join(segments, ".")
		displayName := e.GetName()
		if lbl := pkgLabel(pkg); lbl != "" {
			displayName = fmt.Sprintf("%s (%s)", e.GetName(), lbl)
		}

		enumPath := append(pathPrefix, int32(idx))
		anchor := anchorFor(displayName)

		var values []enumValueDoc
		for vIdx, v := range e.GetValue() {
			values = append(values, enumValueDoc{
				Name:        v.GetName(),
				Number:      v.GetNumber(),
				Description: commentForPath(comments, append(enumPath, 2, int32(vIdx))),
			})
		}

		docs = append(docs, enumDoc{
			Name:        e.GetName(),
			DisplayName: displayName,
			FullName:    fullName,
			Anchor:      anchor,
			Description: commentForPath(comments, enumPath),
			Values:      values,
		})
	}

	return docs
}

func formatType(typeFull string, messages map[string]*messageDoc, enums map[string]*enumDoc) string {
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
	if ed, ok := enums[typeFull]; ok {
		return fmt.Sprintf("[%s](#%s)", display, ed.Anchor)
	}
	return display
}

func anchorFor(s string) string {
	var b strings.Builder
	prevDash := false

	for _, r := range strings.ToLower(s) {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_':
			b.WriteRune(r)
			prevDash = false
		case r == ' ' || r == '-':
			if !prevDash {
				b.WriteRune('-')
				prevDash = true
			}
		// drop all other punctuation (e.g., dots) to match markdown heading IDs
		default:
			// skip
		}
	}

	return strings.Trim(b.String(), "-")
}

func readMethodOptions(opts *descriptorpb.MethodOptions) (bool, bool, []string) {
	if opts == nil {
		return false, false, nil
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

	return authRequired, adminOrSelf, scopes
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
