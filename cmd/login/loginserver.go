package main

import (
	"github.com/weseliu/alphanet/net"
	"github.com/weseliu/alphanet/net/websocket"
	"log"
	"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/cmd/login/protocal"
)

func main() {
	queue := net.NewEventQueue()
	peer := net.PeerManager().NewPeer("login", queue, func(queue net.EventQueue) net.Peer {
		return websocket.NewAcceptor(queue)
	})

	peer.RegisterEvent(net.EventReceive, func(event *net.Event){
		log.Print("EventReceive : ", string(event.Data))
		msg, err := codec.CodecManager().Decode("json", event.Data)
		var userAuth *protocal.UserAuth = nil

		if err != nil {
			log.Fatal(err)
		} else {
			userAuth = msg.(*protocal.UserAuth)
			log.Print("userAuth.name : ", userAuth.Name)
		}
		event.Session.Send(event.Data)
	})

	peer.RegisterEvent(net.EventConnected, func(event *net.Event){
		log.Print("EventConnected !")
	})

	peer.RegisterEvent(net.EventClosed, func(event *net.Event){
		log.Print("EventClosed !")
	})

	peer.Start("http://127.0.0.1:8801/login")
	queue.StartLoop()
	queue.Wait()
	peer.Stop()
}

