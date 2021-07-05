package gen

import (
	"go/format"
	"io/ioutil"
	logger "log"
	"os"
	"path"
	"strings"

	_ "github.com/flosch/pongo2-addons"
	"github.com/flosch/pongo2/v4"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pubgo/xerror"
)

var log = logger.New(os.Stderr, "xprotogen: ", logger.LstdFlags|logger.Lshortfile)

type opts struct {
	onlyService bool
}

type Opt = func(opt *opts) error

func OnlyService() Opt {
	return func(opt *opts) error {
		opt.onlyService = true
		return nil
	}
}

func New(name string, opts ...Opt) *protoGen {
	p := &protoGen{name: name, ImportMap: make(map[string]string)}

	for i := range opts {
		xerror.Panic(opts[i](&p.opts))
	}

	data := xerror.PanicBytes(ioutil.ReadAll(os.Stdin))
	xerror.Panic(proto.Unmarshal(data, &p.request))

	if len(p.request.GetFileToGenerate()) == 0 {
		log.Println("GetFileToGenerate is empty")
	}

	return p
}

type protoGen struct {
	name                string
	request             plugin.CodeGeneratorRequest
	response            plugin.CodeGeneratorResponse
	PathsSourceRelative bool
	ImportMap           map[string]string
	opts                opts
}

func (t *protoGen) Parameter(fn func(key, value string)) {
	for _, param := range strings.Split(t.request.GetParameter(), ",") {
		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = strings.TrimSpace(param[i+1:])
			param = strings.TrimSpace(param[0:i])
		}

		if param == "" {
			continue
		}

		switch {
		case param == "paths":
			if value == "source_relative" {
				t.PathsSourceRelative = true
			} else if value == "import" {
				t.PathsSourceRelative = false
			} else {
				log.Fatalf(`unknown path type %q: want "import" or "source_relative"`, value)
			}
		case param[0] == 'M':
			t.ImportMap[param[1:]] = value
		default:
			fn(param, value)
		}
	}
}

func (t *protoGen) GenWithTpl(fns ...func(fd *FileDescriptor) string) (err error) {
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

		if t.opts.onlyService && len(fd.GetService()) == 0 {
			continue
		}

		pkg, _ := goPackageName(fd)

		fd1 := &FileDescriptor{Pkg: pkg, FileDescriptorProto: fd}
		ctx := pongo2.Context{
			"fileName": fd.GetName(),
			"fd":       fd1,
			"unExport": UnExport,
			"pkg": func() string {
				if strings.Contains(pkg, "/") {
					var names = strings.Split(pkg, "/")
					return names[0]
				}

				if strings.Contains(pkg, ".") {
					var names = strings.Split(pkg, ".")
					return names[0]
				}

				return pkg
			},
		}
		var data []string
		for i := range fns {
			data = append(data, Template(fns[i](fd1), ctx))
		}

		if ext := path.Ext(name); ext == ".proto" {
			name = name[:len(name)-len(ext)]
		}

		var code = strings.Join(data, "\n")
		dt, err := format.Source([]byte(code))
		xerror.Panic(err, code)
		t.response.File = append(t.response.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(name + ".pb." + t.name + ".go"),
			Content: proto.String(string(dt)),
		})
	}

	data := xerror.PanicBytes(proto.Marshal(&t.response))
	xerror.PanicErr(os.Stdout.Write(data))
	return nil
}

type FileDescriptor struct {
	Pkg string
	*descriptor.FileDescriptorProto
}

func (t *FileDescriptor) GetService() []*Service {
	var ss []*Service
	for _, s := range t.FileDescriptorProto.GetService() {
		s1 := &Service{
			ServiceDescriptorProto: s,
			Pkg:                    t.Pkg,
		}
		s1.Srv = s1.GetName()
		ss = append(ss, s1)
	}
	return ss
}

type Service struct {
	Pkg string
	Srv string
	*descriptor.ServiceDescriptorProto
}

func (t *Service) GetMethod() (methods []*Method) {
	for _, mth := range t.ServiceDescriptorProto.GetMethod() {
		m := &Method{Srv: t.GetName(), md: mth, Pkg: t.Pkg}
		m.Name = m.GetName()
		m.InType = m.GetInputType()
		m.OutType = m.GetOutputType()
		m.CS = m.GetClientStreaming()
		m.SS = m.GetServerStreaming()

		httpMethod, httpPath, defaultUrl := m.GetHttpMethod()
		m.HttpMethod = httpMethod
		m.HttpPath = httpPath
		m.DefaultUrl = defaultUrl
		methods = append(methods, m)
	}
	return methods
}

func (t *Service) GetName() string {
	return CamelCase(getTypeName(t.Pkg, t.ServiceDescriptorProto.GetName()))
}

type Method struct {
	InType     string
	OutType    string
	CS         bool
	SS         bool
	Name       string
	HttpMethod string
	HttpPath   string
	DefaultUrl bool
	Srv        string
	Pkg        string
	md         *descriptor.MethodDescriptorProto
}

func (t *Method) GetClientStreaming() bool              { return t.md.GetClientStreaming() }
func (t *Method) GetServerStreaming() bool              { return t.md.GetServerStreaming() }
func (t *Method) GetOptions() *descriptor.MethodOptions { return t.md.GetOptions() }
func (t *Method) GetName() string                       { return CamelCase(t.md.GetName()) }
func (t *Method) GetInputType() string                  { return getTypeName(t.Pkg, t.md.GetInputType()) }
func (t *Method) GetOutputType() string                 { return getTypeName(t.Pkg, t.md.GetOutputType()) }

func (t *Method) GetHttpMethod() (method string, path string, defaultUrl bool) {
	hr, err := ExtractAPIOptions(t.md)
	if err != nil || hr == nil {
		defaultUrl = true
		hr = DefaultAPIOptions(t.Pkg, t.Srv, t.GetName())
	}
	method, path = ExtractHttpMethod(hr)
	return
}
