package cache

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

type Cache struct {
	store  map[string][]byte
	mu     sync.RWMutex
	logger *zap.Logger
}

func Init(logger *zap.Logger) *Cache {
	return &Cache{
		store:  make(map[string][]byte),
		logger: logger,
	}
}

func (c *Cache) Set(key string, val []byte) error {
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
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	start := time.Now()
	c.mu.Lock()
	var hit bool
	defer func() {
		c.mu.Unlock()
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
