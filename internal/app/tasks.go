package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/domain/tasks"
)

func (a *App) makeTaskJob(task *domain.Task, attempt int) func() {
	return func() {
		// create task instance in db
		ti := domain.InitTaskInstance(task.ID, attempt)
		ti, err := a.repo.CreateTaskInstance(context.Background(), ti)
		if err != nil {
			a.log.Error("failed save new task instance to db", err)
		}

		// run resolve func
		a.log.Debug(fmt.Sprintf("resolving task, attempt %d", attempt), task.Code)
		ti, err = task.ResolveFn(context.TODO(), task, ti)
		if err != nil {
			a.log.Error("task failed", err)
			if attempt < a.cfg.Task.DefaultRetryNumber {
				_, err := a.scheduler.Every(a.cfg.Task.DefaultRetryDelay).LimitRunsTo(1).WaitForSchedule().Do(a.makeTaskJob(task, attempt+1))
				if err != nil {
					a.log.Error("failed to schedule retry", err)
				}
			} else {
				a.log.Error("task reached retry limit", err)
			}
		}

		// save results to db
		a.taskResultMutex.Lock()
		// this is needed to ensure transactions won't fail on parallel creation of
		// same dimensions
		// TODO: probably better move to db level
		ti, err = a.repo.UpdateTaskInstance(context.TODO(), ti)
		a.taskResultMutex.Unlock()
		if err != nil {
			a.log.Error("failed to save updated task instance to db", err)
		}

		a.log.Debug("task result", fmt.Sprintf("%+v, err: %s", ti, err))
	}
}

func (a *App) scheduleTasks() error {
	for _, t := range a.tasks {
		if !t.Active {
			continue
		}

		var err error
		taskJob := a.makeTaskJob(t, 0)
		if strings.Contains(t.Schedule, "*") || strings.Contains(t.Schedule, " ") {
			// if this is cron string
			_, err = a.scheduler.Cron(t.Schedule).Tag(t.Code, t.Type).Do(taskJob)
		} else {
			// if this is something like 1h etc.
			_, err = a.scheduler.Every(t.Schedule).Tag(t.Code, t.Type).Do(taskJob)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) loadTasksPresets() error {
	a.log.Info("loading tasks", fmt.Sprint(a.cfg.Data.Presets.TaskPresetsPaths))
	for _, path := range a.cfg.Data.Presets.TaskPresetsPaths {
		a.log.Debug("reading from file", path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		var presets []domain.CreateTaskArgs
		err = json.Unmarshal(data, &presets)
		if err != nil {
			return err
		}

		for _, preset := range presets {
			ctx := context.Background()

			category, err := a.repo.GetOrCreateTaskCategory(ctx, &domain.TaskCategory{Title: preset.Category})
			if err != nil {
				a.log.Fatal("failed to get or create category", err)
			}

			var tags []*domain.TaskTag
			for _, t := range preset.Tags {
				tag, err := a.repo.GetOrCreateTaskTag(ctx, &domain.TaskTag{Title: t})
				if err != nil {
					a.log.Fatal("failed to get or create tag", err)
				}
				tags = append(tags, tag)
			}

			task, err := tasks.CreateTask(preset, category, tags)
			if err != nil {
				a.log.Fatal("failed to create task", err)
			}

			// saving to db, must have id after that
			task, err = a.repo.CreateOrUpdateTask(ctx, task)
			if err != nil {
				a.log.Fatal("failed to save task to db", err)
			}

			a.tasks = append(a.tasks, task)
			a.log.Info("loaded task", fmt.Sprintf("%+v", task))
		}
	}

	return nil
}
