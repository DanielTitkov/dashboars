package job

import (
	"github.com/DanielTitkov/dashboars/internal/app"
	"github.com/DanielTitkov/dashboars/internal/configs"
	"github.com/DanielTitkov/dashboars/logger"
)

type Job struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func New(
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Job {
	return &Job{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
}
