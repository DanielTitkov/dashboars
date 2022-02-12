package tasks

import (
	"context"
	"errors"

	"github.com/DanielTitkov/dashboars/internal/domain"
)

const ()

type FailingTaskArgs struct {
}

func NewFailingTask(args domain.CreateTaskArgs) (*domain.Task, error) {
	return &domain.Task{
		Type:        args.Type,
		Code:        args.Code,
		Title:       args.Title,
		Description: args.Description,
		Active:      args.Active,
		Display:     args.Display,
		Schedule:    args.Schedule,
		Args:        &FailingTaskArgs{},
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

func (a *FailingTaskArgs) GetInt(string) (int, bool) {
	return 0, false
}

func FailingTaskResolveFn(ctx context.Context, t *domain.Task, args domain.TaskArgs) (*domain.Item, error) {
	return nil, errors.New("test error from failing task")
}
