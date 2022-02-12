package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

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
		ti, err = a.repo.UpdateTaskInstance(context.TODO(), ti)
		if err != nil {
			a.log.Error("failed to save updated task instance to db", err)
		}

		a.log.Debug("task result", fmt.Sprintf("%+v", ti))
	}
}

func (a *App) scheduleTasks() error {
	for _, t := range a.tasks {
		if !t.Active {
			continue
		}

		_, err := a.scheduler.Every(t.Schedule).Tag(t.Code, t.Type).Do(a.makeTaskJob(t, 0))
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) loadTasksPresets() error {
	data, err := ioutil.ReadFile(a.cfg.Data.Presets.TasksPresetsPath)
	if err != nil {
		return err
	}

	var presets []domain.CreateTaskArgs
	err = json.Unmarshal(data, &presets)
	if err != nil {
		return err
	}

	for _, preset := range presets {
		task, err := tasks.CreateTask(preset)
		if err != nil {
			a.log.Fatal("failed to create task", err)
		}

		// saving to db, must have id after that
		task, err = a.repo.CreateOrUpdateTask(context.Background(), task)
		if err != nil {
			a.log.Fatal("failed to save task to db", err)
		}

		a.tasks = append(a.tasks, task)
		a.log.Info("loaded task", fmt.Sprintf("%+v", task))
	}

	return nil
}
