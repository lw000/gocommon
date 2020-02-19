package tycache

import (
	"reflect"
	"testing"

	"github.com/bluele/gcache"
)

func TestNewMemCache(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want MemCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemCache(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Set(t *testing.T) {
	type fields struct {
		store gcache.Cache
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				store: tt.fields.store,
			}
			if err := m.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("memCache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_memCache_Get(t *testing.T) {
	type fields struct {
		store gcache.Cache
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				store: tt.fields.store,
			}
			got, err := m.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("memCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Exist(t *testing.T) {
	type fields struct {
		store gcache.Cache
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				store: tt.fields.store,
			}
			if got := m.Exist(tt.args.key); got != tt.want {
				t.Errorf("memCache.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Remove(t *testing.T) {
	type fields struct {
		store gcache.Cache
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				store: tt.fields.store,
			}
			if got := m.Remove(tt.args.key); got != tt.want {
				t.Errorf("memCache.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Select(t *testing.T) {
	type fields struct {
		store gcache.Cache
	}
	tests := []struct {
		name   string
		fields fields
		want   map[interface{}]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				store: tt.fields.store,
			}
			if got := m.Select(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memCache.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memCache_Keys(t *testing.T) {
	type fields struct {
		store gcache.Cache
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memCache{
				store: tt.fields.store,
			}
			if got := m.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memCache.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}
