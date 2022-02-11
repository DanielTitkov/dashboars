package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/DanielTitkov/dashboars/internal/configs"
	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/domain/tasks"
	"github.com/DanielTitkov/dashboars/logger"
)

type (
	App struct {
		cfg configs.Config
		log *logger.Logger
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	// repo Repository,
) (*App, error) {
	app := App{
		cfg: cfg,
		log: logger,
		// repo:   repo,
	}

	// init app jobs, caches and preload data (if any)

	return &app, nil
}

func (a *App) LoadTasksPresets() error {
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

		a.log.Info("loaded task", fmt.Sprintf("%+v", task))
	}

	return nil
}
