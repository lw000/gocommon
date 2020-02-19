package tyIdWorker

import (
	"fmt"
	"time"

	// "github.com/gitstliu/go-id-worker"
	// "github.com/bwmarrin/snowflake"
	"github.com/zheng-ji/goSnowFlake"
)

// sonyflake 是 Sony 公司的一个开源项目，基本思路和 snowflake 差不多，不过位分配上稍有不同：

// +-----------------------------------------------------------------------------+
// | 1 Bit Unused | 39 Bit Timestamp |  8 Bit Sequence ID  |   16 Bit Machine ID |
// +-----------------------------------------------------------------------------+
// 这里的时间只用了 39 个 bit，但时间的单位变成了 10ms，所以理论上比 41 位表示的时间还要久(174 years)。

type IdWorker struct {
	worker *goSnowFlake.IdWorker
}

func (iw *IdWorker) Start(workerId int64) error {
	var err error
	iw.worker, err = goSnowFlake.NewIdWorker(workerId)
	if err != nil {
		return err
	}
	return nil
}

func (iw *IdWorker) Id() int64 {
	newId, err := iw.worker.NextId()
	if err != nil {
		return -1
	}
	return newId
}

func (iw *IdWorker) String() string {
	t := time.Now()
	return fmt.Sprintf("%d%d%d%d%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), iw.Id())
}
