package handler

import (
	"github.com/weseliu/alphanet/core"
	"github.com/weseliu/alphanet/net"
	"log"
	"github.com/weseliu/alphanet/cmd/login/encoder"
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
	var msg = encoder.Decode(event.Data)
	if msg != nil {
		core.MsgCenter().DispatchMessage(event.Session, msg)
	}
}

func (Self *EventHandler) onEventConnected(event *net.Event){
	log.Print("EventConnected !")
}

func (Self *EventHandler) onEventClosed(event *net.Event){
	log.Print("EventClosed !")
}
