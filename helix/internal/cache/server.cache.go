package cache

import (
	"sync"

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
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = val
	if c.logger != nil {
		c.logger.Info("Cache set", zap.String("key", key))
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.store[key]
	if c.logger != nil {
		if ok {
			c.logger.Info("Cache get hit", zap.String("key", key))
		} else {
			c.logger.Info("Cache get miss", zap.String("key", key))
		}
	}
	return val, ok
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
	if c.logger != nil {
		c.logger.Info("Cache delete", zap.String("key", key))
	}
	return nil
}
