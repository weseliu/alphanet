package websocket

import (
	"github.com/weseliu/alphanet/net"
	"github.com/gorilla/websocket"

	"net/http"
	"log"
	"net/url"
)

type wsAcceptor struct {
	*net.PeerBase
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (Self *wsAcceptor) Start(address string) net.Peer {
	if Self.IsRunning() {
		return Self
	}

	Self.SetRunning(true)
	url, err := url.Parse(address)
	if err != nil {
		return Self
	}

	if url.Path == "" {
		log.Fatalln("websocket: expect path in url to listen", address)
		return Self
	}

	Self.SetAddress(address)
	http.HandleFunc(url.Path, func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		session := NewSession(c, Self)
		Self.Add(session)
		session.OnClose = func() {
			Self.Remove(session)
		}
		session.run()
	})

	go func() {
		err = http.ListenAndServe(url.Host, nil)
		if err != nil {
			log.Fatalln(err)
		}
		Self.SetRunning(false)
	}()

	return Self
}

func (Self *wsAcceptor) Stop() {
	if !Self.IsRunning() {
		return
	}
}

func NewAcceptor(q net.EventQueue) net.Peer {
	self := &wsAcceptor{
		PeerBase : &net.PeerBase{
			EventQueue : q,
			SessionManager : net.NewSessionManager(),
		},
	}

	return self
}
