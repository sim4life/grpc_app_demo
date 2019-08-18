// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/repository-service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	domain "github.com/sim4life/grpc_app_demo/internal/gRPC/domain"
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

type AddRepositoryResponse struct {
	AddedRepository      *domain.Repository `protobuf:"bytes,1,opt,name=addedRepository,proto3" json:"addedRepository,omitempty"`
	Error                *Error             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddRepositoryResponse) Reset()         { *m = AddRepositoryResponse{} }
func (m *AddRepositoryResponse) String() string { return proto.CompactTextString(m) }
func (*AddRepositoryResponse) ProtoMessage()    {}
func (*AddRepositoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d64b488d017c1a6, []int{0}
}

func (m *AddRepositoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRepositoryResponse.Unmarshal(m, b)
}
func (m *AddRepositoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRepositoryResponse.Marshal(b, m, deterministic)
}
func (m *AddRepositoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRepositoryResponse.Merge(m, src)
}
func (m *AddRepositoryResponse) XXX_Size() int {
	return xxx_messageInfo_AddRepositoryResponse.Size(m)
}
func (m *AddRepositoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRepositoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddRepositoryResponse proto.InternalMessageInfo

func (m *AddRepositoryResponse) GetAddedRepository() *domain.Repository {
	if m != nil {
		return m.AddedRepository
	}
	return nil
}

func (m *AddRepositoryResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d64b488d017c1a6, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*AddRepositoryResponse)(nil), "service.AddRepositoryResponse")
	proto.RegisterType((*Error)(nil), "service.Error")
}

func init() { proto.RegisterFile("service/repository-service.proto", fileDescriptor_1d64b488d017c1a6) }

var fileDescriptor_1d64b488d017c1a6 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0xa5, 0x6a, 0x2d, 0x8d, 0xa0, 0x38, 0x20, 0x2e, 0x3d, 0x48, 0x29, 0x1e, 0xbc, 0xb8, 0x81,
	0xaa, 0x48, 0xc1, 0x8b, 0x8a, 0x57, 0x91, 0x78, 0xf3, 0x52, 0xd2, 0x9d, 0x71, 0x0d, 0x34, 0x3b,
	0x61, 0xb2, 0x0a, 0xe2, 0xcf, 0x8b, 0xd9, 0x6d, 0x5d, 0x4a, 0x6f, 0x79, 0x6f, 0xde, 0xbc, 0xf7,
	0x32, 0x6a, 0x1c, 0x49, 0xbe, 0x5c, 0x41, 0x5a, 0x28, 0x70, 0x74, 0x35, 0xcb, 0xf7, 0x65, 0x4b,
	0xe5, 0x41, 0xb8, 0x66, 0x18, 0xb4, 0x70, 0x74, 0x8a, 0xec, 0xad, 0xab, 0x3a, 0xca, 0x46, 0x31,
	0xf9, 0x51, 0x27, 0xf7, 0x88, 0x66, 0x4d, 0x1b, 0x8a, 0x81, 0xab, 0x48, 0x70, 0xa7, 0x8e, 0x2c,
	0x22, 0x75, 0x46, 0x59, 0x6f, 0xdc, 0xbb, 0x38, 0x98, 0x42, 0xde, 0x78, 0xe5, 0x9d, 0xa5, 0x4d,
	0x29, 0x9c, 0xab, 0x3e, 0x89, 0xb0, 0x64, 0x3b, 0x69, 0xe7, 0x30, 0x5f, 0xf5, 0x7a, 0xfa, 0x63,
	0x4d, 0x33, 0x9c, 0xdc, 0xa8, 0x7e, 0xc2, 0x00, 0x6a, 0xaf, 0x60, 0xa4, 0x94, 0x30, 0x34, 0xe9,
	0x0d, 0x99, 0x1a, 0x78, 0x8a, 0xd1, 0x96, 0x94, 0x4c, 0x86, 0x66, 0x05, 0xa7, 0xcf, 0xea, 0xf8,
	0x3f, 0xea, 0xb5, 0x31, 0x86, 0x99, 0xda, 0xb5, 0x88, 0xb0, 0xa5, 0xdd, 0xe8, 0x6c, 0x9d, 0xbe,
	0xf5, 0xab, 0x0f, 0xb3, 0xb7, 0xdb, 0xd2, 0xd5, 0x1f, 0x9f, 0x8b, 0xbc, 0x60, 0xaf, 0xa3, 0xf3,
	0xd7, 0x4b, 0xf7, 0x4e, 0xba, 0x94, 0x50, 0xcc, 0x6d, 0x08, 0x73, 0x24, 0xcf, 0xda, 0x55, 0x35,
	0x49, 0x65, 0x97, 0xba, 0x34, 0x2f, 0x8f, 0xba, 0x35, 0x5c, 0xec, 0xa7, 0x2b, 0x5e, 0xfd, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x72, 0x10, 0x57, 0x48, 0x8b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepositoryServiceClient is the client API for RepositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryServiceClient interface {
	Add(ctx context.Context, in *domain.Repository, opts ...grpc.CallOption) (*AddRepositoryResponse, error)
}

type repositoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryServiceClient(cc *grpc.ClientConn) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) Add(ctx context.Context, in *domain.Repository, opts ...grpc.CallOption) (*AddRepositoryResponse, error) {
	out := new(AddRepositoryResponse)
	err := c.cc.Invoke(ctx, "/service.RepositoryService/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServiceServer is the server API for RepositoryService service.
type RepositoryServiceServer interface {
	Add(context.Context, *domain.Repository) (*AddRepositoryResponse, error)
}

// UnimplementedRepositoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepositoryServiceServer struct {
}

func (*UnimplementedRepositoryServiceServer) Add(ctx context.Context, req *domain.Repository) (*AddRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterRepositoryServiceServer(s *grpc.Server, srv RepositoryServiceServer) {
	s.RegisterService(&_RepositoryService_serviceDesc, srv)
}

func _RepositoryService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(domain.Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.RepositoryService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).Add(ctx, req.(*domain.Repository))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _RepositoryService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/repository-service.proto",
}
