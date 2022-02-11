package tasks

import (
	"errors"

	"github.com/DanielTitkov/dashboars/internal/domain"
)

func getTaskConstructorFn(args domain.CreateTaskArgs) (func(domain.CreateTaskArgs) (*domain.Task, error), error) {
	switch args.Type {
	case domain.TaskTypeRandom:
		return NewRandomTask, nil
	default:
		return nil, errors.New("unknown task type: " + args.Type)
	}
}

func CreateTask(args domain.CreateTaskArgs) (*domain.Task, error) {
	constructorFn, err := getTaskConstructorFn(args)
	if err != nil {
		return nil, err
	}

	return constructorFn(args)
}
