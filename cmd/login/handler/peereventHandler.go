package handler

import (
	"github.com/weseliu/alphanet/core"
	"github.com/weseliu/alphanet/net"
	"log"
	"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/cmd/protocal/connect"
)

type EventHandler struct {
}

func (Self *EventHandler) Start(peer net.Peer) {
	peer.RegisterEventHandler(net.EventReceive, Self.onEventReceive)
	peer.RegisterEventHandler(net.EventConnected, Self.onEventConnected)
	peer.RegisterEventHandler(net.EventClosed, Self.onEventClosed)
}

func (Self *EventHandler) onEventReceive(event *net.Event){
	log.Print("EventReceive : ", string(event.Data))
	msg, err := codec.CodecManager().Decode("pb", event.Data, &connect.CSPACK{})
	if err == nil {
		csPack := msg.(*connect.CSPACK)
		command := codec.BuildMessage(connect.CommandType_name[(int32)(csPack.Cmd)] + "_CS")
		command, err = codec.CodecManager().Decode("pb", csPack.Body, command)

		if csPack.Cmd == (connect.CommandType)(connect.CommandType_value["CMD_LOGIC"]){
			//csMsg, err := codec.CodecManager().Decode("pb", command.(connect.CMD_LOGIC_CS).LogicPkg, &connect.CSMSG{})
			//if err == nil{
			//
			//}
		} else {
			core.MsgCenter().DispatchMessage(event.Session, command)
		}
	} else {
		log.Fatal(err)
	}
}

func (Self *EventHandler) onEventConnected(event *net.Event){
	log.Print("EventConnected !")
}

func (Self *EventHandler) onEventClosed(event *net.Event){
	log.Print("EventClosed !")
}
