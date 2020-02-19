package tyschedule

import (
	"errors"
	"github.com/lw000/gocommon/utils"
	"time"

	"github.com/ouqiang/timewheel"
)

type Task struct {
	Data interface{}
	Fn   func(data interface{})
}

type Schedule struct {
	wheel *timewheel.TimeWheel
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func (s *Schedule) AddTask(second int64, task *Task) (taskid string) {
	taskid = tyutils.UUID()
	if taskid == "" {
		return ""
	}
	s.wheel.AddTimer(time.Second*time.Duration(second), taskid, task)
	return taskid
}

func (s *Schedule) RemoveTask(taskId interface{}) {
	s.wheel.RemoveTimer(taskId)
}

func (s *Schedule) Start() error {
	s.wheel = timewheel.New(time.Second*time.Duration(1), 3600, func(taskData interface{}) {
		task := taskData.(*Task)
		task.Fn(task.Data)
	})

	if s.wheel == nil {
		return errors.New("start schedule failed")
	}

	s.wheel.Start()

	return nil
}

func (s *Schedule) Stop() {
	if s == nil {
		return
	}
	s.wheel.Stop()
}
