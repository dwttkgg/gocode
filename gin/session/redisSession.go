package session

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/garyburd/redigo/redis"
)

const (
	SessionFlagNone = iota
	SessionFlagModify
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	// 设置session 可以先放在内存中 批量导入reids，可以提升性能
	sessionMap map[string]interface{}
	// 读写锁
	rwlock sync.RWMutex
	flag   int
}

func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	return &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}, 16),
		flag:       SessionFlagNone,
	}
}
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 设置值
	r.sessionMap[key] = value
	// 设置完成后修改对应标志位
	r.flag = SessionFlagModify
	return
}
func (r *RedisSession) Save(key string, value interface{}) (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 如果值有变化则需要存入redis,否则不用
	if r.flag != SessionFlagModify {
		return
	}
	// 对内存中的sessionMgr序列号
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return
	}
	r.flag = SessionFlagNone
	return
}
func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	result, ok := r.sessionMap[key]
	if !ok {
		errors.New("key not exists")
	}

	return
}
