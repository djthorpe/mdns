// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gaffer/kernel.proto

package gaffer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type KernelProcess_State int32

const (
	KernelProcess_NONE     KernelProcess_State = 0
	KernelProcess_NEW      KernelProcess_State = 1
	KernelProcess_RUNNING  KernelProcess_State = 2
	KernelProcess_STOPPING KernelProcess_State = 3
	KernelProcess_STOPPED  KernelProcess_State = 4
)

var KernelProcess_State_name = map[int32]string{
	0: "NONE",
	1: "NEW",
	2: "RUNNING",
	3: "STOPPING",
	4: "STOPPED",
}

var KernelProcess_State_value = map[string]int32{
	"NONE":     0,
	"NEW":      1,
	"RUNNING":  2,
	"STOPPING": 3,
	"STOPPED":  4,
}

func (x KernelProcess_State) String() string {
	return proto.EnumName(KernelProcess_State_name, int32(x))
}

func (KernelProcess_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b7683b100b2849c5, []int{4, 0}
}

type KernelProcessEvent_Type int32

const (
	KernelProcessEvent_NONE     KernelProcessEvent_Type = 0
	KernelProcessEvent_NEW      KernelProcessEvent_Type = 1
	KernelProcessEvent_RUNNING  KernelProcessEvent_Type = 2
	KernelProcessEvent_STOPPING KernelProcessEvent_Type = 3
	KernelProcessEvent_STOPPED  KernelProcessEvent_Type = 4
	KernelProcessEvent_STDOUT   KernelProcessEvent_Type = 5
	KernelProcessEvent_STDERR   KernelProcessEvent_Type = 6
)

var KernelProcessEvent_Type_name = map[int32]string{
	0: "NONE",
	1: "NEW",
	2: "RUNNING",
	3: "STOPPING",
	4: "STOPPED",
	5: "STDOUT",
	6: "STDERR",
}

var KernelProcessEvent_Type_value = map[string]int32{
	"NONE":     0,
	"NEW":      1,
	"RUNNING":  2,
	"STOPPING": 3,
	"STOPPED":  4,
	"STDOUT":   5,
	"STDERR":   6,
}

func (x KernelProcessEvent_Type) String() string {
	return proto.EnumName(KernelProcessEvent_Type_name, int32(x))
}

func (KernelProcessEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b7683b100b2849c5, []int{5, 0}
}

// A service defintion is used to start a process
type KernelService struct {
	Path                 string             `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Wd                   string             `protobuf:"bytes,2,opt,name=wd,proto3" json:"wd,omitempty"`
	User                 string             `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Group                string             `protobuf:"bytes,5,opt,name=group,proto3" json:"group,omitempty"`
	Args                 []string           `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
	Timeout              *duration.Duration `protobuf:"bytes,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Sid                  uint32             `protobuf:"varint,8,opt,name=sid,proto3" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *KernelService) Reset()         { *m = KernelService{} }
func (m *KernelService) String() string { return proto.CompactTextString(m) }
func (*KernelService) ProtoMessage()    {}
func (*KernelService) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7683b100b2849c5, []int{0}
}

func (m *KernelService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelService.Unmarshal(m, b)
}
func (m *KernelService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelService.Marshal(b, m, deterministic)
}
func (m *KernelService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelService.Merge(m, src)
}
func (m *KernelService) XXX_Size() int {
	return xxx_messageInfo_KernelService.Size(m)
}
func (m *KernelService) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelService.DiscardUnknown(m)
}

var xxx_messageInfo_KernelService proto.InternalMessageInfo

func (m *KernelService) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *KernelService) GetWd() string {
	if m != nil {
		return m.Wd
	}
	return ""
}

func (m *KernelService) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *KernelService) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *KernelService) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *KernelService) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *KernelService) GetSid() uint32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

