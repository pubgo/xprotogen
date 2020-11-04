// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: examples/proto/hello/api.proto
package hello

import (
	context "context"
	fmt "fmt"
	math "math"

	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TestApi service
type TestApiService interface {
	Version(ctx context.Context, in *TestReq, opts ...client.CallOption) (*TestApiOutput, error)
	VersionTest(ctx context.Context, in *TestReq, opts ...client.CallOption) (*TestApiOutput, error)
}

type testApiService struct {
	c    client.Client
	name string
}

func NewTestApiService(name string, c client.Client) TestApiService {
	return &testApiService{
		c:    c,
		name: name,
	}
}
func (c *TestApiService) Version(ctx context.Context, in *TestReq, opts ...client.CallOption) (*TestApiOutput, error) {

	req := c.c.NewRequest(c.name, "TestApi.Version", in)
	out := new(TestApiOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type TestApi_VersionService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type TestApiVersion struct {
	stream client.Stream
}

func (x *TestApiVersion) Close() error {
	return x.stream.Close()
}

func (x *TestApiVersion) Context() context.Context {
	return x.stream.Context()
}

func (x *TestApiVersion) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *TestApiVersion) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *TestApiVersion) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *TestApiService) VersionTest(ctx context.Context, in *TestReq, opts ...client.CallOption) (*TestApiOutput, error) {

	req := c.c.NewRequest(c.name, "TestApi.VersionTest", in)
	out := new(TestApiOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type TestApi_VersionTestService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type TestApiVersionTest struct {
	stream client.Stream
}

func (x *TestApiVersionTest) Close() error {
	return x.stream.Close()
}

func (x *TestApiVersionTest) Context() context.Context {
	return x.stream.Context()
}

func (x *TestApiVersionTest) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *TestApiVersionTest) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *TestApiVersionTest) Send(m *Message) error {
	return x.stream.Send(m)
}

// Server API for TestApi service
type TestApiHandler interface {
	Version(context.Context, *TestReq, *TestApiOutput) error
	VersionTest(context.Context, *TestReq, *TestApiOutput) error
}

func RegisterTestApiHandler(s server.Server, hdlr TestApiHandler, opts ...server.HandlerOption) error {
	type testApi interface {
		Version(ctx context.Context, in *TestReq, out *TestApiOutput) error
		VersionTest(ctx context.Context, in *TestReq, out *TestApiOutput) error
	}

	type TestApi struct {
		testApi
	}
	h := &testApiHandler{hdlr}
	opts = append(opts, server.EndpointMetadata("Version", map[string]string{"POST": "hello_test_api/version"}))
	opts = append(opts, server.EndpointMetadata("VersionTest", map[string]string{"POST": "/v1/example/versiontest"}))
	return s.Handle(s.NewHandler(&TestApi{h}, opts...))
}

type testApiHandler struct {
	TestApiHandler
}

func (h *TestApiHandler) Version(ctx context.Context, in *TestReq, out *TestApiOutput) error {
	return h.TestApiHandler.Version(ctx, in, out)
}

func (h *TestApiHandler) Version(ctx context.Context, stream server.Stream) error {

	m := new(TestReq)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TestApiHandler.Version(ctx, m, &testApiVersionStream{stream})

}

type TestApi_VersionStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type testApiVersionStream struct {
	stream server.Stream
}

func (x *testApiVersionStream) Close() error {
	return x.stream.Close()
}

func (x *testApiVersionStream) Context() context.Context {
	return x.stream.Context()
}

func (x *testApiVersionStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *testApiVersionStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *TestApiHandler) VersionTest(ctx context.Context, in *TestReq, out *TestApiOutput) error {
	return h.TestApiHandler.VersionTest(ctx, in, out)
}

func (h *TestApiHandler) VersionTest(ctx context.Context, stream server.Stream) error {

	m := new(TestReq)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TestApiHandler.VersionTest(ctx, m, &testApiVersionTestStream{stream})

}

type TestApi_VersionTestStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type testApiVersionTestStream struct {
	stream server.Stream
}

func (x *testApiVersionTestStream) Close() error {
	return x.stream.Close()
}

func (x *testApiVersionTestStream) Context() context.Context {
	return x.stream.Context()
}

