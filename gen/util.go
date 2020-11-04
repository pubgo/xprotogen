package gen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pubgo/xerror"
	options "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/types/descriptorpb"
)

var rgxSyntaxError = regexp.MustCompile(`(\d+):\d+: `)

// formatBuffer format go code, panic when code is invalid
func formatBuffer(code string) (_ string, err error) {
	defer xerror.RespErr(&err)

	buf := bytes.NewBufferString(code)
	output, err := format.Source(buf.Bytes())
	if err == nil {
		return string(output), nil
	}

	matches := rgxSyntaxError.FindStringSubmatch(err.Error())
	if matches == nil {
		return "", xerror.New("failed to format template")
	}

	lineNum, _ := strconv.Atoi(matches[1])
	scanner := bufio.NewScanner(buf)
	errBuf := &bytes.Buffer{}
	line := 1
	for ; scanner.Scan(); line++ {
		if delta := line - lineNum; delta < -5 || delta > 5 {
			continue
		}

		if line == lineNum {
			errBuf.WriteString(">>>> ")
		} else {
			fmt.Fprintf(errBuf, "% 4d ", line)
		}
		errBuf.Write(scanner.Bytes())
		errBuf.WriteByte('\n')
	}

	return "", xerror.New("failed to format template\n\n" + string(errBuf.Bytes()))
}

// baseName returns the last path element of the name, with the last dotted suffix removed.
func baseName(name string) string {
	// First, find the last element
	if i := strings.LastIndex(name, "/"); i >= 0 {
		name = name[i+1:]
	}
	// Now drop the suffix
	if i := strings.LastIndex(name, "."); i >= 0 {
		name = name[0:i]
	}
	return name
}

// getGoPackage returns the file's go_package option.
// If it containts a semicolon, only the part before it is returned.
func getGoPackage(fd *descriptor.FileDescriptorProto) string {
	pkg := fd.GetOptions().GetGoPackage()
	if strings.Contains(pkg, ";") {
		parts := strings.Split(pkg, ";")
		if len(parts) > 2 {
			log.Fatalf("protoc-gen-nrpc: go_package '%s' contains more than 1 ';'", pkg)
		}
		pkg = parts[1]
	}

	return pkg
}

// goPackageOption interprets the file's go_package option.
// If there is no go_package, it returns ("", "", false).
// If there's a simple name, it returns ("", Pkg, true).
// If the option implies an import path, it returns (impPath, Pkg, true).
func goPackageOption(d *descriptor.FileDescriptorProto) (impPath, pkg string, ok bool) {
	pkg = getGoPackage(d)
	if pkg == "" {
		return
	}

	ok = true
	// The presence of a slash implies there's an import path.
	slash := strings.LastIndex(pkg, "/")
	if slash < 0 {
		return
	}

	impPath, pkg = pkg, pkg[slash+1:]
	// A semicolon-delimited suffix overrides the package name.
	sc := strings.IndexByte(impPath, ';')
	if sc < 0 {
		return
	}

	impPath, pkg = impPath[:sc], impPath[sc+1:]
	return
}

// goPackageName returns the Go package name to use in the
// generated Go file.  The result explicit reports whether the name
// came from an option go_package statement.  If explicit is false,
// the name was derived from the protocol buffer's package statement
// or the input file name.
func goPackageName(d *descriptor.FileDescriptorProto) (name string, explicit bool) {
	// Does the file have a "go_package" option?
	if _, pkg, ok := goPackageOption(d); ok {
		return pkg, true
	}

	// Does the file have a package clause?
	if pkg := d.GetPackage(); pkg != "" {
		return pkg, false
	}

	// Use the file base name.
	return baseName(d.GetName()), false
}

// splitMessageTypeName split a message type into (package name, type name)
func splitMessageTypeName(name string) (string, string) {
	if len(name) == 0 {
		log.Fatal("Empty message type")
	}
	if name[0] != '.' {
		log.Fatalf("Expect message type name to start with '.', but it is '%s'", name)
	}
	lastDot := strings.LastIndex(name, ".")
	return name[1:lastDot], name[lastDot+1:]
}

func getTypeName(pkgName, mthName string) string {
	mthName = strings.Trim(mthName, ".")
	mthName = strings.TrimLeft(mthName, pkgName)
	mthName = strings.Trim(mthName, ".")
	return mthName
}

