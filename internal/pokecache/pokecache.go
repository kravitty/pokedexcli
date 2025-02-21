package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	interval time.Duration
	cache    map[string]cacheEntry
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	fmt.Println("NewCache")
	c := &Cache{
		interval: interval,
		cache:    make(map[string]cacheEntry),
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Println("Add")
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, exists bool) {
	//fmt.Println("Get")
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	fmt.Println(entry.val)
	return entry.val, true
}

func (c *Cache) reapLoop() {
	fmt.Println("reapLoop")
	ticker := time.Tick(c.interval)
	for range ticker {
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
