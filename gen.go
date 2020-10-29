package xprotogen

import (
	"fmt"
	"io/ioutil"
	logger "log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pubgo/xerror"
	options "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/types/descriptorpb"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

var log = logger.New(os.Stderr, "", logger.LstdFlags|logger.Lshortfile)

func New(name string) *protoGen {
	p := &protoGen{name: name}

	data := xerror.PanicBytes(ioutil.ReadAll(os.Stdin))
	xerror.Panic(proto.Unmarshal(data, &p.request))

	if len(p.request.GetFileToGenerate()) == 0 {
		log.Println("GetFileToGenerate is empty")
	}

	return p
}

type protoGen struct {
	init                func(pkg string, fd *descriptor.FileDescriptorProto, j *jen.File) error
	name                string
	request             plugin.CodeGeneratorRequest
	response            plugin.CodeGeneratorResponse
	PathsSourceRelative bool
}

func (t *protoGen) Save() (err error) {
	defer xerror.RespErr(&err)
	data := xerror.PanicBytes(proto.Marshal(&t.response))
	xerror.PanicErr(os.Stdout.Write(data))
	return nil
}

type Service struct {
	*descriptor.ServiceDescriptorProto
	Name string
	Pkg  string
	J    *jen.File
}

func (t *Service) P(a ...interface{}) *jen.Statement {
	return t.J.Id(fmt.Sprintln(a...))
}

func (t *Service) TypeName(mthName string) string {
	return getTypeName(t.Pkg, mthName)
}

func (t *protoGen) Init(fn func(pkg string, fd *descriptor.FileDescriptorProto, j *jen.File) error) {
	t.init = fn
}

func (t *protoGen) Service(fn func(s *Service)) (err error) {
	defer xerror.RespErr(&err)

	for _, name := range t.request.GetFileToGenerate() {
		var fd *descriptor.FileDescriptorProto
		for _, fd = range t.request.GetProtoFile() {
			if name == fd.GetName() {
				break
			}
		}

		if fd == nil {
			return xerror.Fmt("could not find the .proto file for %s", name)
		}

		if len(fd.GetService()) == 0 {
			continue
		}

		pkg, _ := goPackageName(fd)
		j := jen.NewFile(pkg)

		if t.init != nil {
			xerror.Panic(t.init(pkg, fd, j))
		}

		for _, ss := range fd.GetService() {
			fn(&Service{J: j, ServiceDescriptorProto: ss, Pkg: pkg, Name: CamelCase(ss.GetName())})
		}

		if ext := path.Ext(name); ext == ".proto" {
			name = name[:len(name)-len(ext)]
		}

		t.response.File = append(t.response.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(name + ".pb." + t.name + ".go"),
			Content: proto.String(j.GoString()),
		})
	}

	return nil
}

func (t *protoGen) Parameter(fn func(key, value string)) {
	parameter := t.request.GetParameter()
	for _, param := range strings.Split(parameter, ",") {
		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = strings.TrimSpace(param[i+1:])
			param = strings.TrimSpace(param[0:i])
		}

		if param == "" {
			continue
		}

		// case "import_prefix", "import_path":
		switch param {
		case "paths":
			if value == "source_relative" {
				t.PathsSourceRelative = true
			} else if value == "import" {
				t.PathsSourceRelative = false
			} else {
				log.Fatalf(`unknown path type %q: want "import" or "source_relative"`, value)
			}
		}

		fn(param, value)
	}
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

func splitTypePath(name string) []string {
	if len(name) == 0 {
		log.Fatal("Empty message type")
	}
	if name[0] != '.' {
		log.Fatalf("Expect message type name to start with '.', but it is '%s'", name)
	}
	return strings.Split(name[1:], ".")
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

// CamelCaseSlice is like CamelCase, but the argument is a slice of strings to
// be joined with "_".
func CamelCaseSlice(elem []string) string { return CamelCase(strings.Join(elem, "_")) }

// dottedSlice turns a sliced name into a dotted name.
func dottedSlice(elem []string) string { return strings.Join(elem, ".") }

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

// Is this field a scalar numeric type?
func isScalar(field *descriptor.FieldDescriptorProto) bool {
	if field.Type == nil {
		return false
	}
	switch *field.Type {
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
		descriptor.FieldDescriptorProto_TYPE_FLOAT,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_UINT64,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_FIXED64,
		descriptor.FieldDescriptorProto_TYPE_FIXED32,
		descriptor.FieldDescriptorProto_TYPE_BOOL,
		descriptor.FieldDescriptorProto_TYPE_UINT32,
		descriptor.FieldDescriptorProto_TYPE_ENUM,
		descriptor.FieldDescriptorProto_TYPE_SFIXED32,
		descriptor.FieldDescriptorProto_TYPE_SFIXED64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64:
		return true
	default:
		return false
	}
}

func DefaultAPIOptions(pkg string, srv string, mth string) (*options.HttpRule, error) {
	// This generates an HttpRule that matches the gRPC mapping to HTTP/2 described in
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests
	// i.e.:
	//   * method is POST
	//   * path is "/<service name>/<method name>"
	//   * body should contain the serialized request message
	rule := &options.HttpRule{
		Pattern: &options.HttpRule_Post{
			Post: fmt.Sprintf("/%s.%s/%s", pkg, srv, mth),
		},
		Body: "*",
	}
	return rule, nil
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