func fileIsProto3(file *descriptor.FileDescriptorProto) bool {
	return file.GetSyntax() == "proto3"
}

// A GoImportPath is the import path of a Go package. e.g., "google.golang.org/genproto/protobuf".
type GoImportPath string

func (p GoImportPath) String() string { return strconv.Quote(string(p)) }

// And now lots of helper functions.

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// CamelCase returns the CamelCased name.
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercases names, it's extremely unlikely to have two fields
// with different capitalizations.
// In short, _my_field_name_2 becomes XMyFieldName_2.
func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		// Need a capital letter; drop the '_'.
		t = append(t, 'X')
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIILower(s[i+1]) {
			continue // Skip the underscore in s.
		}
		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if isASCIILower(c) {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && isASCIILower(s[i+1]) {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

// Is this field optional?
func isOptional(field *descriptor.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_OPTIONAL
}

// Is this field required?
func isRequired(field *descriptor.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_REQUIRED
}

// Is this field repeated?
func isRepeated(field *descriptor.FieldDescriptorProto) bool {
	return field.Label != nil && *field.Label == descriptor.FieldDescriptorProto_LABEL_REPEATED
}

// DefaultAPIOptions
// This generates an HttpRule that matches the gRPC mapping to HTTP/2 described in
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests
// i.e.:
//   * method is POST
//   * path is "/<service name>/<method name>"
//   * body should contain the serialized request message
func DefaultAPIOptions(pkg string, srv string, mth string) *options.HttpRule {
	log.Println(pkg, srv, mth)
	return &options.HttpRule{
		Pattern: &options.HttpRule_Post{
			Post: camel2Case(fmt.Sprintf("/%s_%s/%s", camel2Case(pkg), camel2Case(srv), camel2Case(mth))),
		},
		Body: "*",
	}
}

func ExtractAPIOptions(mth *descriptorpb.MethodDescriptorProto) (*options.HttpRule, error) {
	if mth.Options == nil {
		return nil, nil
	}

	if !proto.HasExtension(mth.Options, options.E_Http) {
		return nil, nil
	}

	ext, err := proto.GetExtension(mth.Options, options.E_Http)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	opts, ok := ext.(*options.HttpRule)
	if !ok {
		return nil, xerror.Fmt("extension is %T; want an HttpRule", ext)
	}

	return opts, nil
}

func ExtractHttpMethod(opts *options.HttpRule) (method string, path string) {
	var (
		httpMethod   string
		pathTemplate string
	)

	switch {
	case opts.GetGet() != "":
		httpMethod = "GET"
		pathTemplate = opts.GetGet()

	case opts.GetPut() != "":
		httpMethod = "PUT"
		pathTemplate = opts.GetPut()

	case opts.GetPost() != "":
		httpMethod = "POST"
		pathTemplate = opts.GetPost()

	case opts.GetDelete() != "":
		httpMethod = "DELETE"
		pathTemplate = opts.GetDelete()

	case opts.GetPatch() != "":
		httpMethod = "PATCH"
		pathTemplate = opts.GetPatch()

	case opts.GetCustom() != nil:
		custom := opts.GetCustom()
		httpMethod = custom.Kind
		pathTemplate = custom.Path

	default:
		return "", ""
	}

	return httpMethod, pathTemplate
}

func UnExport(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// camel2Case
// 驼峰式写法转为下划线写法
func camel2Case(name string) string {
	name = trim(name)
	buf := new(bytes.Buffer)
	for i, r := range name {
		if !unicode.IsUpper(r) {
			buf.WriteRune(r)
			continue
		}

		if i != 0 {
			buf.WriteRune('_')
		}
		buf.WriteRune(unicode.ToLower(r))
	}
	return strings.ReplaceAll(strings.ReplaceAll(buf.String(), ".", "_"), "__", "_")
}

func trim(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, ".")
	s = strings.Trim(s, "-")
	s = strings.Trim(s, "_")
	s = strings.Trim(s, "/")
	return s
}

func template1(tpl string, m M) string {
	t, err := template.New("main").Parse(tpl)
	xerror.PanicF(err, tpl)

	w := bytes.NewBuffer(nil)
	xerror.PanicF(t.Execute(w, m), tpl)
	return w.String()
}
