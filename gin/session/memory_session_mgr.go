package session

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}
func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}
func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	id := uuid.NewV4()
	// 将id转化成string
	sessionid := id.String()
	session = NewMemorySession(sessionid)
	m.sessionMap[sessionid] = session
	return
}
func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	session, ok := m.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
	}
	return
}
