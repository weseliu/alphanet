package main

import (
	"github.com/weseliu/alphanet/cmd/connect/handler"

	"github.com/weseliu/alphanet/codec"
	"github.com/weseliu/alphanet/db"
	"github.com/weseliu/alphanet/net"
	"github.com/weseliu/alphanet/net/websocket"
	_ "github.com/weseliu/alphanet/codec/pb"
	"github.com/weseliu/alphanet/util"
)

func main() {
	codec.RegisterPbMessageMeta()
	config := util.Configs("./conf/connect.json")

	db.Instance().Open(config.String("dsn"))
	var logHandler handler.LoginHandler
	logHandler.Start()
	queue := net.NewEventQueue()
	peer := net.PeerManager().NewPeer("connect", queue, func(queue net.EventQueue) net.Peer {
		return websocket.NewAcceptor(queue)
	})

	var eventHandler handler.EventHandler
	eventHandler.Start(peer)

	peer.Start(config.String("url"))
	queue.StartLoop()
	queue.Wait()
	peer.Stop()
	db.Instance().Close()
}
