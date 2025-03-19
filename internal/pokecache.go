package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data   map[string]cacheEntry
	mu     sync.Mutex
	expiry time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func newCache(expiry time.Duration) *Cache {
	cache := &Cache{
		data:   make(map[string]cacheEntry),
		expiry: expiry,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	if time.Since(entry.createdAt) > c.expiry {
		delete(c.data, key)
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.expiry)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > c.expiry {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}
