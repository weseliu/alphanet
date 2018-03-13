package net

import (
	"github.com/weseliu/alphanet/net/websocket"
)

type PeerType int32

const (
	PeerTypeWebSocket PeerType = iota
	PeerTypeSocket
)

func NewAcceptorPeer(name string, address string, peerType PeerType) Peer {
	var queue = NewEventQueue()
	var peer Peer = nil

	switch peerType {
	case PeerTypeWebSocket:
		{
			peer = PeerManager().NewPeer(name, func() Peer {
				return websocket.NewAcceptor(queue)
			})
		}
	case PeerTypeSocket:
		{

		}
	}

	defer func() {
		peer.Start(address)
		queue.StartLoop()
		queue.Wait()
		peer.Stop()
	}()
	return peer
}
