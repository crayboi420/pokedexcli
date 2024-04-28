package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	mapEntry map[string]CacheEntry
	mtx      *sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) Cache {
	ch := Cache{
		mapEntry: map[string]CacheEntry{},
		interval: interval,
		mtx:      &sync.RWMutex{},
	}
	go ch.reapLoop()
	return ch
}

func (ch Cache) Add(key string, val []byte) {
	entry := CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}

	ch.mtx.Lock()
	ch.mapEntry[key] = entry
	defer ch.mtx.Unlock()
}

func (ch Cache) Get(key string) (val []byte, is bool) {
	ch.mtx.RLock()
	ans, is := ch.mapEntry[key]
	ch.mtx.RUnlock()
	if is {
		val = ans.val
	}
	return
}

func (ch Cache) reapLoop() {
	ticker := time.NewTicker(ch.interval)
	for range ticker.C {
		ch.mtx.Lock()
		nowTime := time.Now()
		for key, val := range ch.mapEntry {
			if nowTime.Sub(val.createdAt) > ch.interval {
				// fmt.Println("Deleting key ",key)
				delete(ch.mapEntry, key)
			}
		}
		ch.mtx.Unlock()
	}
}
