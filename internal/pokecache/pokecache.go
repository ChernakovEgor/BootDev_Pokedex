package pokecache

import (
	// "log"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	entries := make(map[string]CacheEntry)
	cache := &Cache{entries, sync.Mutex{}}
	cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Lock()
	defer c.Unlock()

	c.entries[key] = CacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()

	if val, ok := c.entries[key]; ok {
		return val.val, ok
	} else {
		return nil, ok
	}
}

func (c *Cache) cleanCache(interval time.Duration) {
	c.Lock()
	defer c.Unlock()

	for k, v := range c.entries {
		if v.createdAt.Before(time.Now().Add(-interval)) {
			delete(c.entries, k)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	// remove any entries older than interval
	// use time.Ticker

	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			// log.Printf("tick at %v", t)
			c.cleanCache(interval)
		}
	}()
}
