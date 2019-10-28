// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/account.proto

package go_micro_srv_account

import (
	fmt "fmt"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Request struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{3}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

type InfoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{4}
}

func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse.Unmarshal(m, b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return xxx_messageInfo_InfoResponse.Size(m)
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoResponse) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "go.micro.srv.account.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "go.micro.srv.account.LoginResponse")
	proto.RegisterType((*Request)(nil), "go.micro.srv.account.Request")
	proto.RegisterType((*LogoutResponse)(nil), "go.micro.srv.account.LogoutResponse")
	proto.RegisterType((*InfoResponse)(nil), "go.micro.srv.account.InfoResponse")
}

func init() { proto.RegisterFile("proto/account/account.proto", fileDescriptor_5d492a0187472a3b) }

var fileDescriptor_5d492a0187472a3b = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xdb, 0xb2, 0x75, 0xfa, 0x98, 0x43, 0x1e, 0x43, 0xe6, 0x44, 0x90, 0xa8, 0xe0, 0x29,
	0x82, 0xde, 0xbc, 0x79, 0x11, 0xc4, 0x81, 0xd0, 0xff, 0x20, 0xb6, 0xb1, 0x14, 0x59, 0x5e, 0x4d,
	0xd2, 0xf9, 0xc7, 0x7b, 0x91, 0x35, 0xc9, 0xc8, 0x61, 0x75, 0xa7, 0xf6, 0xcb, 0x97, 0xfc, 0xf2,
	0x7d, 0x2f, 0x70, 0xd1, 0x6a, 0xb2, 0x74, 0x2f, 0xca, 0x92, 0x3a, 0x65, 0xc3, 0x97, 0xf7, 0xab,
	0x38, 0xaf, 0x89, 0xaf, 0x9b, 0x52, 0x13, 0x37, 0x7a, 0xc3, 0xbd, 0xc7, 0x5e, 0x60, 0xba, 0xa2,
	0xba, 0x51, 0x85, 0xfc, 0xee, 0xa4, 0xb1, 0xb8, 0x84, 0xa3, 0xce, 0x48, 0xad, 0xc4, 0x5a, 0x2e,
	0xd2, 0xab, 0xf4, 0xee, 0xb8, 0xd8, 0xe9, 0xad, 0xd7, 0x0a, 0x63, 0x7e, 0x48, 0x57, 0x8b, 0xcc,
	0x79, 0x41, 0xb3, 0x5b, 0x38, 0xf1, 0x1c, 0xd3, 0x92, 0x32, 0x12, 0xe7, 0x30, 0xb6, 0xf4, 0x25,
	0x95, 0xa7, 0x38, 0xc1, 0xce, 0x61, 0x12, 0x6e, 0x9a, 0x41, 0xd6, 0x54, 0xde, 0xcd, 0x9a, 0x8a,
	0x9d, 0xc2, 0x6c, 0x45, 0x35, 0x75, 0x36, 0x20, 0xd8, 0x13, 0x4c, 0x5f, 0xd5, 0x27, 0xed, 0x90,
	0x08, 0xa3, 0x28, 0x57, 0xff, 0x8f, 0x67, 0x90, 0x8b, 0x8d, 0xb0, 0x42, 0xfb, 0x44, 0x5e, 0x3d,
	0xfc, 0xa6, 0x30, 0x79, 0x76, 0x1d, 0xb1, 0x80, 0x71, 0x9f, 0x0d, 0x19, 0xdf, 0x37, 0x03, 0x1e,
	0x0f, 0x60, 0x79, 0xfd, 0xef, 0x1e, 0x9f, 0x2c, 0xc1, 0x77, 0xc8, 0x5d, 0x5a, 0xbc, 0xdc, 0x7f,
	0x20, 0xf0, 0x6e, 0x06, 0x79, 0x71, 0xd5, 0x04, 0xdf, 0x60, 0xb4, 0x2d, 0x7b, 0x08, 0x37, 0x50,
	0x21, 0x9e, 0x13, 0x4b, 0x3e, 0xf2, 0xfe, 0xc9, 0x1f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba,
	0xa6, 0xc1, 0x3f, 0x11, 0x02, 0x00, 0x00,
}
