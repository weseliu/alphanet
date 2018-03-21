package core

import (
	"reflect"
	"github.com/weseliu/alphanet/util"
	"github.com/weseliu/alphanet/net"
)

type MsgHandler interface {
	OnMessage(session net.Session, msg interface{})
}

type msgCenter struct {
	msgHandlers map[reflect.Type] *util.Set
}

var msgCenterInstance *msgCenter = nil
func MsgCenter() *msgCenter  {
	if msgCenterInstance == nil{
		msgCenterInstance = &msgCenter{
			msgHandlers : make(map[reflect.Type] *util.Set),
		}
	}
	return msgCenterInstance
}

func (Self *msgCenter)RegisterMsgHandler(typ reflect.Type, handler interface{}) {
	if typ.Kind() == reflect.Ptr{
		typ = typ.Elem()
	}

	if Self.msgHandlers[typ] == nil {
		Self.msgHandlers[typ] = util.NewSet()
	}

	if handler.(MsgHandler) != nil{
		Self.msgHandlers[typ].Add(handler)
	}
}

func (Self *msgCenter)UnRegisterMsgHandler(typ reflect.Type, handler interface{})  {
	if typ.Kind() == reflect.Ptr{
		typ = typ.Elem()
	}

	if Self.msgHandlers[typ] != nil {
		Self.msgHandlers[typ].Remove(handler)
	}
}

func (Self *msgCenter)UnRegisterAllMsgHandler(handler interface{})  {
	for key := range Self.msgHandlers {
		Self.UnRegisterMsgHandler(key, handler)
	}
}

func (Self *msgCenter)DispatchMessage(session net.Session, msg interface{}) {
	typ := reflect.TypeOf(msg)
	if typ.Kind() == reflect.Ptr{
		typ = typ.Elem()
	}

	if Self.msgHandlers[typ] != nil {
		Self.msgHandlers[typ].Traverse(func(item interface{}) bool {
			var handler = item.(MsgHandler)
			if handler != nil {
				handler.OnMessage(session, msg)
			}
			return true
		})
	}
}