// A returned process identifier
type KernelProcessId struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sid                  uint32   `protobuf:"varint,2,opt,name=sid,proto3" json:"sid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KernelProcessId) Reset()         { *m = KernelProcessId{} }
func (m *KernelProcessId) String() string { return proto.CompactTextString(m) }
func (*KernelProcessId) ProtoMessage()    {}
func (*KernelProcessId) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7683b100b2849c5, []int{1}
}

func (m *KernelProcessId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelProcessId.Unmarshal(m, b)
}
func (m *KernelProcessId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelProcessId.Marshal(b, m, deterministic)
}
func (m *KernelProcessId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelProcessId.Merge(m, src)
}
func (m *KernelProcessId) XXX_Size() int {
	return xxx_messageInfo_KernelProcessId.Size(m)
}
func (m *KernelProcessId) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelProcessId.DiscardUnknown(m)
}

var xxx_messageInfo_KernelProcessId proto.InternalMessageInfo

func (m *KernelProcessId) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KernelProcessId) GetSid() uint32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

// Return a list of processes
type KernelProcessList struct {
	Process              []*KernelProcess `protobuf:"bytes,1,rep,name=process,proto3" json:"process,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *KernelProcessList) Reset()         { *m = KernelProcessList{} }
func (m *KernelProcessList) String() string { return proto.CompactTextString(m) }
func (*KernelProcessList) ProtoMessage()    {}
func (*KernelProcessList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7683b100b2849c5, []int{2}
}

func (m *KernelProcessList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelProcessList.Unmarshal(m, b)
}
func (m *KernelProcessList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelProcessList.Marshal(b, m, deterministic)
}
func (m *KernelProcessList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelProcessList.Merge(m, src)
}
func (m *KernelProcessList) XXX_Size() int {
	return xxx_messageInfo_KernelProcessList.Size(m)
}
func (m *KernelProcessList) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelProcessList.DiscardUnknown(m)
}

var xxx_messageInfo_KernelProcessList proto.InternalMessageInfo

func (m *KernelProcessList) GetProcess() []*KernelProcess {
	if m != nil {
		return m.Process
	}
	return nil
}

// Return a list of executables
type KernelExecutableList struct {
	Executable           []string `protobuf:"bytes,1,rep,name=executable,proto3" json:"executable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KernelExecutableList) Reset()         { *m = KernelExecutableList{} }
func (m *KernelExecutableList) String() string { return proto.CompactTextString(m) }
func (*KernelExecutableList) ProtoMessage()    {}
func (*KernelExecutableList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7683b100b2849c5, []int{3}
}

func (m *KernelExecutableList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelExecutableList.Unmarshal(m, b)
}
func (m *KernelExecutableList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelExecutableList.Marshal(b, m, deterministic)
}
func (m *KernelExecutableList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelExecutableList.Merge(m, src)
}
func (m *KernelExecutableList) XXX_Size() int {
	return xxx_messageInfo_KernelExecutableList.Size(m)
}
func (m *KernelExecutableList) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelExecutableList.DiscardUnknown(m)
}

var xxx_messageInfo_KernelExecutableList proto.InternalMessageInfo

func (m *KernelExecutableList) GetExecutable() []string {
	if m != nil {
		return m.Executable
	}
	return nil
}

// Return information about a process
type KernelProcess struct {
	Id                   *KernelProcessId    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State                KernelProcess_State `protobuf:"varint,2,opt,name=state,proto3,enum=gaffer.KernelProcess_State" json:"state,omitempty"`
	Service              *KernelService      `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KernelProcess) Reset()         { *m = KernelProcess{} }
func (m *KernelProcess) String() string { return proto.CompactTextString(m) }
func (*KernelProcess) ProtoMessage()    {}
func (*KernelProcess) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7683b100b2849c5, []int{4}
}

func (m *KernelProcess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelProcess.Unmarshal(m, b)
}
func (m *KernelProcess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelProcess.Marshal(b, m, deterministic)
}
func (m *KernelProcess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelProcess.Merge(m, src)
}
func (m *KernelProcess) XXX_Size() int {
	return xxx_messageInfo_KernelProcess.Size(m)
}
func (m *KernelProcess) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelProcess.DiscardUnknown(m)
}

var xxx_messageInfo_KernelProcess proto.InternalMessageInfo

func (m *KernelProcess) GetId() *KernelProcessId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KernelProcess) GetState() KernelProcess_State {
	if m != nil {
		return m.State
	}
	return KernelProcess_NONE
}

func (m *KernelProcess) GetService() *KernelService {
	if m != nil {
		return m.Service
	}
	return nil
}

