package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/weseliu/alphanet/net"
)

type wsSession struct {
	*net.SessionBase
	conn *websocket.Conn
	OnClose func()
}

func (Self *wsSession) Send(data interface{}){
	go func() {
		Self.conn.WriteMessage(websocket.BinaryMessage, data.([]byte))
	}()
}

func (Self *wsSession) Close(){
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

	switch t {
	case websocket.TextMessage:
		return 0, net.EventResultCodecError
	case websocket.CloseMessage:
		return 0, net.EventResultRequestClose
	case websocket.BinaryMessage:
		data = raw
	}

	return data, net.EventResultOK
}

func (Self *wsSession) receiveThread() {
	for {
		data, result := Self.readPacket()
		if result == net.EventResultOK {
			ev := net.NewEvent(net.EventReceive, Self)
			ev.Data = data.([]byte)
			ev.Session = Self
			Self.Peer().OnEvent(ev)
		}

		if result == net.EventResultRequestClose || result == net.EventResultSocketError{
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
		SessionBase : &net.SessionBase{
		},
		conn: c,
	}
	session.SetPeer(p)
	return session
}

