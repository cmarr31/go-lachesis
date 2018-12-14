// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roundInfo.proto

package poset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Trilean int32

const (
	Trilean_UNDEFINED Trilean = 0
	Trilean_TRUE      Trilean = 1
	Trilean_FALSE     Trilean = 2
)

var Trilean_name = map[int32]string{
	0: "UNDEFINED",
	1: "TRUE",
	2: "FALSE",
}
var Trilean_value = map[string]int32{
	"UNDEFINED": 0,
	"TRUE":      1,
	"FALSE":     2,
}

func (x Trilean) String() string {
	return proto.EnumName(Trilean_name, int32(x))
}
func (Trilean) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_roundInfo_77608f4a2972004f, []int{0}
}

type RoundEvent struct {
	Consensus            bool     `protobuf:"varint,1,opt,name=Consensus,proto3" json:"Consensus,omitempty"`
	Witness              bool     `protobuf:"varint,2,opt,name=Witness,proto3" json:"Witness,omitempty"`
	Famous               Trilean  `protobuf:"varint,3,opt,name=Famous,proto3,enum=poset.Trilean" json:"Famous,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoundEvent) Reset()         { *m = RoundEvent{} }
func (m *RoundEvent) String() string { return proto.CompactTextString(m) }
func (*RoundEvent) ProtoMessage()    {}
func (*RoundEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_roundInfo_77608f4a2972004f, []int{0}
}
func (m *RoundEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundEvent.Unmarshal(m, b)
}
func (m *RoundEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundEvent.Marshal(b, m, deterministic)
}
func (dst *RoundEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundEvent.Merge(dst, src)
}
func (m *RoundEvent) XXX_Size() int {
	return xxx_messageInfo_RoundEvent.Size(m)
}
func (m *RoundEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RoundEvent proto.InternalMessageInfo

func (m *RoundEvent) GetConsensus() bool {
	if m != nil {
		return m.Consensus
	}
	return false
}

func (m *RoundEvent) GetWitness() bool {
	if m != nil {
		return m.Witness
	}
	return false
}

func (m *RoundEvent) GetFamous() Trilean {
	if m != nil {
		return m.Famous
	}
	return Trilean_UNDEFINED
}

type RoundInfoMessage struct {
	Events               map[string]*RoundEvent `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Queued               bool                   `protobuf:"varint,2,opt,name=queued,proto3" json:"queued,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RoundInfoMessage) Reset()         { *m = RoundInfoMessage{} }
func (m *RoundInfoMessage) String() string { return proto.CompactTextString(m) }
func (*RoundInfoMessage) ProtoMessage()    {}
func (*RoundInfoMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_roundInfo_77608f4a2972004f, []int{1}
}
func (m *RoundInfoMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundInfoMessage.Unmarshal(m, b)
}
func (m *RoundInfoMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundInfoMessage.Marshal(b, m, deterministic)
}
func (dst *RoundInfoMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundInfoMessage.Merge(dst, src)
}
func (m *RoundInfoMessage) XXX_Size() int {
	return xxx_messageInfo_RoundInfoMessage.Size(m)
}
func (m *RoundInfoMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundInfoMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RoundInfoMessage proto.InternalMessageInfo

func (m *RoundInfoMessage) GetEventBlocks() map[string]*RoundEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *RoundInfoMessage) GetQueued() bool {
	if m != nil {
		return m.Queued
	}
	return false
}

func init() {
	proto.RegisterType((*RoundEvent)(nil), "poset.RoundEvent")
	proto.RegisterType((*RoundInfoMessage)(nil), "poset.RoundInfoMessage")
	proto.RegisterMapType((map[string]*RoundEvent)(nil), "poset.RoundInfoMessage.EventsEntry")
	proto.RegisterEnum("poset.Trilean", Trilean_name, Trilean_value)
}

func init() { proto.RegisterFile("roundInfo.proto", fileDescriptor_roundInfo_77608f4a2972004f) }

var fileDescriptor_roundInfo_77608f4a2972004f = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xcd, 0x6a, 0xbb, 0xf5, 0x15, 0x67, 0x7d, 0x07, 0x29, 0xe2, 0x61, 0x4c, 0xd0, 0x22,
	0xd8, 0x43, 0xbd, 0x88, 0x9e, 0xc4, 0xa5, 0x30, 0x98, 0x3b, 0xc4, 0x0d, 0xcf, 0x95, 0x3d, 0x65,
	0x58, 0x93, 0xd9, 0x34, 0x83, 0xfd, 0x61, 0xfe, 0x7f, 0xd2, 0x2c, 0xc3, 0xe1, 0x2d, 0x2f, 0xdf,
	0xf7, 0xcb, 0xf7, 0xbe, 0xc0, 0x71, 0xad, 0x8c, 0x5c, 0x8c, 0xe5, 0xbb, 0xca, 0x56, 0xb5, 0x6a,
	0x14, 0xfa, 0x2b, 0xa5, 0xa9, 0x19, 0x56, 0x00, 0xa2, 0x55, 0xf8, 0x9a, 0x64, 0x83, 0xe7, 0x10,
	0x3e, 0x29, 0xa9, 0x49, 0x6a, 0xa3, 0x13, 0x36, 0x60, 0x69, 0x4f, 0xfc, 0x5d, 0x60, 0x02, 0xdd,
	0xd7, 0x65, 0x23, 0x49, 0xeb, 0xa4, 0x63, 0xb5, 0xdd, 0x88, 0x97, 0x10, 0x14, 0xe5, 0x97, 0x32,
	0x3a, 0xf1, 0x06, 0x2c, 0xed, 0xe7, 0xfd, 0xcc, 0xbe, 0x9e, 0xcd, 0xea, 0x65, 0x45, 0xa5, 0x14,
	0x4e, 0x1d, 0xfe, 0x30, 0x88, 0xc5, 0x6e, 0x91, 0x67, 0xd2, 0xba, 0xfc, 0x20, 0x7c, 0x80, 0xc0,
	0xa6, 0xb7, 0x89, 0x5e, 0x1a, 0xe5, 0x17, 0x0e, 0xfe, 0x6f, 0xcc, 0xb6, 0x2e, 0x2e, 0x9b, 0x7a,
	0x23, 0x1c, 0x82, 0xa7, 0x10, 0x7c, 0x1b, 0x32, 0xb4, 0x70, 0x2b, 0xb9, 0xe9, 0x6c, 0x02, 0xd1,
	0x9e, 0x1d, 0x63, 0xf0, 0x3e, 0x69, 0x63, 0x2b, 0x85, 0xa2, 0x3d, 0xe2, 0x15, 0xf8, 0xeb, 0xb2,
	0x32, 0x64, 0xb9, 0x28, 0x3f, 0xd9, 0x0f, 0xb5, 0xa4, 0xd8, 0xea, 0xf7, 0x9d, 0x3b, 0x76, 0x7d,
	0x03, 0x5d, 0x57, 0x05, 0x8f, 0x20, 0x9c, 0x4f, 0x47, 0xbc, 0x18, 0x4f, 0xf9, 0x28, 0x3e, 0xc0,
	0x1e, 0x1c, 0xce, 0xc4, 0x9c, 0xc7, 0x0c, 0x43, 0xf0, 0x8b, 0xc7, 0xc9, 0x0b, 0x8f, 0x3b, 0x6f,
	0x81, 0xfd, 0xe2, 0xdb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xeb, 0x63, 0xd3, 0x75, 0x01,
	0x00, 0x00,
}