// Monitor discovery changes
type KernelProcessEvent struct {
	Type                 KernelProcessEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=gaffer.KernelProcessEvent_Type" json:"type,omitempty"`
	Id                   *KernelProcessId        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Buf                  []byte                  `protobuf:"bytes,3,opt,name=buf,proto3" json:"buf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KernelProcessEvent) Reset()         { *m = KernelProcessEvent{} }
func (m *KernelProcessEvent) String() string { return proto.CompactTextString(m) }
func (*KernelProcessEvent) ProtoMessage()    {}
func (*KernelProcessEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7683b100b2849c5, []int{5}
}

func (m *KernelProcessEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KernelProcessEvent.Unmarshal(m, b)
}
func (m *KernelProcessEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KernelProcessEvent.Marshal(b, m, deterministic)
}
func (m *KernelProcessEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KernelProcessEvent.Merge(m, src)
}
func (m *KernelProcessEvent) XXX_Size() int {
	return xxx_messageInfo_KernelProcessEvent.Size(m)
}
func (m *KernelProcessEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_KernelProcessEvent.DiscardUnknown(m)
}

var xxx_messageInfo_KernelProcessEvent proto.InternalMessageInfo

func (m *KernelProcessEvent) GetType() KernelProcessEvent_Type {
	if m != nil {
		return m.Type
	}
	return KernelProcessEvent_NONE
}

func (m *KernelProcessEvent) GetId() *KernelProcessId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KernelProcessEvent) GetBuf() []byte {
	if m != nil {
		return m.Buf
	}
	return nil
}

func init() {
	proto.RegisterEnum("gaffer.KernelProcess_State", KernelProcess_State_name, KernelProcess_State_value)
	proto.RegisterEnum("gaffer.KernelProcessEvent_Type", KernelProcessEvent_Type_name, KernelProcessEvent_Type_value)
	proto.RegisterType((*KernelService)(nil), "gaffer.KernelService")
	proto.RegisterType((*KernelProcessId)(nil), "gaffer.KernelProcessId")
	proto.RegisterType((*KernelProcessList)(nil), "gaffer.KernelProcessList")
	proto.RegisterType((*KernelExecutableList)(nil), "gaffer.KernelExecutableList")
	proto.RegisterType((*KernelProcess)(nil), "gaffer.KernelProcess")
	proto.RegisterType((*KernelProcessEvent)(nil), "gaffer.KernelProcessEvent")
}

func init() {
	proto.RegisterFile("gaffer/kernel.proto", fileDescriptor_b7683b100b2849c5)
}

var fileDescriptor_b7683b100b2849c5 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x7f, 0xfe, 0x13, 0xa7, 0x1d, 0x37, 0xfd, 0x99, 0xa5, 0x80, 0x9b, 0xa2, 0x12, 0xf9,
	0x42, 0x4e, 0x0e, 0x38, 0x52, 0x6f, 0x50, 0x89, 0xc6, 0x42, 0x11, 0xc8, 0x89, 0xd6, 0xa9, 0x10,
	0x47, 0x27, 0xde, 0x18, 0x8b, 0x24, 0xb6, 0xec, 0x75, 0x4b, 0xde, 0x8c, 0x97, 0xe0, 0x19, 0xe0,
	0x51, 0xd0, 0xee, 0xda, 0x29, 0x09, 0x71, 0x91, 0xe0, 0x36, 0x3b, 0xf3, 0x99, 0xd9, 0xef, 0xcc,
	0xee, 0xc0, 0xc3, 0x28, 0x98, 0xcf, 0x49, 0xd6, 0xfb, 0x4c, 0xb2, 0x15, 0x59, 0xd8, 0x69, 0x96,
	0xd0, 0x04, 0x69, 0xc2, 0xd9, 0x3e, 0x8b, 0x92, 0x24, 0x5a, 0x90, 0x1e, 0xf7, 0x4e, 0x8b, 0x79,
	0x8f, 0x2c, 0x53, 0xba, 0x16, 0x50, 0xfb, 0x7c, 0x37, 0x18, 0x16, 0x59, 0x40, 0xe3, 0x64, 0x25,
	0xe2, 0xd6, 0x57, 0x09, 0x5a, 0xef, 0x78, 0x55, 0x9f, 0x64, 0x37, 0xf1, 0x8c, 0x20, 0x04, 0x6a,
	0x1a, 0xd0, 0x4f, 0xa6, 0xd4, 0x91, 0xba, 0x87, 0x98, 0xdb, 0xe8, 0x18, 0xe4, 0xdb, 0xd0, 0x94,
	0xb9, 0x47, 0xbe, 0x0d, 0x19, 0x53, 0xe4, 0x24, 0x33, 0x55, 0xc1, 0x30, 0x1b, 0x9d, 0x40, 0x23,
	0xca, 0x92, 0x22, 0x35, 0x1b, 0xdc, 0x29, 0x0e, 0x8c, 0x0c, 0xb2, 0x28, 0x37, 0xb5, 0x8e, 0xc2,
	0x48, 0x66, 0xa3, 0x3e, 0x34, 0x69, 0xbc, 0x24, 0x49, 0x41, 0xcd, 0x66, 0x47, 0xea, 0xea, 0xce,
	0xa9, 0x2d, 0x54, 0xda, 0x95, 0x4a, 0x7b, 0x50, 0xaa, 0xc4, 0x15, 0x89, 0x0c, 0x50, 0xf2, 0x38,
	0x34, 0x0f, 0x3a, 0x52, 0xb7, 0x85, 0x99, 0x69, 0xf5, 0xe1, 0x7f, 0xa1, 0x7c, 0x9c, 0x25, 0x33,
	0x92, 0xe7, 0xc3, 0x90, 0xe9, 0x8c, 0x43, 0xae, 0xbc, 0x85, 0xe5, 0x38, 0xac, 0x92, 0xe4, 0xbb,
	0xa4, 0x01, 0x3c, 0xd8, 0x4a, 0x7a, 0x1f, 0xe7, 0x14, 0xf5, 0xa0, 0x99, 0x8a, 0xa3, 0x29, 0x75,
	0x94, 0xae, 0xee, 0x3c, 0xb2, 0xc5, 0x6c, 0xed, 0x2d, 0x16, 0x57, 0x94, 0x75, 0x01, 0x27, 0x22,
	0xe2, 0x7e, 0x21, 0xb3, 0x82, 0x06, 0xd3, 0x05, 0xe1, 0x85, 0xce, 0x01, 0xc8, 0xc6, 0xc3, 0x6b,
	0x1d, 0xe2, 0x5f, 0x3c, 0xd6, 0x8f, 0xcd, 0xb4, 0xcb, 0x92, 0xe8, 0xf9, 0x46, 0xb1, 0xee, 0x3c,
	0xd9, 0x7b, 0xeb, 0x30, 0xe4, 0xad, 0xbc, 0x84, 0x46, 0x4e, 0x03, 0x4a, 0x78, 0x33, 0xc7, 0xce,
	0xd9, 0x5e, 0xd6, 0xf6, 0x19, 0x82, 0x05, 0xc9, 0xda, 0xca, 0xc5, 0xa3, 0x9a, 0x0a, 0xbf, 0x60,
	0xa7, 0xad, 0xf2, 0xc5, 0x71, 0x45, 0x59, 0x6f, 0xa0, 0xc1, 0x0b, 0xa0, 0x03, 0x50, 0xbd, 0x91,
	0xe7, 0x1a, 0xff, 0xa1, 0x26, 0x28, 0x9e, 0xfb, 0xc1, 0x90, 0x90, 0x0e, 0x4d, 0x7c, 0xed, 0x79,
	0x43, 0xef, 0xad, 0x21, 0xa3, 0x23, 0x38, 0xf0, 0x27, 0xa3, 0xf1, 0x98, 0x9d, 0x14, 0x16, 0xe2,
	0x27, 0x77, 0x60, 0xa8, 0xd6, 0x77, 0x09, 0xd0, 0x96, 0x26, 0xf7, 0x86, 0xac, 0x28, 0xea, 0x83,
	0x4a, 0xd7, 0x29, 0xe1, 0x9d, 0x1e, 0x3b, 0xcf, 0xf6, 0xaa, 0xe7, 0xa4, 0x3d, 0x59, 0xa7, 0x04,
	0x73, 0xb8, 0x1c, 0x8e, 0xfc, 0xe7, 0xe1, 0x18, 0xa0, 0x4c, 0x8b, 0x39, 0xef, 0xf2, 0x08, 0x33,
	0xd3, 0xfa, 0x08, 0x2a, 0x2b, 0xf4, 0x0f, 0x9d, 0x20, 0x00, 0xcd, 0x9f, 0x0c, 0x46, 0xd7, 0x13,
	0xa3, 0x51, 0xda, 0x2e, 0xc6, 0x86, 0xe6, 0x7c, 0x53, 0x40, 0x13, 0x22, 0xd0, 0x05, 0xa8, 0xe3,
	0x78, 0x15, 0xa1, 0xc7, 0xbf, 0x7d, 0x60, 0x97, 0xed, 0x60, 0xbb, 0xc6, 0x8f, 0x2e, 0xa1, 0x75,
	0x95, 0x91, 0x80, 0x92, 0xea, 0x1b, 0xec, 0x7f, 0x99, 0x76, 0x5d, 0xd3, 0xe8, 0x15, 0x00, 0x2e,
	0x56, 0x55, 0x76, 0x1d, 0x56, 0x7b, 0xff, 0x6b, 0xd0, 0x7d, 0x9a, 0xa4, 0x7f, 0x9d, 0x7f, 0x09,
	0x87, 0x25, 0x44, 0xee, 0xc9, 0x3e, 0xdd, 0x1b, 0xe0, 0x8b, 0x72, 0x05, 0xfa, 0xdd, 0xea, 0xe4,
	0xb5, 0xf3, 0x7b, 0xba, 0x5d, 0x61, 0x67, 0xdb, 0x5c, 0x38, 0xf2, 0x69, 0x46, 0x82, 0x25, 0xff,
	0x38, 0xf7, 0x08, 0x69, 0xd7, 0x7f, 0xb7, 0x17, 0xd2, 0x54, 0xe3, 0x97, 0xf6, 0x7f, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x79, 0x7c, 0x46, 0x01, 0x65, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KernelClient is the client API for Kernel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KernelClient interface {
	// Simple ping method to show server is "up"
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// Create a new process
	CreateProcess(ctx context.Context, in *KernelService, opts ...grpc.CallOption) (*KernelProcessId, error)
	// Run a process
	RunProcess(ctx context.Context, in *KernelProcessId, opts ...grpc.CallOption) (*empty.Empty, error)
	// Stop a process
	StopProcess(ctx context.Context, in *KernelProcessId, opts ...grpc.CallOption) (*empty.Empty, error)
	// Return filtered list of processes
	Processes(ctx context.Context, in *KernelProcessId, opts ...grpc.CallOption) (*KernelProcessList, error)
	// Return list of service executables
	Executables(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KernelExecutableList, error)
	// Stream filtered list of events
	StreamEvents(ctx context.Context, in *KernelProcessId, opts ...grpc.CallOption) (Kernel_StreamEventsClient, error)
}

type kernelClient struct {
	cc grpc.ClientConnInterface
}

func NewKernelClient(cc grpc.ClientConnInterface) KernelClient {
	return &kernelClient{cc}
}

func (c *kernelClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gaffer.Kernel/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kernelClient) CreateProcess(ctx context.Context, in *KernelService, opts ...grpc.CallOption) (*KernelProcessId, error) {
	out := new(KernelProcessId)
	err := c.cc.Invoke(ctx, "/gaffer.Kernel/CreateProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kernelClient) RunProcess(ctx context.Context, in *KernelProcessId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gaffer.Kernel/RunProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kernelClient) StopProcess(ctx context.Context, in *KernelProcessId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/gaffer.Kernel/StopProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kernelClient) Processes(ctx context.Context, in *KernelProcessId, opts ...grpc.CallOption) (*KernelProcessList, error) {
	out := new(KernelProcessList)
	err := c.cc.Invoke(ctx, "/gaffer.Kernel/Processes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kernelClient) Executables(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KernelExecutableList, error) {
	out := new(KernelExecutableList)
	err := c.cc.Invoke(ctx, "/gaffer.Kernel/Executables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kernelClient) StreamEvents(ctx context.Context, in *KernelProcessId, opts ...grpc.CallOption) (Kernel_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Kernel_serviceDesc.Streams[0], "/gaffer.Kernel/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &kernelStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Kernel_StreamEventsClient interface {
	Recv() (*KernelProcessEvent, error)
	grpc.ClientStream
}

type kernelStreamEventsClient struct {
	grpc.ClientStream
}

func (x *kernelStreamEventsClient) Recv() (*KernelProcessEvent, error) {
	m := new(KernelProcessEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KernelServer is the server API for Kernel service.
type KernelServer interface {
	// Simple ping method to show server is "up"
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	// Create a new process
	CreateProcess(context.Context, *KernelService) (*KernelProcessId, error)
	// Run a process
	RunProcess(context.Context, *KernelProcessId) (*empty.Empty, error)
	// Stop a process
	StopProcess(context.Context, *KernelProcessId) (*empty.Empty, error)
	// Return filtered list of processes
	Processes(context.Context, *KernelProcessId) (*KernelProcessList, error)
	// Return list of service executables
	Executables(context.Context, *empty.Empty) (*KernelExecutableList, error)
	// Stream filtered list of events
	StreamEvents(*KernelProcessId, Kernel_StreamEventsServer) error
}

// UnimplementedKernelServer can be embedded to have forward compatible implementations.
type UnimplementedKernelServer struct {
}

func (*UnimplementedKernelServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedKernelServer) CreateProcess(ctx context.Context, req *KernelService) (*KernelProcessId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProcess not implemented")
}
func (*UnimplementedKernelServer) RunProcess(ctx context.Context, req *KernelProcessId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunProcess not implemented")
}
func (*UnimplementedKernelServer) StopProcess(ctx context.Context, req *KernelProcessId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopProcess not implemented")
}
func (*UnimplementedKernelServer) Processes(ctx context.Context, req *KernelProcessId) (*KernelProcessList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Processes not implemented")
}
func (*UnimplementedKernelServer) Executables(ctx context.Context, req *empty.Empty) (*KernelExecutableList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Executables not implemented")
}
func (*UnimplementedKernelServer) StreamEvents(req *KernelProcessId, srv Kernel_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}

func RegisterKernelServer(s *grpc.Server, srv KernelServer) {
	s.RegisterService(&_Kernel_serviceDesc, srv)
}

func _Kernel_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KernelServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Kernel/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KernelServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kernel_CreateProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KernelService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KernelServer).CreateProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Kernel/CreateProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KernelServer).CreateProcess(ctx, req.(*KernelService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kernel_RunProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KernelProcessId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KernelServer).RunProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Kernel/RunProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KernelServer).RunProcess(ctx, req.(*KernelProcessId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kernel_StopProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KernelProcessId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KernelServer).StopProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Kernel/StopProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KernelServer).StopProcess(ctx, req.(*KernelProcessId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kernel_Processes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KernelProcessId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KernelServer).Processes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Kernel/Processes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KernelServer).Processes(ctx, req.(*KernelProcessId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kernel_Executables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KernelServer).Executables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaffer.Kernel/Executables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KernelServer).Executables(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kernel_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(KernelProcessId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KernelServer).StreamEvents(m, &kernelStreamEventsServer{stream})
}

type Kernel_StreamEventsServer interface {
	Send(*KernelProcessEvent) error
	grpc.ServerStream
}

type kernelStreamEventsServer struct {
	grpc.ServerStream
}

func (x *kernelStreamEventsServer) Send(m *KernelProcessEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Kernel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gaffer.Kernel",
	HandlerType: (*KernelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Kernel_Ping_Handler,
		},
		{
			MethodName: "CreateProcess",
			Handler:    _Kernel_CreateProcess_Handler,
		},
		{
			MethodName: "RunProcess",
			Handler:    _Kernel_RunProcess_Handler,
		},
		{
			MethodName: "StopProcess",
			Handler:    _Kernel_StopProcess_Handler,
		},
		{
			MethodName: "Processes",
			Handler:    _Kernel_Processes_Handler,
		},
		{
			MethodName: "Executables",
			Handler:    _Kernel_Executables_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _Kernel_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gaffer/kernel.proto",
}
