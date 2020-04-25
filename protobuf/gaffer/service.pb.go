// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gaffer/service.proto

package gaffer

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

// A service definition is used to start one or more processes
type Service struct {
	Sid                  uint32   `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Cwd                  string   `protobuf:"bytes,4,opt,name=cwd,proto3" json:"cwd,omitempty"`
	Args                 []string `protobuf:"bytes,5,rep,name=args,proto3" json:"args,omitempty"`
	User                 string   `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Group                string   `protobuf:"bytes,7,opt,name=group,proto3" json:"group,omitempty"`
	Enabled              bool     `protobuf:"varint,8,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_d89f813df1b8899b, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetSid() uint32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Service) GetCwd() string {
	if m != nil {
		return m.Cwd
	}
	return ""
}

func (m *Service) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Service) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Service) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Service) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// Return a list of services
type ServiceList struct {
	Service              []*Service `protobuf:"bytes,1,rep,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ServiceList) Reset()         { *m = ServiceList{} }
func (m *ServiceList) String() string { return proto.CompactTextString(m) }
func (*ServiceList) ProtoMessage()    {}
func (*ServiceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d89f813df1b8899b, []int{1}
}

func (m *ServiceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceList.Unmarshal(m, b)
}
func (m *ServiceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceList.Marshal(b, m, deterministic)
}
func (m *ServiceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceList.Merge(m, src)
}
func (m *ServiceList) XXX_Size() int {
	return xxx_messageInfo_ServiceList.Size(m)
}
func (m *ServiceList) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceList.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceList proto.InternalMessageInfo

func (m *ServiceList) GetService() []*Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "gaffer.Service")
	proto.RegisterType((*ServiceList)(nil), "gaffer.ServiceList")
}

func init() {
	proto.RegisterFile("gaffer/service.proto", fileDescriptor_d89f813df1b8899b)
}

var fileDescriptor_d89f813df1b8899b = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x41, 0x4e, 0x86, 0x30,
	0x10, 0x85, 0x53, 0x0b, 0x14, 0x87, 0x18, 0x4d, 0xc3, 0x62, 0x96, 0x0d, 0xab, 0xba, 0xc1, 0x44,
	0x37, 0x1e, 0xc2, 0x55, 0x3d, 0x41, 0x81, 0x82, 0x24, 0x0a, 0xa4, 0x05, 0x3d, 0x95, 0x77, 0xfc,
	0x33, 0x2d, 0xec, 0xbe, 0xf9, 0xe6, 0x4d, 0x32, 0x0f, 0xea, 0xc9, 0x8e, 0xa3, 0xf3, 0x2f, 0xc1,
	0xf9, 0xdf, 0xb9, 0x77, 0xed, 0xe6, 0xd7, 0x7d, 0x95, 0x45, 0xb2, 0xcd, 0x3f, 0x03, 0xf1, 0x99,
	0x36, 0xf2, 0x09, 0x78, 0x98, 0x07, 0x64, 0x8a, 0xe9, 0x07, 0x43, 0x28, 0x25, 0x64, 0x8b, 0xfd,
	0x71, 0x78, 0xa7, 0x98, 0xbe, 0x37, 0x91, 0xc9, 0x6d, 0x76, 0xff, 0x42, 0x9e, 0x1c, 0x31, 0x5d,
	0xf6, 0x7f, 0x03, 0x66, 0x51, 0x11, 0x52, 0xca, 0xfa, 0x29, 0x60, 0xae, 0x38, 0xa5, 0x88, 0xc9,
	0x1d, 0xc1, 0x79, 0x2c, 0xd2, 0x25, 0xb1, 0xac, 0x21, 0x9f, 0xfc, 0x7a, 0x6c, 0x28, 0xa2, 0x4c,
	0x83, 0x44, 0x10, 0x6e, 0xb1, 0xdd, 0xb7, 0x1b, 0xb0, 0x54, 0x4c, 0x97, 0xe6, 0x1a, 0x9b, 0x77,
	0xa8, 0xce, 0x77, 0x3f, 0xe6, 0xb0, 0xcb, 0x67, 0x10, 0x67, 0x2f, 0x64, 0x8a, 0xeb, 0xea, 0xf5,
	0xb1, 0x4d, 0xc5, 0xda, 0x33, 0x65, 0xae, 0x7d, 0x57, 0xc4, 0xe2, 0x6f, 0xb7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x19, 0x53, 0xb4, 0x52, 0x10, 0x01, 0x00, 0x00,
}
