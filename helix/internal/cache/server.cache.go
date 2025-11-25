package cache

import "sync"

type Cache struct {
	store map[string][]byte
	mu    sync.RWMutex
}

func Init() *Cache {
	return &Cache{
		store: make(map[string][]byte),
	}
}

func (c *Cache) Set(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = val
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// @TODO: how to check if key exists in c.store
	return c.store[key], nil
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = nil
	return nil
}
