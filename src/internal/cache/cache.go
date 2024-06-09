package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   sync.Mutex
	interval int
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval int) *Cache{
	return &Cache{
		make(map[string]cacheEntry),
		sync.Mutex{},
		interval,
	}
}

func (c *Cache) Add(url string, val []byte) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache[url] = cacheEntry{createdAt: time.Now(), val: val}

	// Start a goroutine to clear expired cache entries periodically
	go func() {
		clearCache(60, c.cache, &c.mux)
	}()

	return nil
}

// clearCache deletes entries that have been in the cache for longer than cacheDuration seconds
func clearCache(cacheDuration int, m map[string]cacheEntry, mux *sync.Mutex) {
	for {
		time.Sleep(time.Duration(cacheDuration) * time.Second)
		mux.Lock()
		for key, entry := range m {
			if time.Since(entry.createdAt) > time.Duration(cacheDuration)*time.Second {
				delete(m, key)
			}
		}
		mux.Unlock()
	}
}

func (c *Cache) Get(url string) ([]byte , error){
	value , ok := c.cache[url]
	if !ok {
		return nil, errors.New("key not found")
	}	
	return value.val , nil
}
