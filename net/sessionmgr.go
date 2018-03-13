package net

import (
	"sync"
	"sync/atomic"
)

type SessionManager interface {
	Add(Session)
	Remove(Session)
	SessionCount() int
	GetSession(int64) Session
	VisitSession(func(Session) bool)
	CloseAllSession()
}

type sessionManager struct {
	sessionMap map[int64]Session
	sessionIDAcc    int64
	sessionMapGuard sync.RWMutex
}

func (Self *sessionManager) Add(session Session) {
	Self.sessionMapGuard.Lock()
	defer Self.sessionMapGuard.Unlock()

	var id = atomic.AddInt64(&Self.sessionIDAcc, 1)
	session.SetID(id)
	Self.sessionMap[id] = session
}

func (Self *sessionManager) Remove(session Session) {
	Self.sessionMapGuard.Lock()
	delete(Self.sessionMap, session.ID())
	Self.sessionMapGuard.Unlock()
}

func (Self *sessionManager) GetSession(id int64) Session {
	Self.sessionMapGuard.RLock()
	defer Self.sessionMapGuard.RUnlock()

	v, ok := Self.sessionMap[id]
	if ok {
		return v
	}

	return nil
}

func (Self *sessionManager) VisitSession(callback func(Session) bool) {
	Self.sessionMapGuard.RLock()
	defer Self.sessionMapGuard.RUnlock()

	for _, ses := range Self.sessionMap {
		if !callback(ses) {
			break
		}
	}
}

func (Self *sessionManager) CloseAllSession() {
	Self.VisitSession(func(session Session) bool {
		session.Close()
		return true
	})
}

func (Self *sessionManager) SessionCount() int {
	Self.sessionMapGuard.Lock()
	defer Self.sessionMapGuard.Unlock()

	return len(Self.sessionMap)
}

func NewSessionManager() SessionManager {
	return &sessionManager{
		sessionMap: make(map[int64]Session),
	}
}
