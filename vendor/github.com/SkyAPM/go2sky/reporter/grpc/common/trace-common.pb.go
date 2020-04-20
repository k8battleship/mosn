// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/trace-common.proto

package common

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

type SpanType int32

const (
	SpanType_Entry SpanType = 0
	SpanType_Exit  SpanType = 1
	SpanType_Local SpanType = 2
)

var SpanType_name = map[int32]string{
	0: "Entry",
	1: "Exit",
	2: "Local",
}

var SpanType_value = map[string]int32{
	"Entry": 0,
	"Exit":  1,
	"Local": 2,
}

func (x SpanType) String() string {
	return proto.EnumName(SpanType_name, int32(x))
}

func (SpanType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{0}
}

type RefType int32

const (
	RefType_CrossProcess RefType = 0
	RefType_CrossThread  RefType = 1
)

var RefType_name = map[int32]string{
	0: "CrossProcess",
	1: "CrossThread",
}

var RefType_value = map[string]int32{
	"CrossProcess": 0,
	"CrossThread":  1,
}

func (x RefType) String() string {
	return proto.EnumName(RefType_name, int32(x))
}

func (RefType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{1}
}

type SpanLayer int32

const (
	SpanLayer_Unknown      SpanLayer = 0
	SpanLayer_Database     SpanLayer = 1
	SpanLayer_RPCFramework SpanLayer = 2
	SpanLayer_Http         SpanLayer = 3
	SpanLayer_MQ           SpanLayer = 4
	SpanLayer_Cache        SpanLayer = 5
)

var SpanLayer_name = map[int32]string{
	0: "Unknown",
	1: "Database",
	2: "RPCFramework",
	3: "Http",
	4: "MQ",
	5: "Cache",
}

var SpanLayer_value = map[string]int32{
	"Unknown":      0,
	"Database":     1,
	"RPCFramework": 2,
	"Http":         3,
	"MQ":           4,
	"Cache":        5,
}

func (x SpanLayer) String() string {
	return proto.EnumName(SpanLayer_name, int32(x))
}

func (SpanLayer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{2}
}

type UpstreamSegment struct {
	GlobalTraceIds       []*UniqueId `protobuf:"bytes,1,rep,name=globalTraceIds,proto3" json:"globalTraceIds,omitempty"`
	Segment              []byte      `protobuf:"bytes,2,opt,name=segment,proto3" json:"segment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpstreamSegment) Reset()         { *m = UpstreamSegment{} }
func (m *UpstreamSegment) String() string { return proto.CompactTextString(m) }
func (*UpstreamSegment) ProtoMessage()    {}
func (*UpstreamSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{0}
}

func (m *UpstreamSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSegment.Unmarshal(m, b)
}
func (m *UpstreamSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSegment.Marshal(b, m, deterministic)
}
func (m *UpstreamSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSegment.Merge(m, src)
}
func (m *UpstreamSegment) XXX_Size() int {
	return xxx_messageInfo_UpstreamSegment.Size(m)
}
func (m *UpstreamSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSegment.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSegment proto.InternalMessageInfo

func (m *UpstreamSegment) GetGlobalTraceIds() []*UniqueId {
	if m != nil {
		return m.GlobalTraceIds
	}
	return nil
}

func (m *UpstreamSegment) GetSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

type UniqueId struct {
	IdParts              []int64  `protobuf:"varint,1,rep,packed,name=idParts,proto3" json:"idParts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UniqueId) Reset()         { *m = UniqueId{} }
func (m *UniqueId) String() string { return proto.CompactTextString(m) }
func (*UniqueId) ProtoMessage()    {}
func (*UniqueId) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5d7d2e503402948, []int{1}
}

func (m *UniqueId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UniqueId.Unmarshal(m, b)
}
func (m *UniqueId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UniqueId.Marshal(b, m, deterministic)
}
func (m *UniqueId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniqueId.Merge(m, src)
}
func (m *UniqueId) XXX_Size() int {
	return xxx_messageInfo_UniqueId.Size(m)
}
func (m *UniqueId) XXX_DiscardUnknown() {
	xxx_messageInfo_UniqueId.DiscardUnknown(m)
}

var xxx_messageInfo_UniqueId proto.InternalMessageInfo

func (m *UniqueId) GetIdParts() []int64 {
	if m != nil {
		return m.IdParts
	}
	return nil
}

func init() {
	proto.RegisterEnum("SpanType", SpanType_name, SpanType_value)
	proto.RegisterEnum("RefType", RefType_name, RefType_value)
	proto.RegisterEnum("SpanLayer", SpanLayer_name, SpanLayer_value)
	proto.RegisterType((*UpstreamSegment)(nil), "UpstreamSegment")
	proto.RegisterType((*UniqueId)(nil), "UniqueId")
}

func init() { proto.RegisterFile("common/trace-common.proto", fileDescriptor_f5d7d2e503402948) }

