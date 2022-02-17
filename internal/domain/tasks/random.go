package tasks

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/util"
)

const (
	argMin = "min"
	argMax = "max"
)

type RandomTaskArgs struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func NewRandomTask(args domain.CreateTaskArgs, cat *domain.TaskCategory, tags []*domain.TaskTag) (*domain.Task, error) {
	var taskArgs RandomTaskArgs
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
		Args:        &taskArgs,
		Category:    cat,
		Tags:        tags,
		ResolveFn:   RandomTaskResolveFn,
	}, nil
}

func (a *RandomTaskArgs) Get(string) (interface{}, bool) {
	var ret interface{}
	return ret, false
}

func (a *RandomTaskArgs) GetString(string) (string, bool) {
	return "", false
}

func (a *RandomTaskArgs) GetFloat(s string) (float64, bool) {
	switch s {
	case argMin:
		return a.Min, true
	case argMax:
		return a.Max, true
	default:
		return 0, false
	}
}

func (a *RandomTaskArgs) GetInt(string) (int, bool) {
	return 0, false
}

func (a *RandomTaskArgs) ToMap() map[string]interface{} {
	return util.ToMap(a)
}

func (a *RandomTaskArgs) ToJSON() json.RawMessage {
	ret, err := json.Marshal(a)
	if err != nil {
		panic(err) // FIXME
	}
	return ret
}

func RandomTaskResolveFn(ctx context.Context, t *domain.Task, ti *domain.TaskInstance) (*domain.TaskInstance, error) {
	min, ok := t.Args.GetFloat(argMin)
	if !ok {
		err := newArgError(argMin)
		return ti.WithError(err, nil), err
	}

	max, ok := t.Args.GetFloat(argMax)
	if !ok {
		err := newArgError(argMax)
		return ti.WithError(err, nil), err
	}

	metricRaw := &domain.Metric{
		Title:       "Raw",
		Description: "Random number as is",
		TaskID:      t.ID,
	}

	metricSquare := &domain.Metric{
		Title:       "Square",
		Description: "Random number squared",
		TaskID:      t.ID,
	}

	v := min + rand.Float64()*(max-min)

	itemRaw := &domain.Item{
		Value:     v,
		Timestamp: time.Now(),
		Dimensions: []*domain.Dimension{
			{
				Title: "city",
				Value: "moscow",
			},
			{
				Title: "source",
				Value: "cian",
			},
		},
		Meta:   make(map[string]interface{}),
		Metric: metricRaw,
	}

	itemSquare := &domain.Item{
		Value:     v * v,
		Timestamp: time.Now(),
		Meta:      make(map[string]interface{}),
		Dimensions: []*domain.Dimension{
			{
				Title: "city",
				Value: "spb",
			},
			{
				Title: "source",
				Value: "cian",
			},
		},
		Metric: metricSquare,
	}

	itemSquare2 := &domain.Item{
		Value:     v * v,
		Timestamp: time.Now(),
		Meta:      make(map[string]interface{}),
		Dimensions: []*domain.Dimension{
			{
				Title: "city",
				Value: "kazan",
			},
			{
				Title: "source",
				Value: "yandex",
			},
		},
		Metric: metricSquare,
	}

	return ti.WithSuccess([]*domain.Item{itemRaw, itemSquare, itemSquare2}), nil
}
