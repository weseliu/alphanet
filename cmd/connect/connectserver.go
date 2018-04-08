package main

import (
	"github.com/weseliu/alphanet/cmd/connect/handler"
	_ "github.com/weseliu/alphanet/codec/json"
	_ "github.com/weseliu/alphanet/codec/pb"
	"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/db"
	"github.com/weseliu/alphanet/net"
	"github.com/weseliu/alphanet/net/websocket"
)

func main() {
	codec.RegisterPbMessageMeta()

	db.Instance().Open("root:@tcp(localhost:3306)/alphanet?charset=utf8")

	var logHandler handler.LoginHandler
	logHandler.Start()
	queue := net.NewEventQueue()
	peer := net.PeerManager().NewPeer("login", queue, func(queue net.EventQueue) net.Peer {
		return websocket.NewAcceptor(queue)
	})

	var eventHandler handler.EventHandler
	eventHandler.Start(peer)
	//
	//peer.RegisterEventHandler(net.EventReceive, func(event *net.Event) {
	//	log.Print("EventReceive : ", string(event.Data))
	//	msg, err := codec.CodecManager().Decode("pb", event.Data, &connect.CSPACK{})
	//	if err == nil {
	//		core.MsgCenter().DispatchMessage(event.Session, msg)
	//	} else {
	//		log.Fatal(err)
	//	}
	//})
	//
	//peer.RegisterEventHandler(net.EventConnected, func(event *net.Event) {
	//	log.Print("EventConnected !")
	//})
	//
	//peer.RegisterEventHandler(net.EventClosed, func(event *net.Event) {
	//	log.Print("EventClosed !")
	//})

	peer.Start("http://127.0.0.1:8801/login")
	queue.StartLoop()
	queue.Wait()
	peer.Stop()
	db.Instance().Close()
}