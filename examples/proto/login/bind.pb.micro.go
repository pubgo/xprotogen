// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: examples/proto/login/bind.proto
package login

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

// Client API for BindTelephone service
type BindTelephoneService interface {
	Check(ctx context.Context, in *CheckRequest, opts ...client.CallOption) (*CheckResponse, error)
	BindVerify(ctx context.Context, in *BindVerifyRequest, opts ...client.CallOption) (*BindVerifyResponse, error)
	BindChange(ctx context.Context, in *BindChangeRequest, opts ...client.CallOption) (*BindChangeResponse, error)
	AutomaticBind(ctx context.Context, in *AutomaticBindRequest, opts ...client.CallOption) (*AutomaticBindResponse, error)
	BindPhoneParse(ctx context.Context, in *BindPhoneParseRequest, opts ...client.CallOption) (*BindPhoneParseResponse, error)
	BindPhoneParseByOneClick(ctx context.Context, in *BindPhoneParseByOneClickRequest, opts ...client.CallOption) (*BindPhoneParseByOneClickResponse, error)
}

type bindTelephoneService struct {
	c    client.Client
	name string
}

func NewBindTelephoneService(name string, c client.Client) BindTelephoneService {
	return &bindTelephoneService{
		c:    c,
		name: name,
	}
}
func (c *BindTelephoneService) Check(ctx context.Context, in *CheckRequest, opts ...client.CallOption) (*CheckResponse, error) {

	req := c.c.NewRequest(c.name, "BindTelephone.Check", in)
	out := new(CheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type BindTelephone_CheckService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type BindTelephoneCheck struct {
	stream client.Stream
}

func (x *BindTelephoneCheck) Close() error {
	return x.stream.Close()
}

func (x *BindTelephoneCheck) Context() context.Context {
	return x.stream.Context()
}

func (x *BindTelephoneCheck) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *BindTelephoneCheck) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *BindTelephoneCheck) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *BindTelephoneService) BindVerify(ctx context.Context, in *BindVerifyRequest, opts ...client.CallOption) (*BindVerifyResponse, error) {

	req := c.c.NewRequest(c.name, "BindTelephone.BindVerify", in)
	out := new(BindVerifyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type BindTelephone_BindVerifyService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type BindTelephoneBindVerify struct {
	stream client.Stream
}

func (x *BindTelephoneBindVerify) Close() error {
	return x.stream.Close()
}

func (x *BindTelephoneBindVerify) Context() context.Context {
	return x.stream.Context()
}

func (x *BindTelephoneBindVerify) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *BindTelephoneBindVerify) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *BindTelephoneBindVerify) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *BindTelephoneService) BindChange(ctx context.Context, in *BindChangeRequest, opts ...client.CallOption) (*BindChangeResponse, error) {

	req := c.c.NewRequest(c.name, "BindTelephone.BindChange", in)
	out := new(BindChangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type BindTelephone_BindChangeService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type BindTelephoneBindChange struct {
	stream client.Stream
}

func (x *BindTelephoneBindChange) Close() error {
	return x.stream.Close()
}

func (x *BindTelephoneBindChange) Context() context.Context {
	return x.stream.Context()
}

func (x *BindTelephoneBindChange) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *BindTelephoneBindChange) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *BindTelephoneBindChange) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *BindTelephoneService) AutomaticBind(ctx context.Context, in *AutomaticBindRequest, opts ...client.CallOption) (*AutomaticBindResponse, error) {

	req := c.c.NewRequest(c.name, "BindTelephone.AutomaticBind", in)
	out := new(AutomaticBindResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type BindTelephone_AutomaticBindService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type BindTelephoneAutomaticBind struct {
	stream client.Stream
}

func (x *BindTelephoneAutomaticBind) Close() error {
	return x.stream.Close()
}

func (x *BindTelephoneAutomaticBind) Context() context.Context {
	return x.stream.Context()
}

func (x *BindTelephoneAutomaticBind) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *BindTelephoneAutomaticBind) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *BindTelephoneAutomaticBind) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *BindTelephoneService) BindPhoneParse(ctx context.Context, in *BindPhoneParseRequest, opts ...client.CallOption) (*BindPhoneParseResponse, error) {

	req := c.c.NewRequest(c.name, "BindTelephone.BindPhoneParse", in)
	out := new(BindPhoneParseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type BindTelephone_BindPhoneParseService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type BindTelephoneBindPhoneParse struct {
	stream client.Stream
}

func (x *BindTelephoneBindPhoneParse) Close() error {
	return x.stream.Close()
}

func (x *BindTelephoneBindPhoneParse) Context() context.Context {
	return x.stream.Context()
}

func (x *BindTelephoneBindPhoneParse) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *BindTelephoneBindPhoneParse) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *BindTelephoneBindPhoneParse) Send(m *Message) error {
	return x.stream.Send(m)
}

