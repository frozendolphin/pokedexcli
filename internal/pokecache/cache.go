package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	MapsCache map[string]cacheEntry
	mu *sync.RWMutex
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func (c Cache) Add(key string, val []byte){
	new_entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.MapsCache[key] = new_entry
}

func (c Cache) Get(key string) ([]byte, bool){
	c.mu.RLock()
	entry, ok := c.MapsCache[key]
	c.mu.RUnlock()
	if !ok {
		return nil, ok
	}
	return entry.val, ok
}

func (c Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.MapsCache {
		elapsed := time.Since(val.createdAt)
		if elapsed > interval {
			delete(c.MapsCache, key)
		}	
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		MapsCache: make(map[string]cacheEntry),
		mu: &sync.RWMutex{},
	}
	ticker := time.NewTicker(interval)
	
	go func() {
		for range ticker.C {
			c.reapLoop(interval)
		}
	}()
	
	return c
}

