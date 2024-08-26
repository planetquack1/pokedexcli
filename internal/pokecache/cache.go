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
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func newCache(interval time.Duration) *Cache {
	return &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
}
