// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/hello/api1.proto

package hello

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type TestReq struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReq) Reset()         { *m = TestReq{} }
func (m *TestReq) String() string { return proto.CompactTextString(m) }
func (*TestReq) ProtoMessage()    {}
func (*TestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b1bf68f1d872202, []int{0}
}

func (m *TestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReq.Unmarshal(m, b)
}
func (m *TestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReq.Marshal(b, m, deterministic)
}
func (m *TestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReq.Merge(m, src)
}
func (m *TestReq) XXX_Size() int {
	return xxx_messageInfo_TestReq.Size(m)
}
func (m *TestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReq.DiscardUnknown(m)
}

var xxx_messageInfo_TestReq proto.InternalMessageInfo

func (m *TestReq) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

type TestApiData struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	SrvVersion           string   `protobuf:"bytes,2,opt,name=srvVersion,proto3" json:"srvVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestApiData) Reset()         { *m = TestApiData{} }
func (m *TestApiData) String() string { return proto.CompactTextString(m) }
func (*TestApiData) ProtoMessage()    {}
func (*TestApiData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b1bf68f1d872202, []int{1}
}

func (m *TestApiData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestApiData.Unmarshal(m, b)
}
func (m *TestApiData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestApiData.Marshal(b, m, deterministic)
}
func (m *TestApiData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestApiData.Merge(m, src)
}
func (m *TestApiData) XXX_Size() int {
	return xxx_messageInfo_TestApiData.Size(m)
}
func (m *TestApiData) XXX_DiscardUnknown() {
	xxx_messageInfo_TestApiData.DiscardUnknown(m)
}

var xxx_messageInfo_TestApiData proto.InternalMessageInfo

func (m *TestApiData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *TestApiData) GetSrvVersion() string {
	if m != nil {
		return m.SrvVersion
	}
	return ""
}

type TestApiOutput struct {
	Code                 int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	NowTime              int64        `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	Data                 *TestApiData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TestApiOutput) Reset()         { *m = TestApiOutput{} }
func (m *TestApiOutput) String() string { return proto.CompactTextString(m) }
func (*TestApiOutput) ProtoMessage()    {}
func (*TestApiOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b1bf68f1d872202, []int{2}
}

func (m *TestApiOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestApiOutput.Unmarshal(m, b)
}
func (m *TestApiOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestApiOutput.Marshal(b, m, deterministic)
}
func (m *TestApiOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestApiOutput.Merge(m, src)
}
func (m *TestApiOutput) XXX_Size() int {
	return xxx_messageInfo_TestApiOutput.Size(m)
}
func (m *TestApiOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TestApiOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TestApiOutput proto.InternalMessageInfo

func (m *TestApiOutput) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TestApiOutput) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TestApiOutput) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *TestApiOutput) GetData() *TestApiData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TestReq)(nil), "hello.TestReq")
	proto.RegisterType((*TestApiData)(nil), "hello.TestApiData")
	proto.RegisterType((*TestApiOutput)(nil), "hello.TestApiOutput")
}

func init() { proto.RegisterFile("examples/proto/hello/api1.proto", fileDescriptor_8b1bf68f1d872202) }

var fileDescriptor_8b1bf68f1d872202 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0xdd, 0x6a, 0x83, 0x30,
	0x18, 0xc5, 0x55, 0xeb, 0xfa, 0xc9, 0x6e, 0x72, 0xe5, 0xc6, 0x20, 0x22, 0x1b, 0x48, 0x61, 0xba,
	0x9f, 0x27, 0xa8, 0xec, 0x7e, 0x90, 0x95, 0xde, 0x8e, 0xd8, 0x66, 0x36, 0xa0, 0x26, 0xd3, 0xd8,
	0xed, 0x75, 0xf6, 0x62, 0x3e, 0x80, 0x4f, 0x31, 0x92, 0x28, 0xeb, 0xcd, 0xf1, 0x3b, 0xe7, 0x3b,
	0xc7, 0x43, 0x3e, 0xc0, 0xec, 0x87, 0xd6, 0xb2, 0x62, 0x5d, 0x26, 0x5b, 0xa1, 0x44, 0x76, 0x64,
	0x55, 0x25, 0x32, 0x2a, 0xf9, 0x53, 0x6a, 0x04, 0xe4, 0x19, 0xe5, 0xe6, 0xa1, 0xe4, 0xea, 0xd8,
	0x17, 0xe9, 0x5e, 0xd4, 0x59, 0x29, 0x4a, 0x61, 0xed, 0x45, 0xff, 0x69, 0x98, 0xcd, 0xea, 0xc9,
	0xa6, 0xe2, 0x35, 0xf8, 0x5b, 0xd6, 0x29, 0xc2, 0xbe, 0x10, 0x06, 0x8f, 0x37, 0xb2, 0x57, 0xa1,
	0x13, 0x39, 0xc9, 0x2a, 0x5f, 0x8d, 0x03, 0xb6, 0x02, 0xb1, 0x9f, 0xf8, 0x03, 0x02, 0xed, 0xdd,
	0x48, 0xfe, 0x4a, 0x15, 0x45, 0x77, 0xe0, 0x9f, 0x58, 0xdb, 0x71, 0xd1, 0x4c, 0x09, 0x18, 0x07,
	0xbc, 0xdc, 0x48, 0xbe, 0x63, 0x2d, 0x99, 0x57, 0x68, 0x0d, 0xd0, 0xb5, 0xa7, 0xdd, 0x64, 0xbc,
	0xf8, 0x37, 0xbe, 0x1b, 0x95, 0x9c, 0x6d, 0xe3, 0x5f, 0x07, 0xae, 0xa6, 0x86, 0xb7, 0x5e, 0xc9,
	0x5e, 0xa1, 0x5b, 0x70, 0xf7, 0xe2, 0xc0, 0x4c, 0x81, 0x97, 0x5f, 0x8e, 0x03, 0x36, 0x9c, 0x18,
	0x44, 0xd7, 0xb0, 0xa8, 0xbb, 0x72, 0xfa, 0xa9, 0x3f, 0x0e, 0x58, 0x53, 0xa2, 0x01, 0xdd, 0x83,
	0xdf, 0x88, 0xef, 0x2d, 0xaf, 0x59, 0xb8, 0x88, 0x9c, 0x64, 0x91, 0x07, 0xe3, 0x80, 0x67, 0x89,
	0xcc, 0x03, 0x7a, 0x04, 0xf7, 0x40, 0x15, 0x0d, 0xdd, 0xc8, 0x49, 0x82, 0x67, 0x94, 0x9a, 0x1b,
	0xa6, 0x67, 0xaf, 0xb4, 0x9d, 0xda, 0x43, 0x0c, 0x16, 0x4b, 0x73, 0xb7, 0x97, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x97, 0xfe, 0xa7, 0xb4, 0x90, 0x01, 0x00, 0x00,
}
