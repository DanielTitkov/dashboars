package tasks

import (
	"errors"

	"github.com/DanielTitkov/dashboars/internal/domain"
)

func getTaskConstructorFn(args domain.CreateTaskArgs) (func(domain.CreateTaskArgs, *domain.TaskCategory, []*domain.TaskTag) (*domain.Task, error), error) {
	switch args.Type {
	case domain.TaskTypeRandom:
		return NewRandomTask, nil
	case domain.TaskTypeFailing:
		return NewFailingTask, nil
	default:
		return nil, errors.New("unknown task type: " + args.Type)
	}
}

func CreateTask(args domain.CreateTaskArgs, cat *domain.TaskCategory, tags []*domain.TaskTag) (*domain.Task, error) {
	constructorFn, err := getTaskConstructorFn(args)
	if err != nil {
		return nil, err
	}

	return constructorFn(args, cat, tags)
}
