package encoder

import (
	"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/cmd/global/protocal/connect"
	"reflect"
	"strings"
	"github.com/weseliu/alphanet/cmd/global/protocal/game"
)

var scPack connect.SCPACK
var scMsg game.SCMSG
var logicPack connect.CMD_LOGIC_SC

func EncodeCmd(cmd interface{}) []byte {
	typ := reflect.TypeOf(cmd)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	cmdName := strings.TrimRight(typ.Name(),"_SC")
	scPack.Cmd = (connect.CommandType)(connect.CommandType_value[cmdName])
	data, err := codec.CodecManager().Encode("pb", cmd)
	if err == nil {
		scPack.Body = data
		data, err := codec.CodecManager().Encode("pb", &scPack)
		if err == nil{
			return data
		}
	}
	return nil
}

func EncodeMsg(code int32, msg interface{}) []byte {
	typ := reflect.TypeOf(msg)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	msgName := strings.TrimRight(typ.Name(),"_SC")
	scMsg.RetCode = code
	scMsg.Msg = (game.MessageType)(game.MessageType_value[msgName])
	data, err := codec.CodecManager().Encode("pb", msg)
	if err == nil{
		scMsg.Body = data
		data, err := codec.CodecManager().Encode("pb", &scMsg)
		if err == nil{
			return data
		}
	}

	return nil
}

func DecodeCmd(data []byte) interface{} {
	msg, err := codec.CodecManager().Decode("pb", data, &connect.CSPACK{})
	if err == nil {
		csPack := msg.(*connect.CSPACK)
		command := codec.BuildMessage(connect.CommandType_name[(int32)(csPack.Cmd)] + "_CS")
		command, err = codec.CodecManager().Decode("pb", csPack.Body, command)

		if err == nil {
			return command
		}
	}
	return nil
}

func DecodeMsg(data []byte) interface{} {
	csMsg, err := codec.CodecManager().Decode("pb", data, &game.CSMSG{})
	if err == nil{
		csPack := csMsg.(*game.CSMSG)
		msg := codec.BuildMessage(connect.CommandType_name[(int32)(csPack.Msg)] + "_CS")
		msg, err = codec.CodecManager().Decode("pb", csPack.Body, msg)
		if err == nil {
			return msg
		}
	}
	return nil
}