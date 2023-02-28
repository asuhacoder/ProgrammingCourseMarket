// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/pkg/pb/auth/auth.proto

package auth

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

type AuthnRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthnRequest) Reset()         { *m = AuthnRequest{} }
func (m *AuthnRequest) String() string { return proto.CompactTextString(m) }
func (*AuthnRequest) ProtoMessage()    {}
func (*AuthnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ef67c2da09a44de, []int{0}
}

func (m *AuthnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthnRequest.Unmarshal(m, b)
}
func (m *AuthnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthnRequest.Marshal(b, m, deterministic)
}
func (m *AuthnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthnRequest.Merge(m, src)
}
func (m *AuthnRequest) XXX_Size() int {
	return xxx_messageInfo_AuthnRequest.Size(m)
}
func (m *AuthnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthnRequest proto.InternalMessageInfo

func (m *AuthnRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthnRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthnReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Introduction         string   `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Permission           string   `protobuf:"bytes,6,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthnReply) Reset()         { *m = AuthnReply{} }
func (m *AuthnReply) String() string { return proto.CompactTextString(m) }
func (*AuthnReply) ProtoMessage()    {}
func (*AuthnReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ef67c2da09a44de, []int{1}
}

func (m *AuthnReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthnReply.Unmarshal(m, b)
}
func (m *AuthnReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthnReply.Marshal(b, m, deterministic)
}
func (m *AuthnReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthnReply.Merge(m, src)
}
func (m *AuthnReply) XXX_Size() int {
	return xxx_messageInfo_AuthnReply.Size(m)
}
func (m *AuthnReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthnReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthnReply proto.InternalMessageInfo

func (m *AuthnReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthnReply) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *AuthnReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthnReply) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *AuthnReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthnReply) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

type AuthzRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthzRequest) Reset()         { *m = AuthzRequest{} }
func (m *AuthzRequest) String() string { return proto.CompactTextString(m) }
func (*AuthzRequest) ProtoMessage()    {}
func (*AuthzRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ef67c2da09a44de, []int{2}
}

func (m *AuthzRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthzRequest.Unmarshal(m, b)
}
func (m *AuthzRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthzRequest.Marshal(b, m, deterministic)
}
func (m *AuthzRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthzRequest.Merge(m, src)
}
func (m *AuthzRequest) XXX_Size() int {
	return xxx_messageInfo_AuthzRequest.Size(m)
}
func (m *AuthzRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthzRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthzRequest proto.InternalMessageInfo

func (m *AuthzRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthzReply struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Introduction         string   `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Permission           string   `protobuf:"bytes,6,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthzReply) Reset()         { *m = AuthzReply{} }
func (m *AuthzReply) String() string { return proto.CompactTextString(m) }
func (*AuthzReply) ProtoMessage()    {}
func (*AuthzReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ef67c2da09a44de, []int{3}
}

func (m *AuthzReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthzReply.Unmarshal(m, b)
}
func (m *AuthzReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthzReply.Marshal(b, m, deterministic)
}
func (m *AuthzReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthzReply.Merge(m, src)
}
func (m *AuthzReply) XXX_Size() int {
	return xxx_messageInfo_AuthzReply.Size(m)
}
func (m *AuthzReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthzReply.DiscardUnknown(m)
}

var xxx_messageInfo_AuthzReply proto.InternalMessageInfo

func (m *AuthzReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthzReply) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *AuthzReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthzReply) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *AuthzReply) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthzReply) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthnRequest)(nil), "auth.AuthnRequest")
	proto.RegisterType((*AuthnReply)(nil), "auth.AuthnReply")
	proto.RegisterType((*AuthzRequest)(nil), "auth.AuthzRequest")
	proto.RegisterType((*AuthzReply)(nil), "auth.AuthzReply")
}

func init() {
	proto.RegisterFile("internal/pkg/pb/auth/auth.proto", fileDescriptor_8ef67c2da09a44de)
}

