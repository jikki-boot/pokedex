package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	contents map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(t time.Duration) *Cache {
	cache := &Cache{
		contents: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}

	go cache.reapLoop(t)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	c.contents[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()

	entry, exists := c.contents[key]
	c.mu.Unlock()

	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()

		currentTime := time.Now()
		ageLimit := currentTime.Add(-interval)

		for k, v := range c.contents {
			if ageLimit.Compare(v.createdAt) < 0 {
				continue
			}

			delete(c.contents, k)
		}

		c.mu.Unlock()
	}
}
