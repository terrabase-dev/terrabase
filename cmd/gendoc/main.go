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
	TypeDisplay string
	Label       string
	Description string
	Required    bool
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

	var b strings.Builder
	anchorMap := make(map[string]string)

	// Precompute anchors and docs for all messages so we can link Request/Response types.
	messageIndex := make(map[string][]string)
	messageDocs := make(map[string][]messageDoc)
	messageByName := make(map[string]*messageDoc)
	anchorByFull := make(map[string]string)
	seenPkgs := make(map[string]bool)
	var pkgs []string
	for _, fd := range fds.File {
		pkg := fd.GetPackage()
		comments := buildCommentMap(fd.GetSourceCodeInfo())
		names, docs := collectMessages(pkg, fd.GetMessageType(), "", anchorMap, comments, []int32{4})
		if slices.Contains(ignorePackages, pkg) {
			continue
		}
		if len(names) > 0 {
			messageIndex[pkg] = append(messageIndex[pkg], names...)
			messageDocs[pkg] = append(messageDocs[pkg], docs...)
			for i := range docs {
				md := docs[i]
				messageByName[md.FullName] = &docs[i]
				anchorByFull[md.FullName] = md.Anchor
			}
		}
		if !seenPkgs[pkg] && !(len(fd.GetService()) == 0 && len(messageDocs[pkg]) == 0) {
			seenPkgs[pkg] = true
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
	markType := func(typeName string) {}
	var markFunc func(string)
	markFunc = func(typeName string) {
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
				markFunc(f.TypeFull)
			}
		}
	}
	markType = markFunc

	renderedMsgPkg := make(map[string]bool)

	for _, pkg := range pkgs {
		fmt.Fprintf(&b, "## %s\n\n", pkg)
		for _, fd := range fds.File {
			if fd.GetPackage() != pkg {
				continue
			}
			comments := buildCommentMap(fd.GetSourceCodeInfo())
			for sIdx, svc := range fd.GetService() {
				svcHeading := fmt.Sprintf("%s (%s)", svc.GetName(), pkgLabel(pkg))
				fmt.Fprintf(&b, "### %s\n\n", svcHeading)
				if desc := commentFor(comments, 6, int32(sIdx)); desc != "" {
					b.WriteString(desc + "\n\n")
				}
				b.WriteString("| Name | Request | Response | Authentication Required | Required Scopes | Description |\n")
				b.WriteString("| --- | --- | --- | --- | --- | --- |\n")

				for mIdx, m := range svc.GetMethod() {
					authReq, adminOrSelf, scopes := readMethodOptions(m.GetOptions())

					name := m.GetName()
					reqFull := strings.TrimPrefix(m.GetInputType(), ".")
					respFull := strings.TrimPrefix(m.GetOutputType(), ".")
					request := formatType(basename(reqFull), reqFull, anchorByFull)
					response := formatType(basename(respFull), respFull, anchorByFull)
					desc := commentFor(comments, 6, int32(sIdx), 2, int32(mIdx))

					var requiredScopes string

					if adminOrSelf {
						requiredScopes = "Admin or self"
					} else {
						requiredScopes = strings.Join(scopes, ", ")
					}

					markType(reqFull)
					markType(respFull)
					fmt.Fprintf(&b, "| `%s` | %s | %s | `%t` | %s | %s |\n", name, request, response, authReq, requiredScopes, desc)
				}

				b.WriteString("\n")
			}

			if msgs := messageIndex[pkg]; len(msgs) > 0 {
				for _, msg := range messageDocs[pkg] {
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
						name := f.Name
						typ := formatType(f.TypeDisplay, f.TypeFull, anchorByFull)
						label := f.Label
						desc := f.Description
						fmt.Fprintf(&b, "| `%s` | `%s` | %s | `%t` | %s |\n", name, typ, label, f.Required, desc)
					}
					b.WriteString("\n")
				}
				renderedMsgPkg[pkg] = true
			}
		}
	}

	// Render message docs for packages without services (e.g., well-known types)
	for _, pkg := range pkgs {
		if renderedMsgPkg[pkg] {
			continue
		}
		docs := messageDocs[pkg]
		if len(docs) == 0 {
			continue
		}
		printed := false
		for _, msg := range docs {
			if !neededMessages[msg.FullName] {
				continue
			}
			if !printed {
				fmt.Fprintf(&b, "### %s Messages\n\n", pkg)
				printed = true
			}
			fmt.Fprintf(&b, "#### %s\n\n", msg.DisplayName)
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
				name := f.Name
				typ := formatType(f.TypeDisplay, f.TypeFull, anchorByFull)
				label := f.Label
				required := "no"
				if f.Required {
					required = "yes"
				}
				desc := f.Description
				fmt.Fprintf(&b, "| %s | %s | %s | %s | %s |\n", name, typ, label, required, desc)
			}
			b.WriteString("\n")
		}
	}

	outPath := filepath.Join("docs", "index.md")
	content := strings.TrimRight(b.String(), "\n") + "\n"
	if err := os.WriteFile(outPath, []byte(content), 0o644); err != nil {
		fatalf("write docs: %v", err)
	}
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

