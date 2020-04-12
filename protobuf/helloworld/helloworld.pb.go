// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld/helloworld.proto

package gopi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_73149fedf49f4319, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "gopi.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "gopi.HelloReply")
}

func init() {
	proto.RegisterFile("helloworld/helloworld.proto", fileDescriptor_73149fedf49f4319)
}

var fileDescriptor_73149fedf49f4319 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x58, 0xd2, 0xf3, 0x0b, 0x32, 0xa5, 0xa4, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x62,
	0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10, 0x25, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e,
	0x20, 0x6d, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x1a, 0x17, 0x17, 0x54, 0x4d,
	0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x3a, 0x4c, 0x11, 0x8c,
	0x6b, 0x54, 0x06, 0x55, 0x07, 0x76, 0x82, 0x90, 0x19, 0x17, 0x4b, 0x40, 0x66, 0x5e, 0xba, 0x90,
	0x98, 0x1e, 0xc4, 0x7e, 0x3d, 0x98, 0xfd, 0x7a, 0xae, 0x20, 0xfb, 0xa5, 0x70, 0x88, 0x0b, 0x19,
	0x70, 0x71, 0x04, 0x27, 0x56, 0x82, 0x0d, 0x12, 0x12, 0xd2, 0x03, 0xf9, 0x40, 0x0f, 0xd9, 0x85,
	0x52, 0x02, 0x28, 0x62, 0x05, 0x39, 0x95, 0x49, 0x6c, 0x60, 0x13, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x72, 0xd5, 0x0d, 0x9d, 0x0c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloworldClient is the client API for Helloworld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloworldClient interface {
	// Simple ping method to show server is "up"
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type helloworldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloworldClient(cc grpc.ClientConnInterface) HelloworldClient {
	return &helloworldClient{cc}
}

func (c *helloworldClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gopi.Helloworld/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/gopi.Helloworld/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloworldServer is the server API for Helloworld service.
type HelloworldServer interface {
	// Simple ping method to show server is "up"
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedHelloworldServer can be embedded to have forward compatible implementations.
type UnimplementedHelloworldServer struct {
}

func (*UnimplementedHelloworldServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedHelloworldServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterHelloworldServer(s *grpc.Server, srv HelloworldServer) {
	s.RegisterService(&_Helloworld_serviceDesc, srv)
}

func _Helloworld_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopi.Helloworld/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopi.Helloworld/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Helloworld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gopi.Helloworld",
	HandlerType: (*HelloworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Helloworld_Ping_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _Helloworld_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}
