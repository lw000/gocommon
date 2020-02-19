package tycache

import (
	"github.com/bluele/gcache"
	"time"
)

// MemCache ...
type MemCache interface {
	Set(key interface{}, value interface{}) error
	SetWithExpire(interface{}, interface{}, time.Duration) error
	Get(key interface{}) (interface{}, error)
	Exist(key interface{}) bool
	Remove(key interface{}) bool
	RemoveAll()
	GetALL() map[interface{}]interface{}
	Keys() []interface{}
	Len() int
}

type memCache struct {
	store gcache.Cache
}

// NewLRU ...
func NewLRU(size int) MemCache {
	return &memCache{store: gcache.New(size).LRU().Build()}
}

// NewLFU ...
func NewLFU(size int) MemCache {
	return &memCache{store: gcache.New(size).LFU().Build()}
}

// NewARC ...
func NewARC(size int) MemCache {
	return &memCache{store: gcache.New(size).ARC().Build()}
}

// Set ...
func (m *memCache) Set(key interface{}, value interface{}) error {
	err := m.store.Set(key, value)
	if err != nil {
		return err
	}

	return nil
}

// Set ...
func (m *memCache) SetWithExpire(key interface{}, value interface{}, t time.Duration) error {
	err := m.store.SetWithExpire(key, value, t)
	if err != nil {
		return err
	}
	return nil
}

// Get ...
func (m *memCache) Get(key interface{}) (interface{}, error) {
	value, err := m.store.Get(key)
	if err != nil {
		return nil, err
	}
	return value, err
}

// Exist ...
func (m *memCache) Exist(key interface{}) bool {
	return m.store.Has(key)
}

// Remove...
func (m *memCache) Remove(key interface{}) bool {
	return m.store.Remove(key)
}

// RemoveAll ...
func (m *memCache) RemoveAll() {
	keys := m.store.Keys(true)
	for _, k := range keys {
		// t := reflect.TypeOf(k)
		// switch t.Kind() {
		// case reflect.String:
		// case reflect.Int32:
		// }
		m.store.Remove(k)
	}
}

// GetALL ...
func (m *memCache) GetALL() map[interface{}]interface{} {
	return m.store.GetALL(true)
}

// Keys ...
func (m *memCache) Keys() []interface{} {
	return m.store.Keys(true)
}

// Len ...
func (m *memCache) Len() int {
	return m.store.Len(true)
}
