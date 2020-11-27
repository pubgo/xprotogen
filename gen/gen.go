package gen

import (
	"go/format"
	"io/ioutil"
	logger "log"
	"os"
	"path"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pubgo/xerror"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

var log = logger.New(os.Stderr, "xprotogen: ", logger.LstdFlags|logger.Lshortfile)

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
	name                string
	request             plugin.CodeGeneratorRequest
	response            plugin.CodeGeneratorResponse
	PathsSourceRelative bool
	ImportMap           map[string]string
}

func (t *protoGen) Parameter(fn func(key, value string)) {
	var data = make(map[string]string)
	for _, param := range strings.Split(t.request.GetParameter(), ",") {
		var value string
		if i := strings.Index(param, "="); i >= 0 {
			value = strings.TrimSpace(param[i+1:])
			param = strings.TrimSpace(param[0:i])
		}

		if param == "" {
			continue
		}

		// case "import_prefix", "import_path":
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
			data[param] = value
		}
	}

	for k, v := range data {
		fn(k, v)
	}
}

func (t *protoGen) GenWithTpl(tpl string) (err error) {
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
		data := template1(tpl, pongo2.Context{
			"fileName": fd.GetName(),
			"pkg":      pkg,
			"fd":       &FileDescriptor{Pkg: pkg, FileDescriptorProto: fd},
			"unExport": UnExport,
		})

		if ext := path.Ext(name); ext == ".proto" {
			name = name[:len(name)-len(ext)]
		}

		dt, err := format.Source([]byte(data))
		xerror.PanicF(err, data)

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
		m := &Method{Srv: t.GetName(), MethodDescriptorProto: mth, Pkg: t.Pkg}
		m.Name = m.GetName()
		m.InType = m.GetInputType()
		m.OutType = m.GetOutputType()
		m.CS = m.GetClientStreaming()
		m.SS = m.GetServerStreaming()

		httpMethod, httpPath := m.GetHttpMethod()
		m.HttpMethod = httpMethod
		m.HttpPath = httpPath
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
	Srv        string
	Pkg        string
	*descriptor.MethodDescriptorProto
}

func (t *Method) GetName() string {
	return CamelCase(t.MethodDescriptorProto.GetName())
}

func (t *Method) GetInputType() string {
	return getTypeName(t.Pkg, t.MethodDescriptorProto.GetInputType())
}

func (t *Method) GetOutputType() string {
	return getTypeName(t.Pkg, t.MethodDescriptorProto.GetOutputType())
}

func (t *Method) GetHttpMethod() (method string, path string) {
	hr, err := ExtractAPIOptions(t.MethodDescriptorProto)
	if err != nil || hr == nil {
		hr = DefaultAPIOptions(t.Pkg, t.Srv, t.GetName())
	}
	return ExtractHttpMethod(hr)
}