func (c *BindTelephoneService) BindPhoneParseByOneClick(ctx context.Context, in *BindPhoneParseByOneClickRequest, opts ...client.CallOption) (*BindPhoneParseByOneClickResponse, error) {

	req := c.c.NewRequest(c.name, "BindTelephone.BindPhoneParseByOneClick", in)
	out := new(BindPhoneParseByOneClickResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Stream auxiliary types and methods.
type BindTelephone_BindPhoneParseByOneClickService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}
type BindTelephoneBindPhoneParseByOneClick struct {
	stream client.Stream
}

func (x *BindTelephoneBindPhoneParseByOneClick) Close() error {
	return x.stream.Close()
}

func (x *BindTelephoneBindPhoneParseByOneClick) Context() context.Context {
	return x.stream.Context()
}

func (x *BindTelephoneBindPhoneParseByOneClick) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *BindTelephoneBindPhoneParseByOneClick) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *BindTelephoneBindPhoneParseByOneClick) Send(m *Message) error {
	return x.stream.Send(m)
}

// Server API for BindTelephone service
type BindTelephoneHandler interface {
	Check(context.Context, *CheckRequest, *CheckResponse) error
	BindVerify(context.Context, *BindVerifyRequest, *BindVerifyResponse) error
	BindChange(context.Context, *BindChangeRequest, *BindChangeResponse) error
	AutomaticBind(context.Context, *AutomaticBindRequest, *AutomaticBindResponse) error
	BindPhoneParse(context.Context, *BindPhoneParseRequest, *BindPhoneParseResponse) error
	BindPhoneParseByOneClick(context.Context, *BindPhoneParseByOneClickRequest, *BindPhoneParseByOneClickResponse) error
}

func RegisterBindTelephoneHandler(s server.Server, hdlr BindTelephoneHandler, opts ...server.HandlerOption) error {
	type bindTelephone interface {
		Check(ctx context.Context, in *CheckRequest, out *CheckResponse) error
		BindVerify(ctx context.Context, in *BindVerifyRequest, out *BindVerifyResponse) error
		BindChange(ctx context.Context, in *BindChangeRequest, out *BindChangeResponse) error
		AutomaticBind(ctx context.Context, in *AutomaticBindRequest, out *AutomaticBindResponse) error
		BindPhoneParse(ctx context.Context, in *BindPhoneParseRequest, out *BindPhoneParseResponse) error
		BindPhoneParseByOneClick(ctx context.Context, in *BindPhoneParseByOneClickRequest, out *BindPhoneParseByOneClickResponse) error
	}

	type BindTelephone struct {
		bindTelephone
	}
	h := &bindTelephoneHandler{hdlr}
	return s.Handle(s.NewHandler(&BindTelephone{h}, opts...))
}

type bindTelephoneHandler struct {
	BindTelephoneHandler
}

func (h *BindTelephoneHandler) Check(ctx context.Context, in *CheckRequest, out *CheckResponse) error {
	return h.BindTelephoneHandler.Check(ctx, in, out)
}