func (x *testApiVersionTestStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *testApiVersionTestStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

// Client API for TestApiV2 service
type TestApiV2Service interface {
	Version(ctx context.Context, in *TestReq, opts ...client.CallOption) (*TestApiOutput, error)
	VersionTest(ctx context.Context, in *TestReq, opts ...client.CallOption) (*TestApiOutput, error)
}

type testApiV2Service struct {
	c    client.Client
	name string
}

func NewTestApiV2Service(name string, c client.Client) TestApiV2Service {
	return &testApiV2Service{
		c:    c,
		name: name,
	}
}
func (c *TestApiV2Service) Version(ctx context.Context, in *TestReq, opts ...client.CallOption) (*TestApiOutput, error) {

	req := c.c.NewRequest(c.name, "TestApiV2.Version", in)
	out := new(TestApiOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type TestApiV2_VersionService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type TestApiV2Version struct {
	stream client.Stream
}

func (x *TestApiV2Version) Close() error {
	return x.stream.Close()
}

func (x *TestApiV2Version) Context() context.Context {
	return x.stream.Context()
}

func (x *TestApiV2Version) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *TestApiV2Version) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *TestApiV2Version) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *TestApiV2Service) VersionTest(ctx context.Context, in *TestReq, opts ...client.CallOption) (*TestApiOutput, error) {

	req := c.c.NewRequest(c.name, "TestApiV2.VersionTest", in)
	out := new(TestApiOutput)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type TestApiV2_VersionTestService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type TestApiV2VersionTest struct {
	stream client.Stream
}

func (x *TestApiV2VersionTest) Close() error {
	return x.stream.Close()
}

func (x *TestApiV2VersionTest) Context() context.Context {
	return x.stream.Context()
}

func (x *TestApiV2VersionTest) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *TestApiV2VersionTest) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *TestApiV2VersionTest) Send(m *Message) error {
	return x.stream.Send(m)
}

// Server API for TestApiV2 service
type TestApiV2Handler interface {
	Version(context.Context, *TestReq, *TestApiOutput) error
	VersionTest(context.Context, *TestReq, *TestApiOutput) error
}

func RegisterTestApiV2Handler(s server.Server, hdlr TestApiV2Handler, opts ...server.HandlerOption) error {
	type testApiV2 interface {
		Version(ctx context.Context, in *TestReq, out *TestApiOutput) error
		VersionTest(ctx context.Context, in *TestReq, out *TestApiOutput) error
	}

	type TestApiV2 struct {
		testApiV2
	}
	h := &testApiV2Handler{hdlr}
	opts = append(opts, server.EndpointMetadata("Version", map[string]string{"POST": "/v2/example/version"}))
	opts = append(opts, server.EndpointMetadata("VersionTest", map[string]string{"POST": "/v2/example/versiontest"}))
	return s.Handle(s.NewHandler(&TestApiV2{h}, opts...))
}

type testApiV2Handler struct {
	TestApiV2Handler
}

func (h *TestApiV2Handler) Version(ctx context.Context, in *TestReq, out *TestApiOutput) error {
	return h.TestApiV2Handler.Version(ctx, in, out)
}

func (h *TestApiV2Handler) Version(ctx context.Context, stream server.Stream) error {

	m := new(TestReq)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TestApiV2Handler.Version(ctx, m, &testApiV2VersionStream{stream})

}

type TestApiV2_VersionStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type testApiV2VersionStream struct {
	stream server.Stream
}

func (x *testApiV2VersionStream) Close() error {
	return x.stream.Close()
}

func (x *testApiV2VersionStream) Context() context.Context {
	return x.stream.Context()
}

func (x *testApiV2VersionStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *testApiV2VersionStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *TestApiV2Handler) VersionTest(ctx context.Context, in *TestReq, out *TestApiOutput) error {
	return h.TestApiV2Handler.VersionTest(ctx, in, out)
}

func (h *TestApiV2Handler) VersionTest(ctx context.Context, stream server.Stream) error {

	m := new(TestReq)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TestApiV2Handler.VersionTest(ctx, m, &testApiV2VersionTestStream{stream})

}

type TestApiV2_VersionTestStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type testApiV2VersionTestStream struct {
	stream server.Stream
}

func (x *testApiV2VersionTestStream) Close() error {
	return x.stream.Close()
}

func (x *testApiV2VersionTestStream) Context() context.Context {
	return x.stream.Context()
}

func (x *testApiV2VersionTestStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *testApiV2VersionTestStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}
