// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gaffer/gaffer.proto

package gaffer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Update a service
type ServiceUpdateRequest struct {
	Service              *Service              `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Fields               *field_mask.FieldMask `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ServiceUpdateRequest) Reset()         { *m = ServiceUpdateRequest{} }
func (m *ServiceUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceUpdateRequest) ProtoMessage()    {}
func (*ServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5f7d067d470eb10, []int{0}
}

func (m *ServiceUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceUpdateRequest.Unmarshal(m, b)
}
func (m *ServiceUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceUpdateRequest.Marshal(b, m, deterministic)
}
func (m *ServiceUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceUpdateRequest.Merge(m, src)
}
func (m *ServiceUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceUpdateRequest.Size(m)
}
func (m *ServiceUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceUpdateRequest proto.InternalMessageInfo

func (m *ServiceUpdateRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ServiceUpdateRequest) GetFields() *field_mask.FieldMask {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceUpdateRequest)(nil), "gaffer.ServiceUpdateRequest")
}

func init() {
	proto.RegisterFile("gaffer/gaffer.proto", fileDescriptor_a5f7d067d470eb10)
}

var fileDescriptor_a5f7d067d470eb10 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4f, 0x4c, 0x4b,
	0x4b, 0x2d, 0xd2, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x94,
	0x74, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x34, 0xa9, 0x34, 0x4d, 0x3f, 0x35, 0xb7,
	0xa0, 0xa4, 0x12, 0xa2, 0x48, 0x4a, 0x01, 0x5d, 0x32, 0x2d, 0x33, 0x35, 0x27, 0x25, 0x3e, 0x37,
	0xb1, 0x38, 0x1b, 0xaa, 0x42, 0x04, 0x6a, 0x76, 0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0x2a, 0x44,
	0x54, 0xa9, 0x94, 0x4b, 0x24, 0x18, 0x22, 0x10, 0x5a, 0x90, 0x92, 0x58, 0x92, 0x1a, 0x94, 0x5a,
	0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xc9, 0xc5, 0x0e, 0x55, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x6d, 0xc4, 0xaf, 0x07, 0x75, 0x14, 0x54, 0x79, 0x10, 0x4c, 0x5e, 0xc8, 0x88, 0x8b, 0x0d, 0x6c,
	0x59, 0xb1, 0x04, 0x13, 0x58, 0xa5, 0x94, 0x1e, 0xc4, 0x2d, 0x7a, 0x30, 0xb7, 0xe8, 0xb9, 0x81,
	0xa4, 0x7d, 0x13, 0x8b, 0xb3, 0x83, 0xa0, 0x2a, 0x8d, 0x9e, 0x33, 0x72, 0xb1, 0xb9, 0x83, 0xcd,
	0x13, 0x32, 0xe3, 0x62, 0x09, 0xc8, 0xcc, 0x4b, 0x17, 0x12, 0xc3, 0xd0, 0xe6, 0x0a, 0xf2, 0x9f,
	0x14, 0x0e, 0x71, 0x21, 0x73, 0x2e, 0x0e, 0xa8, 0x53, 0x8a, 0x71, 0xea, 0x15, 0x46, 0x73, 0xb4,
	0x4f, 0x66, 0x71, 0x89, 0x90, 0x35, 0x17, 0x1b, 0xc4, 0xaf, 0x42, 0x32, 0x68, 0xd2, 0x28, 0x41,
	0x80, 0x5d, 0xb3, 0x3e, 0x17, 0x6b, 0x70, 0x49, 0x62, 0x51, 0x89, 0x90, 0x20, 0x9a, 0xac, 0x67,
	0x0a, 0x56, 0x0d, 0x49, 0x6c, 0x60, 0x27, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xfc,
	0x9a, 0x0d, 0xdb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GafferClient is the client API for Gaffer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GafferClient interface {
	// Simple ping method to show server is "up"
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// Return services
	Services(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServiceList, error)
	// Update a service
	Update(ctx context.Context, in *ServiceUpdateRequest, opts ...grpc.CallOption) (*ServiceList, error)
	// Start a service
	Start(ctx context.Context, in *ServiceId, opts ...grpc.CallOption) (*ServiceList, error)
}

type gafferClient struct {
	cc grpc.ClientConnInterface
}

func NewGafferClient(cc grpc.ClientConnInterface) GafferClient {
	return &gafferClient{cc}
}

func (c *gafferClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gaffer.Gaffer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gafferClient) Services(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServiceList, error) {
	out := new(ServiceList)
	err := c.cc.Invoke(ctx, "/gaffer.Gaffer/Services", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gafferClient) Update(ctx context.Context, in *ServiceUpdateRequest, opts ...grpc.CallOption) (*ServiceList, error) {
	out := new(ServiceList)
	err := c.cc.Invoke(ctx, "/gaffer.Gaffer/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gafferClient) Start(ctx context.Context, in *ServiceId, opts ...grpc.CallOption) (*ServiceList, error) {
	out := new(ServiceList)
	err := c.cc.Invoke(ctx, "/gaffer.Gaffer/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GafferServer is the server API for Gaffer service.
type GafferServer interface {
	// Simple ping method to show server is "up"
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	// Return services
	Services(context.Context, *empty.Empty) (*ServiceList, error)
	// Update a service
	Update(context.Context, *ServiceUpdateRequest) (*ServiceList, error)
	// Start a service
	Start(context.Context, *ServiceId) (*ServiceList, error)
}

// UnimplementedGafferServer can be embedded to have forward compatible implementations.
type UnimplementedGafferServer struct {
}

func (*UnimplementedGafferServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedGafferServer) Services(ctx context.Context, req *empty.Empty) (*ServiceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Services not implemented")
}
func (*UnimplementedGafferServer) Update(ctx context.Context, req *ServiceUpdateRequest) (*ServiceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedGafferServer) Start(ctx context.Context, req *ServiceId) (*ServiceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}

func RegisterGafferServer(s *grpc.Server, srv GafferServer) {
	s.RegisterService(&_Gaffer_serviceDesc, srv)
}

func _Gaffer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GafferServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Gaffer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GafferServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gaffer_Services_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GafferServer).Services(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Gaffer/Services",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GafferServer).Services(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gaffer_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GafferServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Gaffer/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GafferServer).Update(ctx, req.(*ServiceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gaffer_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GafferServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Gaffer/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GafferServer).Start(ctx, req.(*ServiceId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gaffer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gaffer.Gaffer",
	HandlerType: (*GafferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Gaffer_Ping_Handler,
		},
		{
			MethodName: "Services",
			Handler:    _Gaffer_Services_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Gaffer_Update_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Gaffer_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gaffer/gaffer.proto",
}
