package main

import "github.com/weseliu/alphanet/net"

func main() {
	net.NewAcceptorPeer("login", "http://127.0.0.1:8801/login", net.PeerTypeWebSocket)
}

