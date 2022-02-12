package domain

import (
	"context"
	"time"
)

type (
	TaskArgs interface {
		Get(string) (interface{}, bool)
		GetString(string) (string, bool)
		GetFloat(string) (float64, bool)
		GetInt(string) (int, bool)
		ToMap() map[string]interface{}
	}

	Task struct {
		ID          int                                                                `json:"id"`
		Type        string                                                             `json:"type"`
		Code        string                                                             `json:"code"`
		Title       string                                                             `json:"title"`
		Description string                                                             `json:"description"`
		Active      bool                                                               `json:"active"`
		Display     bool                                                               `json:"display"`
		Schedule    string                                                             `json:"schedule"`
		Args        TaskArgs                                                           `json:"-"`
		ResolveFn   func(context.Context, *Task, *TaskInstance) (*TaskInstance, error) `json:"-"`
	}

	TaskInstance struct {
		ID        int
		TaskID    int
		StartTime time.Time
		EndTime   time.Time
		Success   bool
		Running   bool
		Error     error
		Attempt   int
		Items     []*Item
		Meta      map[string]interface{}
	}

	Item struct {
		ID             int
		TaskInstanceID int
		Value          float64
		Timestamp      time.Time
		Meta           map[string]interface{}
	}

	SystemSymmary struct {
		ID              int
		Users           int
		Tasks           int
		ActiveTasks     int
		Items           int
		AvgItemsPerTask float64
		CreateTime      time.Time
	}
)
