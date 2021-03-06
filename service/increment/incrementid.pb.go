// Code generated by protoc-gen-go.
// source: incrementid.proto
// DO NOT EDIT!

/*
Package increment is a generated protocol buffer package.

It is generated from these files:
	incrementid.proto

It has these top-level messages:
	IncrIdNameRequest
	IncrIdReply
	IncrBoolReply
	IncrIdNameValueRequest
	NoneReply
*/
package increment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IncrIdNameRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *IncrIdNameRequest) Reset()                    { *m = IncrIdNameRequest{} }
func (m *IncrIdNameRequest) String() string            { return proto.CompactTextString(m) }
func (*IncrIdNameRequest) ProtoMessage()               {}
func (*IncrIdNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type IncrIdReply struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *IncrIdReply) Reset()                    { *m = IncrIdReply{} }
func (m *IncrIdReply) String() string            { return proto.CompactTextString(m) }
func (*IncrIdReply) ProtoMessage()               {}
func (*IncrIdReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type IncrBoolReply struct {
	Ret bool `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
}

func (m *IncrBoolReply) Reset()                    { *m = IncrBoolReply{} }
func (m *IncrBoolReply) String() string            { return proto.CompactTextString(m) }
func (*IncrBoolReply) ProtoMessage()               {}
func (*IncrBoolReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type IncrIdNameValueRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *IncrIdNameValueRequest) Reset()                    { *m = IncrIdNameValueRequest{} }
func (m *IncrIdNameValueRequest) String() string            { return proto.CompactTextString(m) }
func (*IncrIdNameValueRequest) ProtoMessage()               {}
func (*IncrIdNameValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type NoneReply struct {
}

func (m *NoneReply) Reset()                    { *m = NoneReply{} }
func (m *NoneReply) String() string            { return proto.CompactTextString(m) }
func (*NoneReply) ProtoMessage()               {}
func (*NoneReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*IncrIdNameRequest)(nil), "increment.IncrIdNameRequest")
	proto.RegisterType((*IncrIdReply)(nil), "increment.IncrIdReply")
	proto.RegisterType((*IncrBoolReply)(nil), "increment.IncrBoolReply")
	proto.RegisterType((*IncrIdNameValueRequest)(nil), "increment.IncrIdNameValueRequest")
	proto.RegisterType((*NoneReply)(nil), "increment.NoneReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for IncrementId service

type IncrementIdClient interface {
	GetIncrId(ctx context.Context, in *IncrIdNameRequest, opts ...grpc.CallOption) (*IncrIdReply, error)
	GetIncrIdByAmount(ctx context.Context, in *IncrIdNameValueRequest, opts ...grpc.CallOption) (*IncrIdReply, error)
	CheckIncrKeyExist(ctx context.Context, in *IncrIdNameRequest, opts ...grpc.CallOption) (*IncrBoolReply, error)
	CreateIncrKey(ctx context.Context, in *IncrIdNameValueRequest, opts ...grpc.CallOption) (*NoneReply, error)
}

type incrementIdClient struct {
	cc *grpc.ClientConn
}

func NewIncrementIdClient(cc *grpc.ClientConn) IncrementIdClient {
	return &incrementIdClient{cc}
}

func (c *incrementIdClient) GetIncrId(ctx context.Context, in *IncrIdNameRequest, opts ...grpc.CallOption) (*IncrIdReply, error) {
	out := new(IncrIdReply)
	err := grpc.Invoke(ctx, "/increment.IncrementId/GetIncrId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementIdClient) GetIncrIdByAmount(ctx context.Context, in *IncrIdNameValueRequest, opts ...grpc.CallOption) (*IncrIdReply, error) {
	out := new(IncrIdReply)
	err := grpc.Invoke(ctx, "/increment.IncrementId/GetIncrIdByAmount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementIdClient) CheckIncrKeyExist(ctx context.Context, in *IncrIdNameRequest, opts ...grpc.CallOption) (*IncrBoolReply, error) {
	out := new(IncrBoolReply)
	err := grpc.Invoke(ctx, "/increment.IncrementId/CheckIncrKeyExist", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementIdClient) CreateIncrKey(ctx context.Context, in *IncrIdNameValueRequest, opts ...grpc.CallOption) (*NoneReply, error) {
	out := new(NoneReply)
	err := grpc.Invoke(ctx, "/increment.IncrementId/CreateIncrKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IncrementId service

type IncrementIdServer interface {
	GetIncrId(context.Context, *IncrIdNameRequest) (*IncrIdReply, error)
	GetIncrIdByAmount(context.Context, *IncrIdNameValueRequest) (*IncrIdReply, error)
	CheckIncrKeyExist(context.Context, *IncrIdNameRequest) (*IncrBoolReply, error)
	CreateIncrKey(context.Context, *IncrIdNameValueRequest) (*NoneReply, error)
}

func RegisterIncrementIdServer(s *grpc.Server, srv IncrementIdServer) {
	s.RegisterService(&_IncrementId_serviceDesc, srv)
}

func _IncrementId_GetIncrId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrIdNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementIdServer).GetIncrId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementId/GetIncrId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementIdServer).GetIncrId(ctx, req.(*IncrIdNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementId_GetIncrIdByAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrIdNameValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementIdServer).GetIncrIdByAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementId/GetIncrIdByAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementIdServer).GetIncrIdByAmount(ctx, req.(*IncrIdNameValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementId_CheckIncrKeyExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrIdNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementIdServer).CheckIncrKeyExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementId/CheckIncrKeyExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementIdServer).CheckIncrKeyExist(ctx, req.(*IncrIdNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncrementId_CreateIncrKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrIdNameValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementIdServer).CreateIncrKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/increment.IncrementId/CreateIncrKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementIdServer).CreateIncrKey(ctx, req.(*IncrIdNameValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IncrementId_serviceDesc = grpc.ServiceDesc{
	ServiceName: "increment.IncrementId",
	HandlerType: (*IncrementIdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIncrId",
			Handler:    _IncrementId_GetIncrId_Handler,
		},
		{
			MethodName: "GetIncrIdByAmount",
			Handler:    _IncrementId_GetIncrIdByAmount_Handler,
		},
		{
			MethodName: "CheckIncrKeyExist",
			Handler:    _IncrementId_CheckIncrKeyExist_Handler,
		},
		{
			MethodName: "CreateIncrKey",
			Handler:    _IncrementId_CreateIncrKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("incrementid.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xdb, 0x58, 0xc5, 0x4c, 0xa9, 0x98, 0xa1, 0x94, 0x20, 0x0a, 0x76, 0x2f, 0x7a, 0xca,
	0x41, 0x9f, 0xc0, 0x04, 0x91, 0x28, 0x16, 0xc9, 0xc1, 0x7b, 0xcc, 0x0e, 0xb8, 0x98, 0xec, 0xd6,
	0xed, 0x46, 0xcc, 0xd3, 0xf9, 0x6a, 0x92, 0x5d, 0x59, 0x8b, 0x95, 0xea, 0x6d, 0x76, 0xe6, 0xe3,
	0xff, 0xe7, 0x1f, 0x16, 0x22, 0x21, 0x2b, 0x4d, 0x0d, 0x49, 0x23, 0x78, 0xb2, 0xd4, 0xca, 0x28,
	0x0c, 0x7d, 0x8b, 0x9d, 0x41, 0x94, 0xcb, 0x4a, 0xe7, 0x7c, 0x51, 0x36, 0x54, 0xd0, 0x6b, 0x4b,
	0x2b, 0x83, 0x08, 0x23, 0x59, 0x36, 0x14, 0x0f, 0x4f, 0x87, 0xe7, 0x61, 0x61, 0x6b, 0x76, 0x02,
	0x63, 0x07, 0x16, 0xb4, 0xac, 0x3b, 0x3c, 0x80, 0x40, 0x70, 0x0b, 0x8c, 0x8a, 0x40, 0x70, 0x36,
	0x87, 0x49, 0x3f, 0x4e, 0x95, 0xaa, 0x1d, 0x70, 0x08, 0x3b, 0x9a, 0x8c, 0x25, 0xf6, 0x8b, 0xbe,
	0x64, 0x29, 0xcc, 0xbe, 0xad, 0x1e, 0xcb, 0xba, 0xdd, 0xe6, 0x87, 0x53, 0xd8, 0x7d, 0xeb, 0x99,
	0x38, 0xb0, 0x1e, 0xee, 0xc1, 0xc6, 0x10, 0x2e, 0x94, 0x24, 0x6b, 0x71, 0xf1, 0x11, 0xb8, 0x9d,
	0x6c, 0x92, 0x9c, 0x63, 0x06, 0xe1, 0x0d, 0x19, 0xe7, 0x81, 0xc7, 0x89, 0x0f, 0x99, 0x6c, 0x24,
	0x3c, 0x9a, 0x6d, 0x4c, 0xad, 0x24, 0x1b, 0xe0, 0x03, 0x44, 0x5e, 0x24, 0xed, 0xae, 0x1a, 0xd5,
	0x4a, 0x83, 0xf3, 0x5f, 0xc5, 0xd6, 0x33, 0x6c, 0x51, 0xbc, 0x87, 0x28, 0x7b, 0xa6, 0xea, 0xa5,
	0xef, 0xde, 0x51, 0x77, 0xfd, 0x2e, 0x56, 0xe6, 0x8f, 0xf5, 0xe2, 0x1f, 0x53, 0x7f, 0x56, 0x36,
	0xc0, 0x5b, 0x98, 0x64, 0x9a, 0x4a, 0x43, 0x5f, 0x7a, 0xff, 0x59, 0x6e, 0xba, 0x86, 0xf8, 0xfb,
	0xb1, 0xc1, 0xd3, 0x9e, 0xfd, 0x0f, 0x97, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x5a, 0x35,
	0xe9, 0x24, 0x02, 0x00, 0x00,
}
