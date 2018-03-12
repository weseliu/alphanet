// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commandproto.proto

/*
Package command is a generated protocol buffer package.

It is generated from these files:
	commandproto.proto

It has these top-level messages:
	CMD_BASE_CS
	CMD_BASE_SC
	CMD_ERROR_SC
	CMD_AUTH_CS
	CMD_AUTH_SC
	CMD_REPLAY_CS
	CMD_REPLAY_SC
	CMD_LOGIC_CS
	CMD_LOGIC_SC
*/
package command

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

type CommandType int32

const (
	CommandType_CMD_ZERO_PLACEHOLDER CommandType = 0
	CommandType_CMD_ERROR            CommandType = 1
	CommandType_CMD_AUTH             CommandType = 16
	CommandType_CMD_REPLAY           CommandType = 18
	CommandType_CMD_LOGIC            CommandType = 96
)

var CommandType_name = map[int32]string{
	0:  "CMD_ZERO_PLACEHOLDER",
	1:  "CMD_ERROR",
	16: "CMD_AUTH",
	18: "CMD_REPLAY",
	96: "CMD_LOGIC",
}
var CommandType_value = map[string]int32{
	"CMD_ZERO_PLACEHOLDER": 0,
	"CMD_ERROR":            1,
	"CMD_AUTH":             16,
	"CMD_REPLAY":           18,
	"CMD_LOGIC":            96,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CMD_BASE_CS struct {
	Cmd  CommandType `protobuf:"varint,1,opt,name=cmd,enum=command.CommandType" json:"cmd,omitempty"`
	Body []byte      `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *CMD_BASE_CS) Reset()                    { *m = CMD_BASE_CS{} }
func (m *CMD_BASE_CS) String() string            { return proto.CompactTextString(m) }
func (*CMD_BASE_CS) ProtoMessage()               {}
func (*CMD_BASE_CS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CMD_BASE_CS) GetCmd() CommandType {
	if m != nil {
		return m.Cmd
	}
	return CommandType_CMD_ZERO_PLACEHOLDER
}

func (m *CMD_BASE_CS) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type CMD_BASE_SC struct {
	Cmd     CommandType `protobuf:"varint,1,opt,name=cmd,enum=command.CommandType" json:"cmd,omitempty"`
	RetCode int32       `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Body    []byte      `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *CMD_BASE_SC) Reset()                    { *m = CMD_BASE_SC{} }
func (m *CMD_BASE_SC) String() string            { return proto.CompactTextString(m) }
func (*CMD_BASE_SC) ProtoMessage()               {}
func (*CMD_BASE_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CMD_BASE_SC) GetCmd() CommandType {
	if m != nil {
		return m.Cmd
	}
	return CommandType_CMD_ZERO_PLACEHOLDER
}

func (m *CMD_BASE_SC) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CMD_BASE_SC) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type CMD_ERROR_SC struct {
	RetCode     int32  `protobuf:"varint,1,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	RetCodeDesc string `protobuf:"bytes,2,opt,name=ret_code_desc,json=retCodeDesc" json:"ret_code_desc,omitempty"`
}

func (m *CMD_ERROR_SC) Reset()                    { *m = CMD_ERROR_SC{} }
func (m *CMD_ERROR_SC) String() string            { return proto.CompactTextString(m) }
func (*CMD_ERROR_SC) ProtoMessage()               {}
func (*CMD_ERROR_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CMD_ERROR_SC) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CMD_ERROR_SC) GetRetCodeDesc() string {
	if m != nil {
		return m.RetCodeDesc
	}
	return ""
}

