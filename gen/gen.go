package gen

import (
	"fmt"
	"io/ioutil"
	logger "log"
	"os"
	"path"
	"strings"

	"github.com/pubgo/xerror"

	"github.com/dave/jennifer/jen"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
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

func (t *protoGen) Init(init func(fd *FileDescriptor)) (err error) {
	defer xerror.RespErr(&err)

	if init == nil {
		return xerror.New("[init] should not be nil")
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

		init(&FileDescriptor{
			FileDescriptorProto: fd,
			Jen:                 j,
			M:                   M{"pkg": pkg, "fd": fd}})

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

func (t M) Clone() M {
	var data = make(M)
	for k, v := range t {
		data[k] = v
	}
	return data
}

func (t M) Set(k string, v interface{}) {
	t[k] = v
}

func (t M) P(template string) string {
	return template1(template, t)
}

type FileDescriptor struct {
	M
	*descriptor.FileDescriptorProto
	Jen *jen.File
}

func (t *FileDescriptor) GetService() []*Service {
	var ss []*Service
	for _, s := range t.FileDescriptorProto.GetService() {
		s1 := &Service{ServiceDescriptorProto: s, M: t.M.Clone(), Jen: t.Jen}
		s1.Set("srv", s1.GetName())
		s1.Set("srv1", UnExport(s1.GetName()))
		ss = append(ss, s1)
	}
	return ss
}

type Service struct {
	*descriptor.ServiceDescriptorProto
	M
	Jen *jen.File
}

func (t *Service) GetMethod() (methods []*Method) {
	for _, mth := range t.ServiceDescriptorProto.GetMethod() {
		m := &Method{M: t.M.Clone(), MethodDescriptorProto: mth, Jen: t.Jen}
		m.Set("mthName", m.GetName())
		m.Set("inType", m.GetInputType())
		m.Set("outType", m.GetOutputType())
		m.Set("cs", m.GetClientStreaming())
		m.Set("ss", m.GetServerStreaming())
		methods = append(methods, m)
	}
	return methods
}

func (t *Service) GetName() string {
	return CamelCase(getTypeName(fmt.Sprintf("%s", t.M["pkg"]), t.ServiceDescriptorProto.GetName()))
}

type Method struct {
	*descriptor.MethodDescriptorProto
	M
	Jen *jen.File
}

func (t *Method) GetName() string {
	return CamelCase(t.MethodDescriptorProto.GetName())
}

func (t *Method) GetInputType() string {
	return getTypeName(fmt.Sprintf("%s", t.M["pkg"]), t.MethodDescriptorProto.GetInputType())
}

func (t *Method) GetOutputType() string {
	return getTypeName(fmt.Sprintf("%s", t.M["pkg"]), t.MethodDescriptorProto.GetOutputType())
}

func (t *Method) GetHttpMethod() (method string, path string) {
	hr, err := ExtractAPIOptions(t.MethodDescriptorProto)
	if err != nil || hr == nil {
		hr = DefaultAPIOptions(fmt.Sprintf("%s", t.M["pkg"]), fmt.Sprintf("%s", t.M["srv"]), t.GetName())
	}
	return ExtractHttpMethod(hr)
}
