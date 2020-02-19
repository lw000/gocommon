package tyschedule

import (
	"log"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	sche := NewSchedule()
	sche.HandleFunc(func(data interface{}) {
		log.Println(time.Now().Format("2006-01-02 15:04:05.000000000"), data)
	})

	var (
		er     error
		taskId string
	)
	er = sche.Start()
	if er != nil {
		log.Println(er)
		return
	}

	for i, v := range []string{"1111111111", "2222222222", "3333333333", "4444444444"} {
		taskId, er = sche.AddTask(i+1, v)
		if er != nil {
			log.Println(er)
			return
		}
		log.Println("taskId:", taskId)
	}

	time.Sleep(time.Second * time.Duration(5))
	sche.Stop()
}

func TestMonitor(t *testing.T) {
	monitor := NewMonitor()
	monitor.HandleFunc(func(currentRetryCount int, maxRetryCount int) (ok bool) {
		log.Printf("currentRetryCount:%d maxRetryCount:%d", currentRetryCount, maxRetryCount)
		return false
	})

	er := monitor.Start(1, 5)
	if er != nil {
		log.Println(er)
		return
	}

	time.Sleep(time.Second * time.Duration(10))

	monitor.Stop()
}