func (h *BindTelephoneHandler) Check(ctx context.Context, stream server.Stream) error {

	m := new(CheckRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BindTelephoneHandler.Check(ctx, m, &bindTelephoneCheckStream{stream})

}

type BindTelephone_CheckStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type bindTelephoneCheckStream struct {
	stream server.Stream
}

func (x *bindTelephoneCheckStream) Close() error {
	return x.stream.Close()
}

func (x *bindTelephoneCheckStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bindTelephoneCheckStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bindTelephoneCheckStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *BindTelephoneHandler) BindVerify(ctx context.Context, in *BindVerifyRequest, out *BindVerifyResponse) error {
	return h.BindTelephoneHandler.BindVerify(ctx, in, out)
}

func (h *BindTelephoneHandler) BindVerify(ctx context.Context, stream server.Stream) error {

	m := new(BindVerifyRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BindTelephoneHandler.BindVerify(ctx, m, &bindTelephoneBindVerifyStream{stream})

}

type BindTelephone_BindVerifyStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type bindTelephoneBindVerifyStream struct {
	stream server.Stream
}

func (x *bindTelephoneBindVerifyStream) Close() error {
	return x.stream.Close()
}

func (x *bindTelephoneBindVerifyStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bindTelephoneBindVerifyStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bindTelephoneBindVerifyStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *BindTelephoneHandler) BindChange(ctx context.Context, in *BindChangeRequest, out *BindChangeResponse) error {
	return h.BindTelephoneHandler.BindChange(ctx, in, out)
}

func (h *BindTelephoneHandler) BindChange(ctx context.Context, stream server.Stream) error {

	m := new(BindChangeRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BindTelephoneHandler.BindChange(ctx, m, &bindTelephoneBindChangeStream{stream})

}

type BindTelephone_BindChangeStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type bindTelephoneBindChangeStream struct {
	stream server.Stream
}

func (x *bindTelephoneBindChangeStream) Close() error {
	return x.stream.Close()
}

func (x *bindTelephoneBindChangeStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bindTelephoneBindChangeStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bindTelephoneBindChangeStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *BindTelephoneHandler) AutomaticBind(ctx context.Context, in *AutomaticBindRequest, out *AutomaticBindResponse) error {
	return h.BindTelephoneHandler.AutomaticBind(ctx, in, out)
}

func (h *BindTelephoneHandler) AutomaticBind(ctx context.Context, stream server.Stream) error {

	m := new(AutomaticBindRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BindTelephoneHandler.AutomaticBind(ctx, m, &bindTelephoneAutomaticBindStream{stream})

}

type BindTelephone_AutomaticBindStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type bindTelephoneAutomaticBindStream struct {
	stream server.Stream
}

func (x *bindTelephoneAutomaticBindStream) Close() error {
	return x.stream.Close()
}

func (x *bindTelephoneAutomaticBindStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bindTelephoneAutomaticBindStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bindTelephoneAutomaticBindStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *BindTelephoneHandler) BindPhoneParse(ctx context.Context, in *BindPhoneParseRequest, out *BindPhoneParseResponse) error {
	return h.BindTelephoneHandler.BindPhoneParse(ctx, in, out)
}

func (h *BindTelephoneHandler) BindPhoneParse(ctx context.Context, stream server.Stream) error {

	m := new(BindPhoneParseRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BindTelephoneHandler.BindPhoneParse(ctx, m, &bindTelephoneBindPhoneParseStream{stream})

}

type BindTelephone_BindPhoneParseStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type bindTelephoneBindPhoneParseStream struct {
	stream server.Stream
}

func (x *bindTelephoneBindPhoneParseStream) Close() error {
	return x.stream.Close()
}

func (x *bindTelephoneBindPhoneParseStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bindTelephoneBindPhoneParseStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bindTelephoneBindPhoneParseStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (h *BindTelephoneHandler) BindPhoneParseByOneClick(ctx context.Context, in *BindPhoneParseByOneClickRequest, out *BindPhoneParseByOneClickResponse) error {
	return h.BindTelephoneHandler.BindPhoneParseByOneClick(ctx, in, out)
}

func (h *BindTelephoneHandler) BindPhoneParseByOneClick(ctx context.Context, stream server.Stream) error {

	m := new(BindPhoneParseByOneClickRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.BindTelephoneHandler.BindPhoneParseByOneClick(ctx, m, &bindTelephoneBindPhoneParseByOneClickStream{stream})

}

type BindTelephone_BindPhoneParseByOneClickStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
}

type bindTelephoneBindPhoneParseByOneClickStream struct {
	stream server.Stream
}

func (x *bindTelephoneBindPhoneParseByOneClickStream) Close() error {
	return x.stream.Close()
}

func (x *bindTelephoneBindPhoneParseByOneClickStream) Context() context.Context {
	return x.stream.Context()
}

func (x *bindTelephoneBindPhoneParseByOneClickStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *bindTelephoneBindPhoneParseByOneClickStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

// POST /user/bind-telephone/check
// POST /user/bind-telephone/bind-verify
// POST /user/bind-telephone/bind-change
// POST /user/bind-telephone/automatic-bind
// POST /user/bind-telephone/bind-phone-parse
// POST /user/bind-telephone/bind-phone-parse-by-one-click
