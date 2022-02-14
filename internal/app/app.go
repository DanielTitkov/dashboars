package app

import (
	"context"
	"sync"
	"time"

	"github.com/go-co-op/gocron"

	"github.com/DanielTitkov/dashboars/internal/configs"
	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/logger"
)

type (
	App struct {
		cfg             configs.Config
		log             *logger.Logger
		scheduler       *gocron.Scheduler
		repo            Repository
		tasks           []*domain.Task
		systemSummary   *domain.SystemSymmary
		taskResultMutex sync.Mutex
	}
	Repository interface {
		CreateOrUpdateTask(context.Context, *domain.Task) (*domain.Task, error)
		CreateTaskInstance(context.Context, *domain.TaskInstance) (*domain.TaskInstance, error)
		UpdateTaskInstance(context.Context, *domain.TaskInstance) (*domain.TaskInstance, error)

		// for system summary
		GetItemCount(ctx context.Context) (int, error)
		GetMetricCount(ctx context.Context) (int, error)
		GetTaskCount(ctx context.Context) (int, error)
		GetActiveTaskCount(ctx context.Context) (int, error)
		GetTaskInstanceCount(ctx context.Context) (int, error)
		GetFailedTaskInstanceCount(ctx context.Context) (int, error)
		GetRunningTaskInstanceCount(ctx context.Context) (int, error)
		GetSuccesfulTaskInstanceCount(ctx context.Context) (int, error)
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) (*App, error) {
	s := gocron.NewScheduler(time.UTC)
	s.SetMaxConcurrentJobs(cfg.Task.MaxConcurrency, gocron.WaitMode)
	// s.WaitForScheduleAll()

	app := App{
		cfg:       cfg,
		log:       logger,
		scheduler: s,
		repo:      repo,
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
	go app.UpdateSystemSummaryJob() // TODO: move to jobs?

	s.StartAsync()

	return &app, nil
}
