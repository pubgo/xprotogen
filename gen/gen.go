package gen

import (
	"fmt"
	"io/ioutil"
	logger "log"
	"os"
	"path"
	"strings"

	"github.com/dave/jennifer/jen"
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
		switch param {
		case "paths":
			if value == "source_relative" {
				t.PathsSourceRelative = true
			} else if value == "import" {
				t.PathsSourceRelative = false
			} else {
				log.Fatalf(`unknown path type %q: want "import" or "source_relative"`, value)
			}
		default:
			if len(param) > 0 && param[0] == 'M' {
				t.ImportMap[param[1:]] = value
			}
		}
		data[param] = value
	}

	for k, v := range data {
		fn(k, v)
	}
}

func (t *protoGen) Init(init func(fd *FileDescriptor)) (err error) {
	defer xerror.RespErr(&err)

	if init == nil {
		return xerror.New("[init] is nil")
	}

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
		var j = jen.NewFile(pkg)

		var data = make(M)
		data.Set("pkg", pkg)
		data.Set("fd", fd)

		init(&FileDescriptor{
			pkg:                 pkg,
			FileDescriptorProto: fd,
			J:                   j,
			M:                   data})

		if ext := path.Ext(name); ext == ".proto" {
			name = name[:len(name)-len(ext)]
		}

		t.response.File = append(t.response.File, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(name + ".pb." + t.name + ".go"),
			Content: proto.String(j.GoString()),
		})
	}

	data := xerror.PanicBytes(proto.Marshal(&t.response))
	xerror.PanicErr(os.Stdout.Write(data))
	return nil
}

type M map[string]interface{}

func (t M) Set(k string, v interface{}) {
	t[k] = v
}

func (t M) Fmt(template string, args ...interface{}) string {
	return fmt.Sprintf(Template(template, t), args...)
}

func (t M) P(template string, args ...interface{}) *jen.Statement {
	return jen.Id(t.Fmt(template, args...))
}

type Method struct {
	*descriptor.MethodDescriptorProto
	M
	ss *Service
}

func (t *Method) GetName() string {
	return CamelCase(t.MethodDescriptorProto.GetName())
}

func (t *Method) GetInputType() string {
	return getTypeName(t.ss.Pkg, t.MethodDescriptorProto.GetInputType())
}

func (t *Method) GetOutputType() string {
	return getTypeName(t.ss.Pkg, t.MethodDescriptorProto.GetOutputType())
}

func (t *Method) GetHttpMethod() (method string, path string) {
	hr, err := ExtractAPIOptions(t.MethodDescriptorProto)
	if err != nil || hr == nil {
		hr = DefaultAPIOptions(t.ss.Pkg, t.ss.GetName(), t.GetName())
	}
	t.GetServerStreaming()
	return ExtractHttpMethod(hr)
}

type FileDescriptor struct {
	M
	*descriptor.FileDescriptorProto
	J   *jen.File
	pkg string
}

func (t *FileDescriptor) Pkg() string {
	return t.pkg
}
func (t *FileDescriptor) GetService() []*Service {
	var ss []*Service
	for _, s := range t.FileDescriptorProto.GetService() {
		ss = append(ss, &Service{
			M:                      t.M,
			ServiceDescriptorProto: s,
			Pkg:                    t.pkg,
			J:                      t.J,
		})
	}
	return ss
}

type Service struct {
	M
	*descriptor.ServiceDescriptorProto
	Pkg string
	J   *jen.File
}

func (t *Service) GetMethod() (methods []*Method) {
	for _, mth := range t.ServiceDescriptorProto.GetMethod() {
		methods = append(methods, &Method{
			M:                     t.M,
			ss:                    t,
			MethodDescriptorProto: mth,
		})
	}
	return methods
}

func (t *Service) GetName() string {
	return getTypeName(t.Pkg, t.ServiceDescriptorProto.GetName())
}
