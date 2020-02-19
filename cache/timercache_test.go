package tycache

import (
	"reflect"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func TestNewTimerCache(t *testing.T) {
	type args struct {
		defaultExpiration time.Duration
		cleanupInterval   time.Duration
	}
	tests := []struct {
		name string
		args args
		want *TimerCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTimerCache(tt.args.defaultExpiration, tt.args.cleanupInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTimerCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimerCache_Add(t *testing.T) {
	type fields struct {
		cache *cache.Cache
	}
	type args struct {
		key   string
		value interface{}
		d     time.Duration
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
			tc := &TimerCache{
				cache: tt.fields.cache,
			}
			if err := tc.Add(tt.args.key, tt.args.value, tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("TimerCache.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimerCache_Remove(t *testing.T) {
	type fields struct {
		cache *cache.Cache
	}
	type args struct {
		key string
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
			tc := &TimerCache{
				cache: tt.fields.cache,
			}
			tc.Remove(tt.args.key)
		})
	}
}

func TestTimerCache_Get(t *testing.T) {
	type fields struct {
		cache *cache.Cache
	}
	type args struct {
		key string
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
			tc := &TimerCache{
				cache: tt.fields.cache,
			}
			gotValue, gotOk := tc.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("TimerCache.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("TimerCache.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
