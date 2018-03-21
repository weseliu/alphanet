package core

import (
	"reflect"
	"github.com/weseliu/alphanet/util"
)

type MsgHandler interface {
	OnMessage(msg interface{})
}

type msgCenter struct {
	msgHandlers map[reflect.Type] *util.Set
}

var msgCenterInstance *msgCenter = nil
func MsgCenter() *msgCenter  {
	if msgCenterInstance == nil{
		msgCenterInstance = &msgCenter{}
	}
	return msgCenterInstance
}

func (Self *msgCenter)RegisterMsgHandler(typ reflect.Type, handler interface{}) {
	if Self.msgHandlers[typ] == nil {
		Self.msgHandlers[typ] = util.NewSet()
	}

	if handler.(MsgHandler) != nil{
		Self.msgHandlers[typ].Add(handler)
	}
}

func (Self *msgCenter)UnRegisterMsgHandler(typ reflect.Type, handler interface{})  {
	if Self.msgHandlers[typ] != nil {
		Self.msgHandlers[typ].Remove(handler)
	}
}

func (Self *msgCenter)UnRegisterAllMsgHandler(handler interface{})  {
	for key := range Self.msgHandlers {
		Self.UnRegisterMsgHandler(key, handler)
	}
}

func (Self *msgCenter)DispatchMessage(msg interface{}) {
	typ := reflect.TypeOf(msg)
	if typ.Kind() == reflect.Ptr{
		typ = typ.Elem()
	}

	if Self.msgHandlers[typ] != nil {
		Self.msgHandlers[typ].Traverse(func(item interface{}) bool {
			var handler = item.(MsgHandler)
			if handler != nil {
				handler.OnMessage(msg)
			}
			return true
		})
	}
}
