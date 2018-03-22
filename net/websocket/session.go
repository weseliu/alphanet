package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/weseliu/alphanet/net"
)

type wsSession struct {
	*net.SessionBase
	conn    *websocket.Conn
	OnClose func()
}

func (Self *wsSession) Send(data interface{}) {
	go func() {
		Self.conn.WriteMessage(websocket.BinaryMessage, data.([]byte))
	}()
}

func (Self *wsSession) Close() {
	if Self.OnClose != nil {
		Self.OnClose()
	}
	Self.conn.Close()
}

func (Self *wsSession) readPacket() (data interface{}, result net.EventResult) {
	t, raw, err := Self.conn.ReadMessage()
	if err != nil {
		return 0, net.EventResultSocketError
	}

	if t == websocket.TextMessage || t == websocket.BinaryMessage {
		return raw, net.EventResultOK
	}

	return 0, net.EventResultRequestClose
}

func (Self *wsSession) receiveThread() {
	for {
		data, result := Self.readPacket()

		var event *net.Event = nil
		if result == net.EventResultOK {
			event = net.NewEvent(net.EventReceive, Self)
			event.Data = data.([]byte)
		} else if result == net.EventResultRequestClose || result == net.EventResultSocketError {
			event = net.NewEvent(net.EventClosed, Self)
		}
		Self.Peer().OnEvent(event)

		if result == net.EventResultRequestClose || result == net.EventResultSocketError {
			Self.Close()
			break
		}
	}
}

func (Self *wsSession) run() {
	go Self.receiveThread()
}

func NewSession(c *websocket.Conn, p net.Peer) *wsSession {
	session := &wsSession{
		SessionBase: &net.SessionBase{},
		conn:        c,
	}
	session.SetPeer(p)
	return session
}