type CMD_AUTH_CS struct {
	GameId        int32  `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	IdentityToken string `protobuf:"bytes,2,opt,name=identity_token,json=identityToken" json:"identity_token,omitempty"`
}

func (m *CMD_AUTH_CS) Reset()                    { *m = CMD_AUTH_CS{} }
func (m *CMD_AUTH_CS) String() string            { return proto.CompactTextString(m) }
func (*CMD_AUTH_CS) ProtoMessage()               {}
func (*CMD_AUTH_CS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CMD_AUTH_CS) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *CMD_AUTH_CS) GetIdentityToken() string {
	if m != nil {
		return m.IdentityToken
	}
	return ""
}

type CMD_AUTH_SC struct {
	GameId      int32  `protobuf:"varint,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	RetCode     int32  `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	RetCodeDesc string `protobuf:"bytes,3,opt,name=ret_code_desc,json=retCodeDesc" json:"ret_code_desc,omitempty"`
	SessionId   string `protobuf:"bytes,4,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
}

func (m *CMD_AUTH_SC) Reset()                    { *m = CMD_AUTH_SC{} }
func (m *CMD_AUTH_SC) String() string            { return proto.CompactTextString(m) }
func (*CMD_AUTH_SC) ProtoMessage()               {}
func (*CMD_AUTH_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CMD_AUTH_SC) GetGameId() int32 {
	if m != nil {
		return m.GameId
	}
	return 0
}

func (m *CMD_AUTH_SC) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CMD_AUTH_SC) GetRetCodeDesc() string {
	if m != nil {
		return m.RetCodeDesc
	}
	return ""
}

func (m *CMD_AUTH_SC) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type CMD_REPLAY_CS struct {
	AuthChannel int32  `protobuf:"varint,1,opt,name=auth_channel,json=authChannel" json:"auth_channel,omitempty"`
	Account     string `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
	SessionId   string `protobuf:"bytes,3,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
}

func (m *CMD_REPLAY_CS) Reset()                    { *m = CMD_REPLAY_CS{} }
func (m *CMD_REPLAY_CS) String() string            { return proto.CompactTextString(m) }
func (*CMD_REPLAY_CS) ProtoMessage()               {}
func (*CMD_REPLAY_CS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CMD_REPLAY_CS) GetAuthChannel() int32 {
	if m != nil {
		return m.AuthChannel
	}
	return 0
}

func (m *CMD_REPLAY_CS) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CMD_REPLAY_CS) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type CMD_REPLAY_SC struct {
	RetCode     int32  `protobuf:"varint,1,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	RetCodeDesc string `protobuf:"bytes,2,opt,name=ret_code_desc,json=retCodeDesc" json:"ret_code_desc,omitempty"`
}

func (m *CMD_REPLAY_SC) Reset()                    { *m = CMD_REPLAY_SC{} }
func (m *CMD_REPLAY_SC) String() string            { return proto.CompactTextString(m) }
func (*CMD_REPLAY_SC) ProtoMessage()               {}
func (*CMD_REPLAY_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CMD_REPLAY_SC) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *CMD_REPLAY_SC) GetRetCodeDesc() string {
	if m != nil {
		return m.RetCodeDesc
	}
	return ""
}

type CMD_LOGIC_CS struct {
	LogicType uint32 `protobuf:"varint,1,opt,name=logic_type,json=logicType" json:"logic_type,omitempty"`
	LogicPkg  []byte `protobuf:"bytes,2,opt,name=logic_pkg,json=logicPkg,proto3" json:"logic_pkg,omitempty"`
}

func (m *CMD_LOGIC_CS) Reset()                    { *m = CMD_LOGIC_CS{} }
func (m *CMD_LOGIC_CS) String() string            { return proto.CompactTextString(m) }
func (*CMD_LOGIC_CS) ProtoMessage()               {}
func (*CMD_LOGIC_CS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CMD_LOGIC_CS) GetLogicType() uint32 {
	if m != nil {
		return m.LogicType
	}
	return 0
}

func (m *CMD_LOGIC_CS) GetLogicPkg() []byte {
	if m != nil {
		return m.LogicPkg
	}
	return nil
}

