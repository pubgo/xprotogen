// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/hello/api.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("examples/proto/hello/api.proto", fileDescriptor_f19a39aeef8b3a71) }

var fileDescriptor_f19a39aeef8b3a71 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9,
	0xd7, 0x4f, 0x2c, 0xc8, 0xd4, 0x03, 0xf3, 0x85, 0x58, 0xc1, 0x02, 0x52, 0x32, 0xe9, 0xf9, 0xf9,
	0xe9, 0x39, 0xa9, 0x20, 0x09, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc,
	0x62, 0x88, 0x22, 0x29, 0x79, 0x5c, 0x86, 0x18, 0x42, 0x14, 0x18, 0xcd, 0x65, 0xe4, 0x62, 0x0f,
	0x49, 0x2d, 0x2e, 0x71, 0x2c, 0xc8, 0x14, 0x32, 0xe4, 0x62, 0x0f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc,
	0xcf, 0x13, 0xe2, 0xd3, 0x03, 0xab, 0xd4, 0x03, 0x49, 0x05, 0xa5, 0x16, 0x4a, 0x89, 0x20, 0xf1,
	0x1d, 0x0b, 0x32, 0xfd, 0x4b, 0x4b, 0x0a, 0x4a, 0x4b, 0x94, 0x18, 0x84, 0xe2, 0xb9, 0xb8, 0xa1,
	0x5a, 0x40, 0x32, 0x44, 0x6a, 0xd3, 0x6a, 0xba, 0xfc, 0x64, 0x32, 0x93, 0x8a, 0x15, 0xa3, 0x96,
	0x93, 0x3c, 0x17, 0x4b, 0x09, 0x48, 0xa3, 0xb8, 0x7e, 0x99, 0xa1, 0x3e, 0xd4, 0xad, 0xfa, 0x65,
	0x10, 0x13, 0x41, 0x12, 0x46, 0x5b, 0x19, 0xb9, 0x38, 0xa1, 0xba, 0xc3, 0x8c, 0x84, 0xfc, 0x49,
	0x75, 0xa1, 0x1c, 0xd8, 0x2a, 0x09, 0x25, 0x61, 0xfd, 0x32, 0x23, 0x74, 0xf3, 0xad, 0x18, 0xb5,
	0x84, 0xc2, 0xc9, 0x71, 0xbf, 0x12, 0xd8, 0x50, 0x19, 0x25, 0x71, 0x2c, 0x86, 0x82, 0x1c, 0x6d,
	0xc5, 0xa8, 0x95, 0xc4, 0x06, 0x0e, 0x5e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x7d,
	0x93, 0x94, 0xc6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestApiClient is the client API for TestApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestApiClient interface {
	Version(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
	VersionTest(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
}

type testApiClient struct {
	cc *grpc.ClientConn
}

func NewTestApiClient(cc *grpc.ClientConn) TestApiClient {
	return &testApiClient{cc}
}

func (c *testApiClient) Version(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApi/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) VersionTest(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApi/VersionTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestApiServer is the server API for TestApi service.
type TestApiServer interface {
	Version(context.Context, *TestReq) (*TestApiOutput, error)
	VersionTest(context.Context, *TestReq) (*TestApiOutput, error)
}

// UnimplementedTestApiServer can be embedded to have forward compatible implementations.
type UnimplementedTestApiServer struct {
}

func (*UnimplementedTestApiServer) Version(ctx context.Context, req *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedTestApiServer) VersionTest(ctx context.Context, req *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VersionTest not implemented")
}

func RegisterTestApiServer(s *grpc.Server, srv TestApiServer) {
	s.RegisterService(&_TestApi_serviceDesc, srv)
}

func _TestApi_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApi/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).Version(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_VersionTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).VersionTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApi/VersionTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).VersionTest(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.TestApi",
	HandlerType: (*TestApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _TestApi_Version_Handler,
		},
		{
			MethodName: "VersionTest",
			Handler:    _TestApi_VersionTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/proto/hello/api.proto",
}

// TestApiV2Client is the client API for TestApiV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestApiV2Client interface {
	Version(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
	VersionTest(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error)
}

type testApiV2Client struct {
	cc *grpc.ClientConn
}

func NewTestApiV2Client(cc *grpc.ClientConn) TestApiV2Client {
	return &testApiV2Client{cc}
}

func (c *testApiV2Client) Version(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApiV2/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiV2Client) VersionTest(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestApiOutput, error) {
	out := new(TestApiOutput)
	err := c.cc.Invoke(ctx, "/hello.TestApiV2/VersionTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestApiV2Server is the server API for TestApiV2 service.
type TestApiV2Server interface {
	Version(context.Context, *TestReq) (*TestApiOutput, error)
	VersionTest(context.Context, *TestReq) (*TestApiOutput, error)
}

// UnimplementedTestApiV2Server can be embedded to have forward compatible implementations.
type UnimplementedTestApiV2Server struct {
}

func (*UnimplementedTestApiV2Server) Version(ctx context.Context, req *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedTestApiV2Server) VersionTest(ctx context.Context, req *TestReq) (*TestApiOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VersionTest not implemented")
}

func RegisterTestApiV2Server(s *grpc.Server, srv TestApiV2Server) {
	s.RegisterService(&_TestApiV2_serviceDesc, srv)
}

func _TestApiV2_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiV2Server).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApiV2/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiV2Server).Version(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApiV2_VersionTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiV2Server).VersionTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.TestApiV2/VersionTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiV2Server).VersionTest(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestApiV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.TestApiV2",
	HandlerType: (*TestApiV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _TestApiV2_Version_Handler,
		},
		{
			MethodName: "VersionTest",
			Handler:    _TestApiV2_VersionTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/proto/hello/api.proto",
}
