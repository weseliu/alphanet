package main

import (
	"github.com/weseliu/alphanet/net"
	"github.com/weseliu/alphanet/net/websocket"
)

func main() {
	net.NewAcceptorPeer("login", "http://127.0.0.1:8801/login", func(queue net.EventQueue) net.Peer {
		return websocket.NewAcceptor(queue)
	})
}

