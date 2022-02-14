package app

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/dashboars/internal/domain"
)

func (a *App) GetSystemSummary(ctx context.Context) (*domain.SystemSymmary, error) {
	if a.systemSummary == nil {
		a.log.Debug("system summary requested but not found, gathering...", "")
		err := a.updateSystemSummary(ctx)
		if err != nil {
			return nil, err
		}
	}

	return a.systemSummary, nil
}

func (a *App) updateSystemSummary(ctx context.Context) error {
	a.log.Debug("updating system summary", "")

	metricCount, err := a.repo.GetMetricCount(ctx)
	if err != nil {
		return err
	}

	taskCount, err := a.repo.GetTaskCount(ctx)
	if err != nil {
		return err
	}

	activeTaskCount, err := a.repo.GetActiveTaskCount(ctx)
	if err != nil {
		return err
	}

	taskInstanceCount, err := a.repo.GetTaskInstanceCount(ctx)
	if err != nil {
		return err
	}

	failedTaskCount, err := a.repo.GetFailedTaskInstanceCount(ctx)
	if err != nil {
		return err
	}

	runningTaskCount, err := a.repo.GetRunningTaskInstanceCount(ctx)
	if err != nil {
		return err
	}

	successfulTaskCount, err := a.repo.GetSuccesfulTaskInstanceCount(ctx)
	if err != nil {
		return err
	}

	itemCount, err := a.repo.GetItemCount(ctx)
	if err != nil {
		return err
	}

	a.systemSummary = &domain.SystemSymmary{
		Tasks:                   taskCount,
		ActiveTasks:             activeTaskCount,
		FailedTasks:             failedTaskCount,
		RunningTasks:            runningTaskCount,
		CompletedTasks:          successfulTaskCount,
		Metrics:                 metricCount,
		CollectedItems:          itemCount,
		AvgItemsPerTask:         float64(itemCount) / float64(taskCount),
		AvgItemsPerTaskInstance: float64(itemCount) / float64(taskInstanceCount),
	}

	a.log.Debug("system summary updated", fmt.Sprintf("%+v", a.systemSummary))
	return nil
}
