package cache

import (
	"context"
	"errors"
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"sync"
	"time"
)

// LocalCodeCache 本地缓存实现
type LocalCodeCache struct {
	cache      *lru.Cache
	mutex      sync.Mutex
	expiration time.Duration
	maps       sync.Map
}

func NewLocalCodeCache(c *lru.Cache, expiration time.Duration) *LocalCodeCache {
	return &LocalCodeCache{
		cache:      c,
		expiration: expiration,
	}
}

func (l *LocalCodeCache) Set(ctx context.Context, biz string, phone string, code string) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	key := l.key(biz, phone)

	l.maps.Store(key, code)

	if l.cache != nil {
		l.cache.Add(key, code)
	}

	return nil
}

func (l *LocalCodeCache) Verify(ctx context.Context, biz string, phone string, inputCode string) (bool, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	key := l.key(biz, phone)

	storedCode, ok := l.maps.Load(key)
	if !ok {
		return false, errors.New("Code not found")
	}

	if storedCode.(string) == inputCode {
		return true, nil
	}

	return false, nil
}

func (l *LocalCodeCache) key(biz string, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
