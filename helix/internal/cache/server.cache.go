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

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.store[key]
	return val, ok
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
	return nil
}
