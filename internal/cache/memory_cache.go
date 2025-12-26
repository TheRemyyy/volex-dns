package cache

import (
	"log"
	"sync"
	"time"

	"github.com/miekg/dns"
)

type cacheEntry struct {
	msg       *dns.Msg
	expiresAt time.Time
}

type MemoryCache struct {
	sync.RWMutex
	m map[string]*cacheEntry
}

func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{
		m: make(map[string]*cacheEntry),
	}
	go c.cleanupLoop()
	return c
}

func (c *MemoryCache) Get(key string) (*dns.Msg, bool) {
	c.RLock()
	entry, found := c.m[key]
	c.RUnlock()

	if found && time.Now().Before(entry.expiresAt) {
		return entry.msg, true
	}
	return nil, found
}

func (c *MemoryCache) Set(key string, msg *dns.Msg, ttl uint32) {
	c.Lock()
	c.m[key] = &cacheEntry{
		msg:       msg,
		expiresAt: time.Now().Add(time.Duration(ttl) * time.Second),
	}
	c.Unlock()
}

func (c *MemoryCache) cleanupLoop() {
	for {
		time.Sleep(10 * time.Minute)
		c.Lock()
		for k, v := range c.m {
			if time.Now().After(v.expiresAt) {
				delete(c.m, k)
			}
		}
		log.Printf("Cache cleanup done. Items remaining: %d", len(c.m))
		c.Unlock()
	}
}