type CMD_LOGIC_SC struct {
	LogicType uint32 `protobuf:"varint,1,opt,name=logic_type,json=logicType" json:"logic_type,omitempty"`
	LogicPkg  []byte `protobuf:"bytes,2,opt,name=logic_pkg,json=logicPkg,proto3" json:"logic_pkg,omitempty"`
}

func (m *CMD_LOGIC_SC) Reset()                    { *m = CMD_LOGIC_SC{} }
func (m *CMD_LOGIC_SC) String() string            { return proto.CompactTextString(m) }
func (*CMD_LOGIC_SC) ProtoMessage()               {}
func (*CMD_LOGIC_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CMD_LOGIC_SC) GetLogicType() uint32 {
	if m != nil {
		return m.LogicType
	}
	return 0
}

func (m *CMD_LOGIC_SC) GetLogicPkg() []byte {
	if m != nil {
		return m.LogicPkg
	}
	return nil
}

func init() {
	proto.RegisterType((*CMD_BASE_CS)(nil), "command.CMD_BASE_CS")
	proto.RegisterType((*CMD_BASE_SC)(nil), "command.CMD_BASE_SC")
	proto.RegisterType((*CMD_ERROR_SC)(nil), "command.CMD_ERROR_SC")
	proto.RegisterType((*CMD_AUTH_CS)(nil), "command.CMD_AUTH_CS")
	proto.RegisterType((*CMD_AUTH_SC)(nil), "command.CMD_AUTH_SC")
	proto.RegisterType((*CMD_REPLAY_CS)(nil), "command.CMD_REPLAY_CS")
	proto.RegisterType((*CMD_REPLAY_SC)(nil), "command.CMD_REPLAY_SC")
	proto.RegisterType((*CMD_LOGIC_CS)(nil), "command.CMD_LOGIC_CS")
	proto.RegisterType((*CMD_LOGIC_SC)(nil), "command.CMD_LOGIC_SC")
	proto.RegisterEnum("command.CommandType", CommandType_name, CommandType_value)
}

