// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connect.proto

/*
Package connect is a generated protocol buffer package.

It is generated from these files:
	connect.proto

It has these top-level messages:
	CSPACK
	SCPACK
	CMD_ERROR_SC
	CMD_AUTH_CS
	CMD_AUTH_SC
	CMD_REPLAY_CS
	CMD_REPLAY_SC
	CMD_LOGIC_CS
	CMD_LOGIC_SC
*/
package connect

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
	CommandType_CMD_AUTH             CommandType = 2
	CommandType_CMD_REPLAY           CommandType = 3
	CommandType_CMD_LOGIC            CommandType = 4
)

var CommandType_name = map[int32]string{
	0: "CMD_ZERO_PLACEHOLDER",
	1: "CMD_ERROR",
	2: "CMD_AUTH",
	3: "CMD_REPLAY",
	4: "CMD_LOGIC",
}
var CommandType_value = map[string]int32{
	"CMD_ZERO_PLACEHOLDER": 0,
	"CMD_ERROR":            1,
	"CMD_AUTH":             2,
	"CMD_REPLAY":           3,
	"CMD_LOGIC":            4,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CSPACK struct {
	Cmd  CommandType `protobuf:"varint,1,opt,name=cmd,enum=connect.CommandType" json:"cmd,omitempty"`
	Body []byte      `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *CSPACK) Reset()                    { *m = CSPACK{} }
func (m *CSPACK) String() string            { return proto.CompactTextString(m) }
func (*CSPACK) ProtoMessage()               {}
func (*CSPACK) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CSPACK) GetCmd() CommandType {
	if m != nil {
		return m.Cmd
	}
	return CommandType_CMD_ZERO_PLACEHOLDER
}

func (m *CSPACK) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type SCPACK struct {
	Cmd     CommandType `protobuf:"varint,1,opt,name=cmd,enum=connect.CommandType" json:"cmd,omitempty"`
	RetCode int32       `protobuf:"varint,2,opt,name=ret_code,json=retCode" json:"ret_code,omitempty"`
	Body    []byte      `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *SCPACK) Reset()                    { *m = SCPACK{} }
func (m *SCPACK) String() string            { return proto.CompactTextString(m) }
func (*SCPACK) ProtoMessage()               {}
func (*SCPACK) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SCPACK) GetCmd() CommandType {
	if m != nil {
		return m.Cmd
	}
	return CommandType_CMD_ZERO_PLACEHOLDER
}

func (m *SCPACK) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *SCPACK) GetBody() []byte {
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
	LogicPkg []byte `protobuf:"bytes,2,opt,name=logic_pkg,json=logicPkg,proto3" json:"logic_pkg,omitempty"`
}

func (m *CMD_LOGIC_CS) Reset()                    { *m = CMD_LOGIC_CS{} }
func (m *CMD_LOGIC_CS) String() string            { return proto.CompactTextString(m) }
func (*CMD_LOGIC_CS) ProtoMessage()               {}
func (*CMD_LOGIC_CS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CMD_LOGIC_CS) GetLogicPkg() []byte {
	if m != nil {
		return m.LogicPkg
	}
	return nil
}

type CMD_LOGIC_SC struct {
	LogicPkg []byte `protobuf:"bytes,2,opt,name=logic_pkg,json=logicPkg,proto3" json:"logic_pkg,omitempty"`
}

func (m *CMD_LOGIC_SC) Reset()                    { *m = CMD_LOGIC_SC{} }
func (m *CMD_LOGIC_SC) String() string            { return proto.CompactTextString(m) }
func (*CMD_LOGIC_SC) ProtoMessage()               {}
func (*CMD_LOGIC_SC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CMD_LOGIC_SC) GetLogicPkg() []byte {
	if m != nil {
		return m.LogicPkg
	}
	return nil
}

func init() {
	proto.RegisterType((*CSPACK)(nil), "connect.CSPACK")
	proto.RegisterType((*SCPACK)(nil), "connect.SCPACK")
	proto.RegisterType((*CMD_ERROR_SC)(nil), "connect.CMD_ERROR_SC")
	proto.RegisterType((*CMD_AUTH_CS)(nil), "connect.CMD_AUTH_CS")
	proto.RegisterType((*CMD_AUTH_SC)(nil), "connect.CMD_AUTH_SC")
	proto.RegisterType((*CMD_REPLAY_CS)(nil), "connect.CMD_REPLAY_CS")
	proto.RegisterType((*CMD_REPLAY_SC)(nil), "connect.CMD_REPLAY_SC")
	proto.RegisterType((*CMD_LOGIC_CS)(nil), "connect.CMD_LOGIC_CS")
	proto.RegisterType((*CMD_LOGIC_SC)(nil), "connect.CMD_LOGIC_SC")
	proto.RegisterEnum("connect.CommandType", CommandType_name, CommandType_value)
}

