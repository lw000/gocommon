package tycache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type TimerCache struct {
	cache *cache.Cache
}

func NewTimerCache(defaultExpiration, cleanupInterval time.Duration) *TimerCache {
	return &TimerCache{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (tc *TimerCache) Add(key string, value interface{}, d time.Duration) error {
	if err := tc.cache.Add(key, value, d); err != nil {
		return err
	}
	return nil
}

func (tc *TimerCache) Remove(key string) {
	tc.cache.Delete(key)
}

func (tc *TimerCache) Get(key string) (value interface{}, ok bool) {

	return tc.cache.Get(key)
}
