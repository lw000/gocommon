package tyutils

import (
	"time"
)

func TExecTime(info string, f func() (interface{}, error)) (data interface{}, tm time.Duration, err error) {
	start := time.Now()
	data, err = f()
	return data, time.Since(start), err
}
