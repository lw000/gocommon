package tyblw

import (
	"github.com/panjf2000/ants"
)

// 协成池

type CoroutinePool struct {
	pool *ants.Pool
}

func New() *CoroutinePool {
	return &CoroutinePool{}
}

func (co *CoroutinePool) Start(size int) error {
	var err error
	co.pool, err = ants.NewPool(size)
	if err != nil {
		return err
	}
	return nil
}

func (co *CoroutinePool) Tune(size uint) {
	co.pool.Tune(size)
}

func (co *CoroutinePool) Close() {
	co.pool.Release()
}

func (co *CoroutinePool) Submit(task func()) {
	if err := co.pool.Submit(task); err != nil {
	}
}
