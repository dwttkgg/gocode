package session

import (
	"errors"
	"sync"
)

type MemmorySession struct {
	sessionId string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

// 构造函数
func NewMemorySession(id string) *MemmorySession {
	return &MemmorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
}
func (m *MemmorySession) Set(key string, value interface{}) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}
func (m *MemmorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return nil, err
	}
	return
}
func (m *MemmorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return
}
func (m *MemmorySession) Save() (err error) {
	return
}
