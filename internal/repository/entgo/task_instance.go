package entgo

import (
	"context"
	"errors"
	"fmt"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/taskinstance"
)

func (r *EntgoRepository) GetTaskInstanceCount(ctx context.Context) (int, error) {
	return r.client.TaskInstance.Query().Count(ctx)
}

func (r *EntgoRepository) GetFailedTaskInstanceCount(ctx context.Context) (int, error) {
	return r.client.TaskInstance.Query().Where(taskinstance.And(
		taskinstance.RunningEQ(false),
		taskinstance.SuccessEQ(false),
	)).Count(ctx)
}

func (r *EntgoRepository) GetRunningTaskInstanceCount(ctx context.Context) (int, error) {
	return r.client.TaskInstance.Query().Where(taskinstance.And(
		taskinstance.RunningEQ(true),
	)).Count(ctx)
}

func (r *EntgoRepository) GetSuccesfulTaskInstanceCount(ctx context.Context) (int, error) {
	return r.client.TaskInstance.Query().Where(taskinstance.And(
		taskinstance.SuccessEQ(true),
	)).Count(ctx)
}

func (r *EntgoRepository) CreateTaskInstance(ctx context.Context, ti *domain.TaskInstance) (*domain.TaskInstance, error) {
	task, err := r.client.Task.Query().Where(task.IDEQ(ti.TaskID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	taskInstance, err := r.client.TaskInstance.
		Create().
		SetTask(task).
		SetStartTime(ti.StartTime).
		SetEndTime(ti.EndTime).
		SetSuccess(ti.Success).
		SetRunning(ti.Running).
		SetAttempt(ti.Attempt).
		SetMeta(ti.Meta).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainTaskInstance(taskInstance, ti.TaskID), nil
}

func (r *EntgoRepository) UpdateTaskInstance(ctx context.Context, ti *domain.TaskInstance) (*domain.TaskInstance, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	if len(ti.Items) > 0 {
		// TODO: maybe move to separate method
		bulk := make([]*ent.ItemCreate, len(ti.Items))
		for i, itm := range ti.Items {
			// TODO: this is a request or two for item. Maybe there's a way to make this more efficient.
			metric, err := getOrCreateMetricInTransaction(ctx, tx, itm.Metric)
			if err != nil {
				// supposedly tx is already rollbacked here
				return nil, err
			}

			bulk[i] = tx.Item.Create().
				SetTaskInstanceID(ti.ID).
				SetMetric(metric).
				SetValue(itm.Value).
				SetTimestamp(itm.Timestamp).
				SetMeta(itm.Meta)
		}
		_, err = tx.Item.CreateBulk(bulk...).Save(ctx)
		if err != nil {
			return nil, rollback(tx, err)
		}
	}

	taskInstance, err := tx.TaskInstance.UpdateOneID(ti.ID).
		SetStartTime(ti.StartTime).
		SetEndTime(ti.EndTime).
		SetSuccess(ti.Success).
		SetRunning(ti.Running).
		SetMeta(ti.Meta).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return entToDomainTaskInstance(taskInstance, ti.TaskID), tx.Commit()
}

func entToDomainTaskInstance(ti *ent.TaskInstance, taskIDs ...int) *domain.TaskInstance {
	var taskID int
	if len(taskIDs) == 1 {
		taskID = taskIDs[0]
	}

	if ti.Edges.Task != nil {
		taskID = ti.Edges.Task.ID
	}

	var errStr error
	if ti.Error != nil {
		errStr = errors.New(*ti.Error)
	}

	return &domain.TaskInstance{
		ID:        ti.ID,
		TaskID:    taskID,
		StartTime: ti.StartTime,
		EndTime:   ti.EndTime,
		Error:     errStr,
		Success:   *ti.Success,
		Running:   ti.Running,
		Meta:      ti.Meta,
		Items:     nil, // TODO add edge
	}
}