var fileDescriptor_f5d7d2e503402948 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0xa7, 0x6d, 0x92, 0x49, 0x44, 0x57, 0x3e, 0x05, 0x4e, 0x55, 0xc5, 0xa1, 0x8a,
	0xa8, 0x0d, 0xe5, 0x09, 0x20, 0x14, 0x51, 0xa9, 0x45, 0xc6, 0x49, 0x84, 0xc4, 0x01, 0x69, 0xb2,
	0x1e, 0x36, 0x96, 0xed, 0xdd, 0x65, 0x76, 0xa3, 0xe0, 0x1b, 0xcf, 0xc3, 0x53, 0xa2, 0x8d, 0xc9,
	0x85, 0xe3, 0xff, 0xeb, 0x9b, 0x6f, 0x57, 0xfa, 0xe1, 0xb9, 0x34, 0x6d, 0x6b, 0x74, 0xe6, 0x19,
	0x25, 0xdd, 0xf6, 0x21, 0xb5, 0x6c, 0xbc, 0xb9, 0xfe, 0x0e, 0x97, 0x1b, 0xeb, 0x3c, 0x13, 0xb6,
	0x2b, 0x52, 0x2d, 0x69, 0x9f, 0xbc, 0x81, 0x67, 0xaa, 0x31, 0x5b, 0x6c, 0xd6, 0x01, 0x7f, 0x28,
	0xdd, 0x3c, 0xba, 0x1a, 0xde, 0x4c, 0xef, 0x26, 0xe9, 0x46, 0x57, 0x3f, 0xf7, 0xf4, 0x50, 0x16,
	0xff, 0x01, 0xc9, 0x1c, 0x46, 0xae, 0xbf, 0x9e, 0xc7, 0x57, 0xd1, 0xcd, 0xac, 0x38, 0xc5, 0xeb,
	0x97, 0x30, 0x3e, 0x5d, 0x05, 0xaa, 0x2a, 0x73, 0x64, 0xdf, 0x1b, 0x87, 0xc5, 0x29, 0x2e, 0x16,
	0x30, 0x5e, 0x59, 0xd4, 0xeb, 0xce, 0x52, 0x32, 0x81, 0xf3, 0x7b, 0xed, 0xb9, 0x13, 0x83, 0x64,
	0x0c, 0x67, 0xf7, 0xbf, 0x2a, 0x2f, 0xa2, 0x50, 0x3e, 0x1a, 0x89, 0x8d, 0x88, 0x17, 0xaf, 0x60,
	0x54, 0xd0, 0x8f, 0x23, 0x2a, 0x60, 0xb6, 0x64, 0xe3, 0x5c, 0xce, 0x46, 0x92, 0x73, 0x62, 0x90,
	0x5c, 0xc2, 0xf4, 0xd8, 0xac, 0x77, 0x4c, 0x58, 0x8a, 0x68, 0xb1, 0x81, 0x49, 0x30, 0x3f, 0x62,
	0x47, 0x9c, 0x4c, 0x61, 0xb4, 0xd1, 0xb5, 0x36, 0x07, 0x2d, 0x06, 0xc9, 0x0c, 0xc6, 0x1f, 0xd0,
	0xe3, 0x16, 0x1d, 0x89, 0x28, 0xa8, 0x8a, 0x7c, 0xf9, 0x91, 0xb1, 0xa5, 0x83, 0xe1, 0x5a, 0xc4,
	0xe1, 0xf1, 0x4f, 0xde, 0x5b, 0x31, 0x4c, 0x2e, 0x20, 0x7e, 0xfa, 0x22, 0xce, 0xc2, 0x27, 0x96,
	0x28, 0x77, 0x24, 0xce, 0xdf, 0xff, 0x8e, 0xe0, 0xb5, 0x61, 0x95, 0xa2, 0x0d, 0x45, 0xea, 0xea,
	0xee, 0x80, 0x4d, 0x5d, 0xe9, 0xd0, 0xb4, 0xa9, 0x26, 0x1f, 0x24, 0x69, 0x83, 0x5a, 0xed, 0x51,
	0x51, 0x8a, 0x8a, 0xb4, 0xcf, 0xa3, 0x6f, 0xb7, 0xaa, 0xf2, 0xbb, 0xfd, 0x36, 0x95, 0xa6, 0xcd,
	0x56, 0x75, 0xf7, 0x2e, 0x7f, 0xca, 0x94, 0xb9, 0x73, 0x75, 0x97, 0x31, 0x59, 0xc3, 0x9e, 0x38,
	0x53, 0x6c, 0x65, 0xd6, 0x4f, 0xf4, 0x27, 0x7e, 0xb1, 0xaa, 0xbb, 0xaf, 0xff, 0xc4, 0x9f, 0x7b,
	0x69, 0x1e, 0x66, 0x93, 0xa6, 0xd9, 0x5e, 0x1c, 0x07, 0x7c, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x23, 0xcc, 0xa6, 0xdd, 0x01, 0x00, 0x00,
}