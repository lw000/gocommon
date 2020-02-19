package tycache

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type LeveldbCache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Exist(key string) bool
	Clear(key string) error
	Close() error
}

// LeveldbCache 缓存数据库
type leveldbCache struct {
	store *leveldb.DB
}

type Compare struct {
}

func (c *Compare) Compare(a, b []byte) int {
	return 0
}

func (c *Compare) Name() string {
	return "111"
}

func (c *Compare) Separator(dst, a, b []byte) []byte {
	return nil
}

func (c *Compare) Successor(dst, b []byte) []byte {
	return nil
}

// NewLDB 创建新的DB实例
func NewLDB(path string) (LeveldbCache, error) {
	var err error
	c := &leveldbCache{}
	// opt := opt2.Options{}
	// opt.BlockCacher = opt2.LRUCacher
	// opt.Comparer = &Compare{}
	c.store, err = leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Set ...
func (lc *leveldbCache) Set(key string, value []byte) error {
	err := lc.store.Put([]byte(key), value, nil)
	if err != nil {
	}
	return nil
}

// Get ...
func (lc *leveldbCache) Get(key string) ([]byte, error) {
	v, err := lc.store.Get([]byte(key), nil)
	if err != nil {

	}
	return v, nil
}

// Exist ...
func (lc *leveldbCache) Exist(key string) bool {
	ret, err := lc.store.Has([]byte(key), nil)
	if err != nil {

	}
	return ret
}

// Close ...
func (lc *leveldbCache) Close() error {
	err := lc.store.Close()
	return err
}

// Clear ...
func (lc *leveldbCache) Clear(key string) error {
	err := lc.store.Delete([]byte(key), nil)
	return err
}
