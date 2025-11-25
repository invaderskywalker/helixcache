package cache

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Cache struct {
	store   map[string][]byte // value
	expires map[string]int64  // expiry timestamp (unix millis)
	hits    map[string]uint64 // access count (for LFU)
	mu      sync.RWMutex
	logger  *zap.Logger

	// sizeLimit int               // optional max entries or memory limit
}

func Init(logger *zap.Logger) *Cache {
	cache := Cache{
		store:   make(map[string][]byte),
		expires: make(map[string]int64),
		hits:    make(map[string]uint64),
		logger:  logger,
	}
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// check now if time has passed ttl
				sample := 20
				fmt.Println("length of store ", len(cache.store), sample)
				i := 0
				for key, exp := range cache.expires {
					if i >= sample {
						break
					}
					if exp > 0 && time.Now().UnixMilli() > exp {
						delete(cache.store, key)
						delete(cache.expires, key)
						delete(cache.hits, key)
						if cache.logger != nil {
							cache.logger.Info("Cache expired", zap.String("key", key))
						}
					}
					i++
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return &cache
}

func (c *Cache) Set(key string, val []byte, ttlMillis int64) error {
	start := time.Now()
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		if c.logger != nil {
			duration := time.Since(start)
			c.logger.Info("Cache set", zap.String("key", key), zap.Duration("duration", duration))
		}
	}()
	c.store[key] = val
	if ttlMillis > 0 {
		c.expires[key] = time.Now().UnixMilli() + ttlMillis
	} else {
		c.expires[key] = 0
	}

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	start := time.Now()
	c.mu.RLock()
	var hit bool
	defer func() {
		c.mu.RUnlock()
		if c.logger != nil {
			duration := time.Since(start)
			if hit {
				c.logger.Info("Cache get hit", zap.String("key", key), zap.Duration("duration", duration))
			} else {
				c.logger.Info("Cache get miss", zap.String("key", key), zap.Duration("duration", duration))
			}
		}
	}()
	val, ok := c.store[key]
	if !ok {
		hit = false
		return nil, false
	}
	exp, hasExp := c.expires[key]
	if hasExp && exp > 0 && time.Now().UnixMilli() > exp {
		delete(c.store, key)
		delete(c.expires, key)
		delete(c.hits, key)
		hit = false
		return nil, false
	}
	c.hits[key]++
	hit = ok
	return val, ok
}

func (c *Cache) Delete(key string) error {
	start := time.Now()
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
		if c.logger != nil {
			duration := time.Since(start)
			c.logger.Info("Cache delete", zap.String("key", key), zap.Duration("duration", duration))
		}
	}()
	delete(c.store, key)
	return nil
}
