package logic

import (
	"github.com/weseliu/alphanet/core"
	"reflect"
	"github.com/weseliu/alphanet/cmd/global/protocal/game"
	"log"
	"github.com/weseliu/alphanet/cmd/game/handler"
)

type roleManager struct {
	
}

var roleManagerInstance *roleManager
func RoleManager() *roleManager {
	if roleManagerInstance == nil {
		roleManagerInstance = &roleManager{}
	}
	return roleManagerInstance
}

func (Self *roleManager)Start()  {
	core.MsgCenter().RegisterMsgHandler(reflect.TypeOf((*game.MSG_LOGIN_CS)(nil)), Self)
}

func (Self *roleManager)End()  {
	core.MsgCenter().UnRegisterMsgHandler(reflect.TypeOf((*game.MSG_LOGIN_CS)(nil)), Self)
}

func (Self *roleManager) OnMessage(msg interface{}, param interface{}) {
	roleId := param.(int64)
	if ptr, ok := msg.(*game.MSG_LOGIN_CS); ok {
		Self.onRoleLogin(roleId, ptr)
	}
}

func (Self *roleManager) onRoleLogin(roleId int64, msg *game.MSG_LOGIN_CS) {
	log.Println("onRoleLogin : ", roleId)
	rsp := &game.MSG_LOGIN_SC{
		RetCode: 0,
		RetMsg:[]byte("login success!"),
	}
	handler.SendMessage(roleId, rsp, 0)
}