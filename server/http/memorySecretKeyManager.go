package server

import (
	"sync"
	"time"
)

type MemorySecretKeyManager struct {
	UserTokenExpiries map[string]*SecretKeyInfo
	lock              sync.Mutex
}

func (manager *MemorySecretKeyManager) ManagerInit() {
	manager.UserTokenExpiries = make(map[string]*SecretKeyInfo)
	go manager.GC()
}

func (manager *MemorySecretKeyManager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	for key, value := range manager.UserTokenExpiries {
		if !value.valid() {
			delete(manager.UserTokenExpiries, key)
		}
	}
	time.AfterFunc(time.Minute*1, func() { manager.GC() })
}

func (manager *MemorySecretKeyManager) CreateSecretKeyByUser(userid string, expiry int64) string {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	key := generateSecretKey()
	manager.UserTokenExpiries[userid] = &SecretKeyInfo{key, expiry}
	return key
}

func (manager *MemorySecretKeyManager) GetSecretKeyByUser(userid string) string {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	info := manager.UserTokenExpiries[userid]
	if info == nil {
		return ""
	}

	if !info.valid() {
		delete(manager.UserTokenExpiries, userid)
	}

	return info.SecretKey
}

func (manager *MemorySecretKeyManager) DeleteSecretKeyByUser(userid string) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	delete(manager.UserTokenExpiries, userid)
}
