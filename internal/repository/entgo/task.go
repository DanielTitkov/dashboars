package entgo

import (
	"context"

	"github.com/DanielTitkov/dashboars/internal/domain/tasks"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
)

func (r *EntgoRepository) CreateOrUpdateTask(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	// query task by code
	tsk, err := r.client.Task.
		Query().
		Where(task.CodeEQ(t.Code)).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create task
		tsk, err = r.client.Task.
			Create().
			SetCode(t.Code).
			SetType(task.Type(t.Type)).
			SetTitle(t.Title).
			SetDescription(t.Description).
			SetActive(t.Active).
			SetDisplay(t.Display).
			SetSchedule(t.Schedule).
			SetArgs(t.Args.ToMap()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return entToDomainTask(tsk)
	}

	// update task
	tsk, err = tsk.Update().
		SetTitle(t.Title).
		SetDescription(t.Description).
		SetActive(t.Active).
		SetDisplay(t.Display).
		SetSchedule(t.Schedule).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainTask(tsk)
}

func entToDomainCreateTaskArgs(t *ent.Task) domain.CreateTaskArgs {
	return domain.CreateTaskArgs{
		ID:          t.ID,
		Code:        t.Code,
		Type:        t.Type.String(),
		Title:       t.Title,
		Display:     t.Display,
		Active:      t.Active,
		Description: t.Description,
		Schedule:    t.Schedule,
		Args:        t.Args,
	}
}

func entToDomainTask(t *ent.Task) (*domain.Task, error) {
	return tasks.CreateTask(entToDomainCreateTaskArgs(t))
}