func init() { proto.RegisterFile("commandproto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x09, 0x1d, 0x6b, 0xfb, 0xda, 0x54, 0x91, 0x35, 0x89, 0x22, 0x34, 0x69, 0x44, 0x02,
	0x4d, 0x1c, 0x7a, 0x80, 0xbf, 0xa0, 0xb8, 0x11, 0x0b, 0x4a, 0x49, 0xe5, 0x94, 0x03, 0x5c, 0x4c,
	0x66, 0x5b, 0x59, 0x94, 0x25, 0x8e, 0x1a, 0xef, 0x90, 0x3b, 0x7f, 0x38, 0xb2, 0xe3, 0x8c, 0x74,
	0x13, 0x08, 0x69, 0x97, 0x36, 0xfe, 0xfc, 0xfc, 0xfb, 0xbe, 0x17, 0xbf, 0x00, 0x62, 0xb2, 0x2c,
	0xd3, 0x8a, 0xd7, 0x07, 0xa9, 0xe4, 0xca, 0xfc, 0xa2, 0xb1, 0xd5, 0xfc, 0x10, 0x66, 0x78, 0xbb,
	0xa1, 0x9f, 0xd6, 0x49, 0x40, 0x71, 0x82, 0xde, 0xc1, 0x88, 0x95, 0x7c, 0xe9, 0x5c, 0x38, 0x97,
	0x8b, 0x0f, 0x67, 0x2b, 0x5b, 0xb5, 0xc2, 0xdd, 0xff, 0xbe, 0xad, 0x05, 0xd1, 0x05, 0x08, 0xc1,
	0xc9, 0xb5, 0xe4, 0xed, 0xf2, 0xf9, 0x85, 0x73, 0x39, 0x27, 0xe6, 0xd9, 0xe7, 0x03, 0x54, 0x82,
	0xff, 0x1b, 0xf5, 0x0a, 0x26, 0x07, 0xa1, 0x28, 0x93, 0x5c, 0x18, 0xdc, 0x0b, 0x32, 0x3e, 0x08,
	0x85, 0x25, 0x17, 0xf7, 0x2e, 0xa3, 0x81, 0xcb, 0x16, 0xe6, 0xda, 0x25, 0x20, 0x24, 0x26, 0xda,
	0x66, 0x78, 0xdc, 0x39, 0x3e, 0xee, 0x83, 0xdb, 0x6f, 0x51, 0x2e, 0x1a, 0x66, 0xf0, 0x53, 0x32,
	0xb3, 0xfb, 0x1b, 0xd1, 0x30, 0x7f, 0xdb, 0x85, 0x5e, 0x7f, 0xdb, 0x5f, 0xe9, 0xfe, 0x5f, 0xc2,
	0x38, 0x4b, 0x4b, 0x41, 0x73, 0x6e, 0x61, 0xa7, 0x7a, 0x19, 0x72, 0xf4, 0x16, 0x16, 0x39, 0x17,
	0x95, 0xca, 0x55, 0x4b, 0x95, 0x2c, 0x44, 0x65, 0x61, 0x6e, 0xaf, 0xee, 0xb5, 0xe8, 0xff, 0x72,
	0x06, 0xbc, 0x04, 0xff, 0x9d, 0xf7, 0x8f, 0xae, 0x1f, 0xc5, 0x1e, 0x3d, 0x8a, 0x8d, 0xce, 0x01,
	0x1a, 0xd1, 0x34, 0xb9, 0xac, 0x34, 0xfa, 0xc4, 0x14, 0x4c, 0xad, 0x12, 0x72, 0xbf, 0x00, 0x57,
	0xa7, 0x20, 0xc1, 0x2e, 0x5a, 0x7f, 0xd7, 0x7d, 0xbd, 0x81, 0x79, 0x7a, 0xa7, 0x6e, 0x28, 0xbb,
	0x49, 0xab, 0x4a, 0xdc, 0xda, 0x30, 0x33, 0xad, 0xe1, 0x4e, 0x42, 0x4b, 0x18, 0xa7, 0x8c, 0xc9,
	0xbb, 0x4a, 0xd9, 0xd6, 0xfa, 0xe5, 0x03, 0xb3, 0xd1, 0x43, 0xb3, 0xaf, 0x47, 0x66, 0x4f, 0xbf,
	0x92, 0x2f, 0xdd, 0x0d, 0x47, 0xf1, 0xe7, 0x10, 0xeb, 0xec, 0xe7, 0x00, 0xb7, 0x32, 0xcb, 0x19,
	0x55, 0x6d, 0xdd, 0x01, 0x5d, 0x32, 0x35, 0x8a, 0x1e, 0x22, 0xf4, 0x1a, 0xba, 0x05, 0xad, 0x8b,
	0xcc, 0xce, 0xe3, 0xc4, 0x08, 0xbb, 0x22, 0x3b, 0x66, 0x25, 0xf8, 0x29, 0xac, 0xf7, 0x0c, 0x66,
	0x83, 0xe1, 0x45, 0x4b, 0x38, 0xd3, 0xe8, 0x1f, 0x01, 0x89, 0xe9, 0x2e, 0x5a, 0xe3, 0xe0, 0x2a,
	0x8e, 0x36, 0x01, 0xf1, 0x9e, 0x21, 0x17, 0xa6, 0xf7, 0x23, 0xea, 0x39, 0x68, 0x0e, 0x93, 0x7e,
	0x24, 0x3c, 0x0f, 0x2d, 0x00, 0xfe, 0xbc, 0x2d, 0x0f, 0xf5, 0xc5, 0x26, 0xa1, 0xf7, 0xf3, 0xfa,
	0xd4, 0x7c, 0x9f, 0x1f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xc7, 0x89, 0xb0, 0xb5, 0x03,
	0x00, 0x00,
}
