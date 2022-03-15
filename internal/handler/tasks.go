package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/jfyne/live"
)

type (
	TasksInstance struct {
		Session string
		Tasks   []*domain.Task
		Error   error
	}
)

func (h *Handler) NewTasksInstance(ctx context.Context, s live.Socket) *TasksInstance {
	m, ok := s.Assigns().(*TasksInstance)
	if !ok {
		tasks, err := h.app.GetTasks(ctx, 10, 0)
		return &TasksInstance{
			Session: fmt.Sprint(s.Session()),
			Tasks:   tasks,
			Error:   err,
		}
	}

	return m
}

func (h *Handler) Tasks() live.Handler {
	t, err := template.ParseFiles(h.t+"layout.html", h.t+"tasks.html")
	if err != nil {
		log.Fatal(err)
	}

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		return h.NewTasksInstance(ctx, s), nil
	})

	return lvh
}
