package net

type Session interface {
	SetID(id int64)
	ID() int64

	SetPeer(peer Peer)
	Peer() Peer

	SetTag(tag interface{})
	Tag() interface{}

	Send(interface{})
	Close()
}

type SessionBase struct {
	id int64
	peer Peer
	tag interface{}
}

func (Self *SessionBase) SetID(id int64){
	Self.id = id
}

func (Self *SessionBase) ID() int64{
	return Self.id
}

func (Self *SessionBase) SetPeer(peer Peer){
	Self.peer = peer
}

func (Self *SessionBase) Peer() Peer{
	return Self.peer
}

func (Self *SessionBase) SetTag(tag interface{}){
	Self.tag = tag
}

func (Self *SessionBase) Tag() interface{}{
	return Self.tag
}

func (Self *SessionBase) Send(data interface{}){
}

func (Self *SessionBase) Close(){
}


