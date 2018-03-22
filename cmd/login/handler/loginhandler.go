package handler

import (
	"github.com/weseliu/alphanet/net"
	"github.com/weseliu/alphanet/core"
	"reflect"
	"github.com/weseliu/alphanet/cmd/login/protocal"
	"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/cmd/login/ado"
)

type LoginHandler struct {

}

func (Self *LoginHandler)Start() {
	core.MsgCenter().RegisterMsgHandler(reflect.TypeOf((*protocal.UserAuth)(nil)), Self)
}

func (Self *LoginHandler)OnMessage(session net.Session, msg interface{}) {
	if msg.(*protocal.UserAuth) != nil {
		Self.onUserAuth(session, msg)
	}
}

func (Self *LoginHandler)onUserAuth(session net.Session, msg interface{})  {
	var userAuth = msg.(*protocal.UserAuth)

	var authResult = &protocal.AuthResult{
		Ret : 0,
		Msg : "",
		Token : userAuth.Name + " : " + userAuth.Password,
	}

	var user = ado.User().GetUser(userAuth.Id)
	if user != nil {
		if user.Password != userAuth.Password {
			authResult.Ret = -1
			authResult.Msg = "password error!"
		}
	} else {
		authResult.Ret = 0
		authResult.Msg = "user is not exist!"
	}

	data, err := codec.CodecManager().Encode("json", authResult)
	if err == nil {
		session.Send(data)
	}
}