var fileDescriptor_8ef67c2da09a44de = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x24, 0x90, 0x54, 0xb0, 0xea, 0x01, 0x59, 0x1c, 0xa2, 0x1e, 0x00, 0x45, 0x1c, 0xb8, 0x10,
	0x4b, 0xf0, 0x01, 0x10, 0x38, 0x23, 0xa1, 0x1e, 0xb9, 0x39, 0xad, 0x95, 0x58, 0x89, 0x1f, 0xf8,
	0x21, 0xd4, 0x7c, 0x0f, 0x1f, 0x8a, 0xec, 0x34, 0xc5, 0x15, 0xfd, 0x80, 0x5e, 0xac, 0xdd, 0x59,
	0xcd, 0x68, 0xc6, 0xbb, 0x70, 0xc3, 0x84, 0xa5, 0x5a, 0x90, 0x1e, 0xab, 0xae, 0xc1, 0xaa, 0xc6,
	0xc4, 0xd9, 0x36, 0x3c, 0xa5, 0xd2, 0xd2, 0x4a, 0x94, 0xfa, 0xba, 0x78, 0x81, 0x79, 0xe5, 0x6c,
	0x2b, 0x96, 0xf4, 0xcb, 0x51, 0x63, 0xd1, 0x15, 0x64, 0x94, 0x13, 0xd6, 0xe7, 0xc9, 0x6d, 0x72,
	0x7f, 0xb1, 0x1c, 0x1b, 0xb4, 0x80, 0x73, 0x45, 0x8c, 0xf9, 0x96, 0x7a, 0x9d, 0x9f, 0x86, 0xc1,
	0xae, 0x2f, 0x7e, 0x12, 0x80, 0xad, 0x84, 0xea, 0x37, 0x5e, 0xc0, 0xca, 0x8e, 0x8a, 0x49, 0x20,
	0x34, 0x08, 0x41, 0xea, 0x1c, 0x9b, 0xc8, 0xa1, 0xf6, 0x98, 0x20, 0x9c, 0xe6, 0x67, 0x23, 0xe6,
	0x6b, 0x54, 0xc0, 0x9c, 0x09, 0xab, 0xe5, 0xda, 0xad, 0x2c, 0x93, 0x22, 0x4f, 0xc3, 0x6c, 0x0f,
	0xfb, 0xb3, 0x98, 0xc5, 0x16, 0xaf, 0x01, 0x14, 0xd5, 0x9c, 0x19, 0xe3, 0x79, 0xb3, 0x30, 0x8a,
	0x90, 0xe2, 0x6e, 0x0c, 0x3a, 0x44, 0x41, 0xff, 0xfb, 0xdc, 0x85, 0x19, 0x8e, 0x3a, 0xcc, 0x63,
	0x0b, 0xa9, 0x77, 0x89, 0x30, 0x64, 0xe1, 0xeb, 0x11, 0x2a, 0xc3, 0x66, 0xe3, 0x55, 0x2e, 0x2e,
	0xf7, 0x30, 0xd5, 0x6f, 0x8a, 0x93, 0x89, 0x30, 0xc4, 0x84, 0xe1, 0x00, 0x61, 0xd8, 0x12, 0x5e,
	0xab, 0xcf, 0xe7, 0x86, 0xd9, 0xd6, 0xd5, 0xe5, 0x4a, 0x72, 0x5c, 0x19, 0xd7, 0x92, 0x07, 0x82,
	0x3f, 0xb4, 0x6c, 0x34, 0xe1, 0x9c, 0x89, 0xe6, 0x4d, 0x3a, 0x6d, 0xe8, 0x3b, 0xd1, 0x1d, 0xb5,
	0xf8, 0xd0, 0xcd, 0xd5, 0xb3, 0x70, 0x6f, 0x4f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xc4,
	0x92, 0x30, 0x92, 0x02, 0x00, 0x00,
}