func collectMessages(pkg string, msgs []*descriptorpb.DescriptorProto, prefix string, anchors map[string]string, comments map[string]string, pathPrefix []int32) ([]string, []messageDoc) {
	var names []string
	var docs []messageDoc

	for idx, m := range msgs {
		segments := []string{}
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
		names = append(names, fullName)
		anchors[fullName] = anchor

		msgPath := append(pathPrefix, int32(idx))
		docs = append(docs, messageDoc{
			Name:        m.GetName(),
			DisplayName: displayName,
			FullName:    fullName,
			Anchor:      anchor,
			Description: commentForPath(comments, msgPath),
			Fields:      collectFields(m, comments, msgPath),
		})

		// Recurse into nested messages
		if nested := m.GetNestedType(); len(nested) > 0 {
			childPrefix := strings.Join(segments, ".")
			childNames, childDocs := collectMessages(pkg, nested, childPrefix, anchors, comments, append(msgPath, 3))
			names = append(names, childNames...)
			docs = append(docs, childDocs...)
		}
	}
	return names, docs
}

func anchorFor(fullName string) string {
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
	anchor := b.String()
	anchor = strings.Trim(anchor, "-")
	return anchor
}

func formatType(display, typeFull string, anchors map[string]string) string {
	if strings.HasPrefix(typeFull, "google.protobuf.") {
		return fmt.Sprintf("`%s`", display)
	}
	if anchor, ok := anchors[typeFull]; ok {
		return fmt.Sprintf("[%s](#%s)", display, anchor)
	}
	return display
}

func basename(typeName string) string {
	parts := strings.Split(typeName, ".")
	return parts[len(parts)-1]
}

func collectFields(msg *descriptorpb.DescriptorProto, comments map[string]string, path []int32) []fieldDoc {
	var fields []fieldDoc
	for i, f := range msg.GetField() {
		fullType := fieldType(f)
		fields = append(fields, fieldDoc{
			Name:        f.GetName(),
			TypeFull:    fullType,
			TypeDisplay: basename(fullType),
			Label:       fieldLabel(f),
			Required:    fieldRequired(f),
			Description: commentForPath(comments, append(path, 2, int32(i))),
		})
	}
	return fields
}

func fieldType(f *descriptorpb.FieldDescriptorProto) string {
	if f.TypeName != nil && f.GetTypeName() != "" {
		return strings.TrimPrefix(f.GetTypeName(), ".")
	}
	return strings.ToLower(strings.TrimPrefix(f.GetType().String(), "TYPE_"))
}

func fieldLabel(f *descriptorpb.FieldDescriptorProto) string {
	if f.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
		return "repeated"
	}
	return ""
}

func fieldRequired(f *descriptorpb.FieldDescriptorProto) bool {
	if f == nil {
		return false
	}
	if proto.HasExtension(f.Options, apipb.E_FieldBehavior) {
		if vals, ok := proto.GetExtension(f.Options, apipb.E_FieldBehavior).([]apipb.FieldBehavior); ok {
			for _, v := range vals {
				if v == apipb.FieldBehavior_REQUIRED {
					return true
				}
			}
		}
	}
	return false
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

func buildCommentMap(info *descriptorpb.SourceCodeInfo) map[string]string {
	m := make(map[string]string)
	if info == nil {
		return m
	}
	for _, loc := range info.Location {
		text := strings.TrimSpace(loc.GetLeadingComments())
		if text == "" {
			text = strings.TrimSpace(loc.GetTrailingComments())
		}
		if text == "" {
			continue
		}
		key := pathKey(loc.Path)
		m[key] = text
	}
	return m
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
