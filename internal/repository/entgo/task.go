package entgo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/DanielTitkov/dashboars/internal/domain/tasks"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
)

func (r *EntgoRepository) GetTaskCount(ctx context.Context) (int, error) {
	return r.client.Task.Query().Count(ctx)
}

func (r *EntgoRepository) GetActiveTaskCount(ctx context.Context) (int, error) {
	return r.client.Task.Query().Where(task.ActiveEQ(true)).Count(ctx)
}

func (r *EntgoRepository) CreateOrUpdateTask(ctx context.Context, t *domain.Task) (*domain.Task, error) {
	var tagIDs []int
	for _, tag := range t.Tags {
		if tag.ID == 0 {
			return nil, errors.New("tag must have valid ID")
		}
		tagIDs = append(tagIDs, tag.ID)
	}

	// query task by code
	tsk, err := r.client.Task.
		Query().
		Where(task.CodeEQ(t.Code)).
		WithCategory().
		WithTags().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		// create task
		tskQuery := r.client.Task.
			Create().
			SetCode(t.Code).
			SetType(task.Type(t.Type)).
			SetTitle(t.Title).
			SetDescription(t.Description).
			SetActive(t.Active).
			SetDisplay(t.Display).
			SetSchedule(t.Schedule).
			SetArgs(t.Args.ToMap())

		if t.Category != nil {
			tskQuery.SetCategoryID(t.Category.ID)
		}

		if len(tagIDs) > 0 {
			tskQuery.AddTagIDs(tagIDs...)
		}

		tsk, err = tskQuery.Save(ctx)
		if err != nil {
			return nil, err
		}
		return entToDomainTask(tsk)
	}

	// update task
	tskUpdateQuery := tsk.Update().
		SetTitle(t.Title).
		SetDescription(t.Description).
		SetActive(t.Active).
		SetDisplay(t.Display).
		SetSchedule(t.Schedule).
		SetArgs(t.Args.ToMap())

	if t.Category != nil {
		tskUpdateQuery.SetCategoryID(t.Category.ID)
	}

	if len(tagIDs) > 0 {
		tskUpdateQuery.RemoveTagIDs(tagIDs...)
		tskUpdateQuery.AddTagIDs(tagIDs...)
	}

	tsk, err = tskUpdateQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: maybe query task with edges to get updated tags and categories

	return entToDomainTask(tsk)
}

func entToDomainCreateTaskArgs(t *ent.Task) domain.CreateTaskArgs {
	jsonArgs, err := json.Marshal(t.Args)
	if err != nil {
		panic(err) // FIXME
	}

	return domain.CreateTaskArgs{
		ID:          t.ID,
		Code:        t.Code,
		Type:        t.Type.String(),
		Title:       t.Title,
		Display:     t.Display,
		Active:      t.Active,
		Description: t.Description,
		Schedule:    t.Schedule,
		Args:        jsonArgs,
	}
}

func entToDomainTask(t *ent.Task) (*domain.Task, error) {
	var category *domain.TaskCategory
	if t.Edges.Category != nil {
		category = entToDomainTaskCategory(t.Edges.Category)
	}

	var tags []*domain.TaskTag
	if t.Edges.Tags != nil {
		for _, tag := range t.Edges.Tags {
			tags = append(tags, entToDomainTaskTag(tag))
		}
	}

	return tasks.CreateTask(entToDomainCreateTaskArgs(t), category, tags)
}
