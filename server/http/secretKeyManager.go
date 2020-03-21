package server

import (
	"math/rand"
	"time"
)

const (
	MemoryManager = "memory"
	RedisManager  = "redis"
)

type SecretKeyInfo struct {
	SecretKey string
	expiry    int64
}

type SecretKeyManager interface {
	ManagerInit()
	CreateSecretKeyByUser(userid string, expiry int64) string
	GetSecretKeyByUser(userid string) string
	DeleteSecretKeyByUser(userid string)
}

type SecretKeyManagerFactory struct {
}

func (f SecretKeyManagerFactory) GetSecretKeyManagerManager(typeName string) SecretKeyManager {
	switch typeName {
	case MemoryManager:
		return &MemorySecretKeyManager{}
	case RedisManager:
		return nil
	default:
		return nil
	}
}

func (info *SecretKeyInfo) valid() bool {
	return time.Now().Unix() <= info.expiry
}

func generateSecretKey() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~=+%^*/()[]{}/!@#$?|"
	bytes := []byte(str)
	randomKey := []byte{}
	length := 12
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		randomKey = append(randomKey, bytes[r.Intn(len(bytes))])
	}
	return string(randomKey)
}
