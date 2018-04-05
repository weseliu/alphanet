package handler

import (
	"fmt"
	"github.com/weseliu/alphanet/cmd/login/ado"
	"github.com/weseliu/alphanet/cmd/login/protocal"
	"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/core"
	"github.com/weseliu/alphanet/net"
	"reflect"
	"github.com/weseliu/alphanet/cmd/protocal/connect"
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
	var userAuth = msg.(*protocal.UserAuth)

	var authResult = &protocal.AuthResult{
		Ret:   0,
		Msg:   "",
		Token: userAuth.Name + " : " + userAuth.Password,
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

		userId := ado.User().AddUser(&ado.UserModel{
			Name:     userAuth.Name,
			Password: userAuth.Password,
			Age:      10,
			Address:  userAuth.DeviceId,
		})
		authResult.Msg = fmt.Sprintf("register user success, id : %d", userId)
	}

	data, err := codec.CodecManager().Encode("json", authResult)
	if err == nil {
		session.Send(data)
	}
}
