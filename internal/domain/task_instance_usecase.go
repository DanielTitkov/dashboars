package domain

import (
	"time"
)

func InitTaskInstance(taskID int, attempt int) *TaskInstance {
	return &TaskInstance{
		TaskID:    taskID,
		Attempt:   attempt,
		StartTime: time.Now(),
		Running:   true,
	}
}

func (ti *TaskInstance) WithSuccess(items []*Item) *TaskInstance {
	ti.Error = nil
	ti.Success = true
	ti.Items = items
	ti.Running = false
	return ti
}

func (ti *TaskInstance) WithError(err error) *TaskInstance {
	ti.Error = err
	ti.Success = false
	ti.Running = false
	return ti
}
