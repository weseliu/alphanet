package net

func NewAcceptorPeer(name string, address string, creator func(queue EventQueue) Peer) Peer {
	queue := NewEventQueue()
	peer := PeerManager().NewPeer(name, queue, creator)

	defer func() {
		peer.Start(address)
		queue.StartLoop()
		queue.Wait()
		peer.Stop()
	}()
	return peer
}
