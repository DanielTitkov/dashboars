package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-co-op/gocron"

	"github.com/DanielTitkov/dashboars/internal/configs"
	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/domain/tasks"
	"github.com/DanielTitkov/dashboars/logger"
)

type (
	App struct {
		cfg       configs.Config
		log       *logger.Logger
		scheduler *gocron.Scheduler
		tasks     []*domain.Task
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	// repo Repository,
) (*App, error) {
	s := gocron.NewScheduler(time.UTC)
	s.SetMaxConcurrentJobs(cfg.Task.MaxConcurrency, gocron.WaitMode)
	s.WaitForScheduleAll()

	app := App{
		cfg:       cfg,
		log:       logger,
		scheduler: s,
		// repo:   repo,
	}

	err := app.loadTasksPresets()
	if err != nil {
		return nil, err
	}

	err = app.scheduleTasks()
	if err != nil {
		return nil, err
	}

	// init app jobs, caches and preload data (if any)

	s.StartAsync()

	return &app, nil
}

func (a *App) makeTaskJob(task *domain.Task, try int) func() {
	return func() {
		// create task instance in db

		// run resolve func
		a.log.Info(fmt.Sprintf("resolving task, attempt %d", try), task.Code)
		result, err := task.ResolveFn(context.TODO(), task, task.Args)
		if err != nil {
			a.log.Error("task failed", err)
			if try < a.cfg.Task.DefaultRetryNumber {
				a.log.Info("scheduling retry", "")
				_, err := a.scheduler.Every(a.cfg.Task.DefaultRetryDelay).LimitRunsTo(1).WaitForSchedule().Do(a.makeTaskJob(task, try+1))
				if err != nil {
					a.log.Error("failed to schedule retry", err)
				}
			} else {
				a.log.Error("task reached retry limit", err)
			}
		}
		a.log.Info("task result", fmt.Sprint(result))
		// save results to db

		// update task instance
	}
}

func (a *App) scheduleTasks() error {
	for _, t := range a.tasks {
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
		// ind, err := a.repo.GetIndicatorByCode(indicator.Code)
		// if err == nil && ind.ID != 0 {
		// 	a.log.Debug("indicator already exists", ind.JSONString())
		// 	continue
		// }
		task, err := tasks.CreateTask(preset)
		if err != nil {
			a.log.Fatal("failed to create task", err)
		}

		// add saving to db
		a.tasks = append(a.tasks, task)
		a.log.Info("loaded task", fmt.Sprintf("%+v", task))
	}

	return nil
}
