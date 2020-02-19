package tycache

import (
	"reflect"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestNewLDBCache(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    LeveldbCache
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLDBCache(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLDBCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLDBCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leveldbCache_Set(t *testing.T) {
	type fields struct {
		store *leveldb.DB
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
			lc := &leveldbCache{
				store: tt.fields.store,
			}
			if err := lc.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("leveldbCache.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_leveldbCache_Get(t *testing.T) {
	type fields struct {
		store *leveldb.DB
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
			lc := &leveldbCache{
				store: tt.fields.store,
			}
			got, err := lc.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("leveldbCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leveldbCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leveldbCache_Exist(t *testing.T) {
	type fields struct {
		store *leveldb.DB
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
			lc := &leveldbCache{
				store: tt.fields.store,
			}
			if got := lc.Exist(tt.args.key); got != tt.want {
				t.Errorf("leveldbCache.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leveldbCache_Remove(t *testing.T) {
	type fields struct {
		store *leveldb.DB
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
			lc := &leveldbCache{
				store: tt.fields.store,
			}
			if got := lc.Remove(tt.args.key); got != tt.want {
				t.Errorf("leveldbCache.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leveldbCache_Select(t *testing.T) {
	type fields struct {
		store *leveldb.DB
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
			lc := &leveldbCache{
				store: tt.fields.store,
			}
			if got := lc.Select(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leveldbCache.Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leveldbCache_Keys(t *testing.T) {
	type fields struct {
		store *leveldb.DB
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
			lc := &leveldbCache{
				store: tt.fields.store,
			}
			if got := lc.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leveldbCache.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leveldbCache_Close(t *testing.T) {
	type fields struct {
		store *leveldb.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &leveldbCache{
				store: tt.fields.store,
			}
			if err := lc.Close(); (err != nil) != tt.wantErr {
				t.Errorf("leveldbCache.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
