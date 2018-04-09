package main

import (
	"github.com/weseliu/alphanet/cmd/connect/handler"

	"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/db"
	"github.com/weseliu/alphanet/net"
	"github.com/weseliu/alphanet/net/websocket"
	_ "github.com/weseliu/alphanet/codec/pb"
)

func main() {
	codec.RegisterPbMessageMeta()

	db.Instance().Open("root:@tcp(localhost:3306)/alphanet?charset=utf8")
	var logHandler handler.LoginHandler
	logHandler.Start()
	queue := net.NewEventQueue()
	peer := net.PeerManager().NewPeer("connect", queue, func(queue net.EventQueue) net.Peer {
		return websocket.NewAcceptor(queue)
	})

	var eventHandler handler.EventHandler
	eventHandler.Start(peer)

	peer.Start("http://127.0.0.1:8801/login")
	queue.StartLoop()
	queue.Wait()
	peer.Stop()
	db.Instance().Close()
}
