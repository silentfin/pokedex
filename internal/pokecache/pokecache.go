package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entry    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entry:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop() // coolest thing ever
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cached, ok := c.entry[key]
	if !ok {
		return nil, false
	} else {
		return cached.val, true
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		interval := time.Now().Add(-c.interval)
		for k, e := range c.entry {
			if e.createdAt.Before(interval) {
				delete(c.entry, k)
			}
		}

		c.mu.Unlock()
	}
}
