// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

/*
Package game is a generated protocol buffer package.

It is generated from these files:
	game.proto

It has these top-level messages:
	CSMSG
	SCMSG
*/
package game

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

type MessageType int32

const (
	MessageType_MSG_ZERO_PLACEHOLDER MessageType = 0
	MessageType_MSG_ERROR            MessageType = 1
	MessageType_MSG_PING             MessageType = 2
	MessageType_MSG_LOGIN            MessageType = 3
	MessageType_MSG_KICKOUT          MessageType = 4
)

var MessageType_name = map[int32]string{
	0: "MSG_ZERO_PLACEHOLDER",
	1: "MSG_ERROR",
	2: "MSG_PING",
	3: "MSG_LOGIN",
	4: "MSG_KICKOUT",
}
var MessageType_value = map[string]int32{
	"MSG_ZERO_PLACEHOLDER": 0,
	"MSG_ERROR":            1,
	"MSG_PING":             2,
	"MSG_LOGIN":            3,
	"MSG_KICKOUT":          4,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CSMSG struct {
	Msg  MessageType `protobuf:"varint,1,opt,name=msg,enum=game.MessageType" json:"msg,omitempty"`
	Body []byte      `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *CSMSG) Reset()                    { *m = CSMSG{} }
func (m *CSMSG) String() string            { return proto.CompactTextString(m) }
func (*CSMSG) ProtoMessage()               {}
func (*CSMSG) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CSMSG) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_MSG_ZERO_PLACEHOLDER
}

func (m *CSMSG) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type SCMSG struct {
	Msg     MessageType `protobuf:"varint,1,opt,name=msg,enum=game.MessageType" json:"msg,omitempty"`
	RetCode int32       `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Body    []byte      `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *SCMSG) Reset()                    { *m = SCMSG{} }
func (m *SCMSG) String() string            { return proto.CompactTextString(m) }
func (*SCMSG) ProtoMessage()               {}
func (*SCMSG) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SCMSG) GetMsg() MessageType {
	if m != nil {
		return m.Msg
	}
	return MessageType_MSG_ZERO_PLACEHOLDER
}

func (m *SCMSG) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *SCMSG) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*CSMSG)(nil), "game.CSMSG")
	proto.RegisterType((*SCMSG)(nil), "game.SCMSG")
	proto.RegisterEnum("game.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("game.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x1c, 0xb8, 0x58, 0x9d, 0x83,
	0x7d, 0x83, 0xdd, 0x85, 0x94, 0xb9, 0x98, 0x73, 0x8b, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0xf8,
	0x8c, 0x04, 0xf5, 0xc0, 0x0a, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x43, 0x2a, 0x0b, 0x52,
	0x83, 0x40, 0xb2, 0x42, 0x42, 0x5c, 0x2c, 0x49, 0xf9, 0x29, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0x3c, 0x41, 0x60, 0xb6, 0x52, 0x34, 0x17, 0x6b, 0xb0, 0x33, 0xd1, 0x26, 0x48, 0x72, 0x71, 0x14,
	0xa5, 0x96, 0xc4, 0x27, 0xe7, 0xa7, 0xa4, 0x82, 0x4d, 0x61, 0x0d, 0x62, 0x2f, 0x4a, 0x2d, 0x71,
	0xce, 0x4f, 0x49, 0x85, 0x1b, 0xce, 0x8c, 0x30, 0x5c, 0x2b, 0x85, 0x8b, 0x1b, 0xc9, 0x08, 0x21,
	0x09, 0x2e, 0x11, 0xdf, 0x60, 0xf7, 0xf8, 0x28, 0xd7, 0x20, 0xff, 0xf8, 0x00, 0x1f, 0x47, 0x67,
	0x57, 0x0f, 0x7f, 0x1f, 0x17, 0xd7, 0x20, 0x01, 0x06, 0x21, 0x5e, 0x2e, 0x4e, 0x90, 0x8c, 0x6b,
	0x50, 0x90, 0x7f, 0x90, 0x00, 0xa3, 0x10, 0x0f, 0x17, 0x07, 0x88, 0x1b, 0xe0, 0xe9, 0xe7, 0x2e,
	0xc0, 0x04, 0x93, 0xf4, 0xf1, 0x77, 0xf7, 0xf4, 0x13, 0x60, 0x16, 0xe2, 0xe7, 0xe2, 0x06, 0x71,
	0xbd, 0x3d, 0x9d, 0xbd, 0xfd, 0x43, 0x43, 0x04, 0x58, 0x92, 0xd8, 0xc0, 0x21, 0x62, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x2b, 0x91, 0x97, 0x37, 0x1f, 0x01, 0x00, 0x00,
}