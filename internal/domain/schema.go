package domain

import (
	"context"
	"encoding/json"
	"time"
)

type (
	TaskArgs interface {
		Get(string) (interface{}, bool)
		GetString(string) (string, bool)
		GetFloat(string) (float64, bool)
		GetInt(string) (int, bool)
		ToMap() map[string]interface{}
		ToJSON() json.RawMessage
	}

	Task struct {
		ID          int
		Type        string
		Code        string
		Title       string
		Description string
		Active      bool
		Display     bool
		Schedule    string
		Category    *TaskCategory
		Metrics     []*Metric
		Tags        []*TaskTag
		Args        TaskArgs                                                           `json:"-"`
		ResolveFn   func(context.Context, *Task, *TaskInstance) (*TaskInstance, error) `json:"-"`
	}

	TaskCategory struct {
		ID          int
		Title       string
		Description string
		Display     bool
		Meta        map[string]interface{}
	}

	TaskTag struct {
		ID          int
		Title       string
		Description string
		Display     bool
		Meta        map[string]interface{}
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

	Metric struct {
		ID          int
		TaskID      int
		Title       string                 `json:"title"`
		Description string                 `json:"description"`
		Meta        map[string]interface{} `json:"meta"`
	}

	Dimension struct {
		ID       int
		MetricID int
		Title    string // e.g. city
		Value    string // e.g. Moscow
		Meta     map[string]interface{}
	}

	Item struct {
		ID             int
		TaskInstanceID int
		Metric         *Metric
		Value          float64
		Dimensions     []*Dimension
		Timestamp      time.Time
		Meta           map[string]interface{}
	}

	SystemSymmary struct {
		ID                      int
		Tasks                   int
		ActiveTasks             int
		RunningTasks            int
		CompletedTasks          int
		FailedTasks             int
		Metrics                 int
		CollectedItems          int
		AvgItemsPerTask         float64
		AvgItemsPerTaskInstance float64
		AvgItemsPerMetric       float64
		CreateTime              time.Time
	}
)
