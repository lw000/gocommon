package tycache

import (
	"reflect"
	"sync"
	"testing"

	"github.com/golang/groupcache/lru"
)

func TestNewLruCache(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *LruCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLruCache(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLruCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLruCache_Add(t *testing.T) {
	type fields struct {
		cache *lru.Cache
		m     sync.RWMutex
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LruCache{
				cache: tt.fields.cache,
				m:     tt.fields.m,
			}
			lru.Add(tt.args.key, tt.args.value)
		})
	}
}

func TestLruCache_Remove(t *testing.T) {
	type fields struct {
		cache *lru.Cache
		m     sync.RWMutex
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LruCache{
				cache: tt.fields.cache,
				m:     tt.fields.m,
			}
			lru.Remove(tt.args.key)
		})
	}
}

func TestLruCache_Clear(t *testing.T) {
	type fields struct {
		cache *lru.Cache
		m     sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LruCache{
				cache: tt.fields.cache,
				m:     tt.fields.m,
			}
			lru.Clear()
		})
	}
}

func TestLruCache_Get(t *testing.T) {
	type fields struct {
		cache *lru.Cache
		m     sync.RWMutex
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue interface{}
		wantOk    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LruCache{
				cache: tt.fields.cache,
				m:     tt.fields.m,
			}
			gotValue, gotOk := lru.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("LruCache.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("LruCache.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestLruCache_Len(t *testing.T) {
	type fields struct {
		cache *lru.Cache
		m     sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := &LruCache{
				cache: tt.fields.cache,
				m:     tt.fields.m,
			}
			if got := lru.Len(); got != tt.want {
				t.Errorf("LruCache.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
