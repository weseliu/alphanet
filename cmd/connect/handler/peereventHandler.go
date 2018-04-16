package handler

import (
	"github.com/weseliu/alphanet/core"
	"github.com/weseliu/alphanet/net"
	"log"
	"reflect"
	"github.com/weseliu/alphanet/cmd/global/protocal/connect"
	"github.com/weseliu/alphanet/cmd/global/interfaces"
	"time"
	"github.com/weseliu/alphanet/cmd/global/ado"
	"fmt"
	"github.com/weseliu/alphanet/cmd/global/encoder"
)

type EventHandler struct {
}

func (Self *EventHandler) Start(peer net.Peer) {
	peer.RegisterEventHandler(net.EventReceive, Self.onEventReceive)
	peer.RegisterEventHandler(net.EventConnected, Self.onEventConnected)
	peer.RegisterEventHandler(net.EventClosed, Self.onEventClosed)

	core.MsgCenter().RegisterMsgHandler(reflect.TypeOf((*connect.CMD_AUTH_CS)(nil)), Self)
	core.MsgCenter().RegisterMsgHandler(reflect.TypeOf((*connect.CMD_LOGIC_CS)(nil)), Self)
}

func (Self *EventHandler) onEventReceive(event *net.Event){
	log.Print("EventReceive : ", string(event.Data))
	var msg = encoder.DecodeCmd(event.Data)
	if msg != nil {
		core.MsgCenter().DispatchMessage(msg, event.Session)
	}
}

func (Self *EventHandler) onEventConnected(event *net.Event){
	log.Print("EventConnected !")
	event.Session.SetID(0)
}

func (Self *EventHandler) onEventClosed(event *net.Event){
	OnRoleExit(event.Session)
}

func (Self *EventHandler) OnMessage(msg interface{}, param interface{}) {
	session := param.(net.Session)
	if _, ok := msg.(*connect.CMD_AUTH_CS); ok {
		Self.onUserAuth(session, msg)
	}

	if _, ok := msg.(*connect.CMD_LOGIC_CS); ok {
		Self.onLogicMessage(session, msg)
	}
}

func (Self *EventHandler) onUserAuth(session net.Session, msg interface{}) {
	var authReq = msg.(*connect.CMD_AUTH_CS)

	token := &interfaces.IdentityToken{}
	if err := token.Decrypt(authReq.IdentityToken); err != nil{
		log.Println(err)
		return
	}

	role := ado.Role().GetRole(token.Id)
	if role == nil {
		role = &ado.RoleModel{
			Id : token.Id,
			Name : fmt.Sprint("role", token.Id),
			Age : 10,
		}
		if ado.Role().AddRole(role) != true{
			role = nil
		}
	}

	authRsq := &connect.CMD_AUTH_SC{}
	authRsq.GameId = authReq.GameId
	if role == nil {
		authRsq.RetCode = -1
		authRsq.RetCodeDesc = "add role fail!"
	} else {
		sessionId := &interfaces.SessionId{
			Id : role.Id,
			Time :time.Now(),
		}
		authRsq.RetCode = 0
		var err error = nil
		authRsq.SessionId, err = sessionId.Encrypt()
		if err != nil {
			log.Println(err)
		}
	}

	if data := encoder.EncodeCmd(authRsq); data != nil{
		session.Send(data)
	}

	session.SetID(role.Id)
	OnRoleEnter(role.Id, session)
}

func (Self *EventHandler) onLogicMessage(session net.Session, msg interface{}) {
	var logicMsg = msg.(*connect.CMD_LOGIC_CS)
	SendRemoteMessage(session.ID(), logicMsg.LogicPkg)
}


