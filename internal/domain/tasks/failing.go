package tasks

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/DanielTitkov/dashboars/internal/domain"
)

type FailingTaskArgs struct {
}

func NewFailingTask(args domain.CreateTaskArgs, cat *domain.TaskCategory, tags []*domain.TaskTag) (*domain.Task, error) {
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

func (a *FailingTaskArgs) ToMap() map[string]interface{} {
	return make(map[string]interface{})
}

func (a *FailingTaskArgs) ToJSON() json.RawMessage {
	return json.RawMessage{}
}

func FailingTaskResolveFn(ctx context.Context, t *domain.Task, ti *domain.TaskInstance) (*domain.TaskInstance, error) {
	err := errors.New("test error from failing task")
	return ti.WithError(err, nil), err
}
