package pokecache

import (
	"sync"
	"time"
)


type Cache struct {
	store 		map[string]cacheEntry
	mu 			*sync.Mutex
}


type cacheEntry struct {
	createdAt	time.Time
	val			[]byte	
}


func NewCache(interval time.Duration) Cache {

	cache := Cache{
		store:  make(map[string]cacheEntry),
		mu:		&sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: 	   val,
	}
}
	
	

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.store[key]
	
	if !ok {
		return nil, false
	}
	return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
	
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, val := range c.store {
		if val.createdAt.Before(now.Add( -last)) {
			delete(c.store, key)
		}
	}
}