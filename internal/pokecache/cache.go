package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	list     map[string]cacheEntry
	mux      sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) Add(key string, val []byte) {
	entry := cacheEntry{createdAt: time.Now(), val: val}
	cache.mux.Lock()
	cache.list[key] = entry
	cache.mux.Unlock()

}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()
	entry, ok := cache.list[key]
	cache.mux.Unlock()
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache) reapLoop() {
	ticket := time.NewTicker(cache.interval)
	defer ticket.Stop()

	for range ticket.C {
		cache.mux.Lock()
		for key, entry := range cache.list {
			if time.Since(entry.createdAt) > cache.interval {
				delete(cache.list, key)

			}

		}
		cache.mux.Unlock()
	}

}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		list:     map[string]cacheEntry{},
		mux:      sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop()

	return cache
}
