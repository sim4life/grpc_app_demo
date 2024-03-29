// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domain/repository.proto

package domain

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

type Repository struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	IsPrivate            bool     `protobuf:"varint,4,opt,name=isPrivate,proto3" json:"isPrivate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf2810b2e6ad19c5, []int{0}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Repository) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Repository) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func init() {
	proto.RegisterType((*Repository)(nil), "domain.Repository")
}

func init() { proto.RegisterFile("domain/repository.proto", fileDescriptor_cf2810b2e6ad19c5) }

var fileDescriptor_cf2810b2e6ad19c5 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0x3f, 0xcf, 0x82, 0x30,
	0x1c, 0x84, 0x53, 0x20, 0xe4, 0xa5, 0xc3, 0x3b, 0x74, 0xd0, 0x0e, 0x0e, 0xc4, 0x89, 0x89, 0x0e,
	0x1a, 0xe3, 0xac, 0x93, 0x1b, 0xe9, 0xe8, 0x42, 0x0a, 0x2d, 0xf8, 0x4b, 0xe8, 0x9f, 0xb4, 0xc5,
	0xc4, 0x6f, 0x6f, 0x02, 0x46, 0xb7, 0xbb, 0x7b, 0x6e, 0x78, 0xf0, 0x56, 0x5a, 0x2d, 0xc0, 0x30,
	0xaf, 0x9c, 0x0d, 0x10, 0xad, 0x7f, 0xd5, 0xce, 0xdb, 0x68, 0x49, 0xbe, 0x82, 0xfd, 0x80, 0x31,
	0xff, 0x32, 0xf2, 0x8f, 0x13, 0x90, 0x14, 0x95, 0xa8, 0x4a, 0x79, 0x02, 0x92, 0x10, 0x9c, 0x19,
	0xa1, 0x15, 0x4d, 0x4a, 0x54, 0x15, 0x7c, 0xc9, 0x64, 0x83, 0xf3, 0x39, 0x28, 0x7f, 0x93, 0x34,
	0x5d, 0x7e, 0x9f, 0x46, 0x76, 0xb8, 0x80, 0xd0, 0x78, 0x78, 0x8a, 0xa8, 0x68, 0x56, 0xa2, 0xea,
	0x8f, 0xff, 0x86, 0xcb, 0xf9, 0x7e, 0x1a, 0x21, 0x3e, 0xe6, 0xae, 0xee, 0xad, 0x66, 0x01, 0xf4,
	0x71, 0x82, 0x41, 0xb1, 0xd1, 0xbb, 0xbe, 0x15, 0xce, 0xb5, 0x52, 0x69, 0xcb, 0xc0, 0x44, 0xe5,
	0x8d, 0x98, 0xd8, 0xc8, 0x9b, 0x2b, 0x5b, 0x0d, 0xbb, 0x7c, 0x11, 0x3e, 0xbc, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x81, 0xa4, 0x9c, 0xb1, 0xcb, 0x00, 0x00, 0x00,
}
