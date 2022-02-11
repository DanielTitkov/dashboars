package tasks

import (
	"context"
	"math/rand"
	"time"

	"github.com/DanielTitkov/dashboars/internal/domain"
)

const (
	argMin = "min"
	argMax = "max"
)

type RandomTaskArgs struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func NewRandomTask(args domain.CreateTaskArgs) (*domain.Task, error) {
	min, err := requireFloat(args.Args, argMin)
	if err != nil {
		return nil, err
	}

	max, err := requireFloat(args.Args, argMax)
	if err != nil {
		return nil, err
	}

	taskArgs := RandomTaskArgs{
		Max: max,
		Min: min,
	}

	return &domain.Task{
		Type:        args.Type,
		Code:        args.Code,
		Title:       args.Title,
		Description: args.Description,
		Active:      args.Active,
		Display:     args.Display,
		Args:        &taskArgs,
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
	case "min":
		return a.Min, true
	case "max":
		return a.Max, true
	default:
		return 0, false
	}
}

func (a *RandomTaskArgs) GetInt(string) (int, bool) {
	return 0, false
}

func RandomTaskResolveFn(ctx context.Context, t *domain.Task, args domain.TaskArgs) (*domain.Item, error) {
	min, ok := args.GetFloat("min")
	if !ok {
		return nil, newArgError("min")
	}

	max, ok := args.GetFloat("max")
	if !ok {
		return nil, newArgError("max")
	}

	v := min + rand.Float64()*(max-min)
	return &domain.Item{
		Value:     v,
		Timestamp: time.Now(),
		Meta:      make(map[string]interface{}),
	}, nil
}
