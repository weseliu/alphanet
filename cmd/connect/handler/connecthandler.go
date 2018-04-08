package handler

import (
	"github.com/weseliu/alphanet/core"
	"github.com/weseliu/alphanet/net"
	"reflect"
	"github.com/weseliu/alphanet/cmd/protocal/connect"
	"github.com/weseliu/alphanet/cmd/connect/encoder"
)

type LoginHandler struct {
}

func (Self *LoginHandler) Start() {
	core.MsgCenter().RegisterMsgHandler(reflect.TypeOf((*connect.CMD_AUTH_CS)(nil)), Self)
}

func (Self *LoginHandler) OnMessage(session net.Session, msg interface{}) {
	if msg.(*connect.CMD_AUTH_CS) != nil {
		Self.onUserAuth(session, msg)
	}
}

func (Self *LoginHandler) onUserAuth(session net.Session, msg interface{}) {
	var authReq = msg.(*connect.CMD_AUTH_CS)
	authRsq := &connect.CMD_AUTH_SC{}
	authRsq.GameId = authReq.GameId
	authRsq.RetCode = 0
	authRsq.SessionId = "12222222222222111111"
	if data := encoder.EncodeCmd(authRsq); data != nil{
		session.Send(data)
	}
}
