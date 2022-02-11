package handler

import (
	"github.com/DanielTitkov/dashboars/internal/app"
	"github.com/DanielTitkov/dashboars/logger"
)

type (
	Handler struct {
		app *app.App
		log *logger.Logger
		t   string // template path
	}
)

func NewHandler(
	app *app.App,
	logger *logger.Logger,
	t string,
) *Handler {
	return &Handler{
		app: app,
		log: logger,
		t:   t,
	}
}