func init() { proto.RegisterFile("connect.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x9d, 0xeb, 0x2c, 0x1f, 0x37, 0x71, 0x30, 0xa2, 0xb0, 0x8c, 0x31, 0xe8, 0x0c, 0x1b, 0x65,
	0x83, 0x3e, 0x6c, 0xbf, 0xc0, 0xc8, 0x66, 0x0d, 0x73, 0x66, 0x23, 0x67, 0x0f, 0xdb, 0x8b, 0x70,
	0x25, 0xe1, 0x1a, 0xd7, 0x52, 0x88, 0xd5, 0x87, 0xbc, 0xef, 0x87, 0x0f, 0xb9, 0x72, 0x17, 0xb7,
	0xac, 0x30, 0xfa, 0xe6, 0x7b, 0x74, 0xee, 0x39, 0xe7, 0x72, 0xaf, 0xc1, 0x63, 0x4a, 0x4a, 0xc1,
	0xf4, 0xc5, 0x6e, 0xaf, 0xb4, 0x42, 0x13, 0x5b, 0x06, 0x11, 0x8c, 0x71, 0x9e, 0x85, 0xf8, 0x1b,
	0xfa, 0x00, 0x2e, 0x6b, 0xf8, 0xca, 0x39, 0x73, 0xce, 0x97, 0x9f, 0x4f, 0x2f, 0x7a, 0x3e, 0x56,
	0x4d, 0x53, 0x48, 0xbe, 0x3d, 0xec, 0x04, 0x31, 0x04, 0x84, 0x60, 0x74, 0xa5, 0xf8, 0x61, 0x75,
	0x72, 0xe6, 0x9c, 0x2f, 0x48, 0xf7, 0x1d, 0x50, 0x18, 0xe7, 0xf8, 0xbf, 0x54, 0x5e, 0xc3, 0x74,
	0x2f, 0x34, 0x65, 0x8a, 0x8b, 0x4e, 0xe9, 0x25, 0x99, 0xec, 0x85, 0xc6, 0x8a, 0x8b, 0x7b, 0x03,
	0xf7, 0xc8, 0x60, 0x03, 0x0b, 0xbc, 0x89, 0x68, 0x4c, 0x48, 0x4a, 0x68, 0x8e, 0x07, 0xed, 0xce,
	0xb0, 0x3d, 0x00, 0xaf, 0x7f, 0xa2, 0x5c, 0xb4, 0xac, 0x93, 0x9f, 0x91, 0xb9, 0x7d, 0x8f, 0x44,
	0xcb, 0x82, 0x0d, 0xcc, 0x8d, 0x5c, 0xf8, 0x63, 0x7b, 0x49, 0x71, 0x8e, 0x5e, 0xc1, 0xa4, 0x2c,
	0x1a, 0x41, 0x2b, 0x6e, 0xc5, 0xc6, 0xa6, 0x5c, 0x73, 0xf4, 0x1e, 0x96, 0x15, 0x17, 0x52, 0x57,
	0xfa, 0x40, 0xb5, 0xaa, 0x85, 0xb4, 0x62, 0x5e, 0x8f, 0x6e, 0x0d, 0x18, 0xfc, 0x76, 0x8e, 0xf4,
	0x72, 0xfc, 0x6f, 0xbd, 0x27, 0xa6, 0x7e, 0x14, 0xdb, 0x7d, 0x14, 0x1b, 0xbd, 0x05, 0x68, 0x45,
	0xdb, 0x56, 0x4a, 0x1a, 0xe9, 0x51, 0x47, 0x98, 0x59, 0x64, 0xcd, 0x83, 0x1a, 0x3c, 0x93, 0x82,
	0xc4, 0x59, 0x12, 0xfe, 0x34, 0x73, 0xbd, 0x83, 0x45, 0x71, 0xab, 0xaf, 0x29, 0xbb, 0x2e, 0xa4,
	0x14, 0x37, 0x36, 0xcc, 0xdc, 0x60, 0xf8, 0x0e, 0x42, 0x2b, 0x98, 0x14, 0x8c, 0xa9, 0x5b, 0xa9,
	0xed, 0x68, 0x7d, 0xf9, 0xc0, 0xcc, 0x7d, 0x68, 0xf6, 0x7d, 0x60, 0xf6, 0xfc, 0x95, 0x7c, 0xba,
	0xdb, 0x70, 0x92, 0x7e, 0x5d, 0x63, 0x93, 0xfd, 0x0d, 0xcc, 0x6e, 0x54, 0x59, 0x31, 0xba, 0xab,
	0x4b, 0x7b, 0x6b, 0xd3, 0x0e, 0xc8, 0xea, 0x72, 0x48, 0xce, 0xf1, 0x93, 0xe4, 0x8f, 0x0c, 0xe6,
	0x47, 0xe7, 0x87, 0x56, 0x70, 0x6a, 0x7a, 0x7f, 0xc5, 0x24, 0xa5, 0x59, 0x12, 0xe2, 0xf8, 0x32,
	0x4d, 0xa2, 0x98, 0xf8, 0x2f, 0x90, 0x07, 0xb3, 0xfb, 0x23, 0xf3, 0x1d, 0xb4, 0x80, 0x69, 0xbf,
	0x54, 0xff, 0x04, 0x2d, 0x01, 0xfe, 0xce, 0xeb, 0xbb, 0x3d, 0xb9, 0x8b, 0xe0, 0x8f, 0xae, 0xc6,
	0xdd, 0x7f, 0xf5, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x12, 0x59, 0xd2, 0x68, 0x03,
	0x00, 0x00,
}
