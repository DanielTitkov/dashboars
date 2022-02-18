package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/util"
	"github.com/hashicorp/go-multierror"
)

const (
	argItemsNumber = "itemsNumber"
	argFailOnEvery = "failOnEvery"
)

type FailingTaskArgs struct {
	ItemsNumber int `json:"itemsNumber"`
	FailOnEvery int `json:"failOnEvery"`
}

func NewFailingTask(args domain.CreateTaskArgs, cat *domain.TaskCategory, tags []*domain.TaskTag) (*domain.Task, error) {
	var taskArgs FailingTaskArgs
	err := json.Unmarshal(args.Args, &taskArgs)
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:          args.ID,
		Type:        args.Type,
		Code:        args.Code,
		Title:       args.Title,
		Description: args.Description,
		Active:      args.Active,
		Display:     args.Display,
		Schedule:    args.Schedule,
		Category:    cat,
		Tags:        tags,
		Args:        &taskArgs,
		ResolveFn:   FailingTaskResolveFn,
	}, nil
}

func (a *FailingTaskArgs) Get(string) (interface{}, bool) {
	var ret interface{}
	return ret, false
}

func (a *FailingTaskArgs) GetString(string) (string, bool) {
	return "", false
}

func (a *FailingTaskArgs) GetFloat(s string) (float64, bool) {
	return 0, false
}

func (a *FailingTaskArgs) GetInt(s string) (int, bool) {
	switch s {
	case argItemsNumber:
		return a.ItemsNumber, true
	case argFailOnEvery:
		return a.FailOnEvery, true
	default:
		return 0, false
	}
}

func (a *FailingTaskArgs) ToMap() map[string]interface{} {
	return util.ToMap(a)
}

func (a *FailingTaskArgs) ToJSON() json.RawMessage {
	return json.RawMessage{}
}

func FailingTaskResolveFn(ctx context.Context, t *domain.Task, ti *domain.TaskInstance) (*domain.TaskInstance, error) {
	nItems, ok := t.Args.GetInt(argItemsNumber)
	if !ok {
		err := newArgError(argItemsNumber)
		return ti.WithError(err, nil), err
	}

	failOnEvery, ok := t.Args.GetInt(argFailOnEvery)
	if !ok {
		err := newArgError(argFailOnEvery)
		return ti.WithError(err, nil), err
	}

	metric := &domain.Metric{
		Title:       "failing_task_metric",
		Description: "Dummy metric for failing task",
		TaskID:      t.ID,
	}

	var items []*domain.Item
	var err error
	for i := 0; i < nItems; i++ {
		if i%failOnEvery == 0 {
			err = multierror.Append(err, fmt.Errorf("test error from failing task, failed item %d", i+1))
			continue
		}

		meta := make(map[string]interface{})
		meta[argItemsNumber] = nItems
		meta[argFailOnEvery] = failOnEvery

		items = append(items, &domain.Item{
			Value:      777,
			Timestamp:  time.Now(),
			Dimensions: []*domain.Dimension{},
			Meta:       meta,
			Metric:     metric,
		})
	}

	if err != nil {
		return ti.WithError(err, items), err
	}

	return ti.WithSuccess(items), nil
}
