package tasks

import (
	"context"
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

func randomTaskArgsFromMap(args map[string]interface{}) (*RandomTaskArgs, error) {
	min, err := requireFloat(args, argMin)
	if err != nil {
		return nil, err
	}

	max, err := requireFloat(args, argMax)
	if err != nil {
		return nil, err
	}

	return &RandomTaskArgs{
		Max: max,
		Min: min,
	}, nil
}

func NewRandomTask(args domain.CreateTaskArgs) (*domain.Task, error) {
	taskArgs, err := randomTaskArgsFromMap(args.Args)
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
		Args:        taskArgs,
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

func RandomTaskResolveFn(ctx context.Context, t *domain.Task, ti *domain.TaskInstance) (*domain.TaskInstance, error) {
	min, ok := t.Args.GetFloat(argMin)
	if !ok {
		return ti.WithError(newArgError(argMin)), newArgError(argMin)
	}

	max, ok := t.Args.GetFloat(argMax)
	if !ok {
		return ti.WithError(newArgError(argMax)), newArgError(argMax)
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
		Meta:      make(map[string]interface{}),
		Metric:    metricRaw,
	}

	itemSquare := &domain.Item{
		Value:     v * v,
		Timestamp: time.Now(),
		Meta:      make(map[string]interface{}),
		Metric:    metricSquare,
	}

	return ti.WithSuccess([]*domain.Item{itemRaw, itemSquare}), nil
}
