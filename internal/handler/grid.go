package handler

import (
	"context"
	"html/template"
	"strings"

	"github.com/bradfitz/iter"

	"github.com/jfyne/live"
)

const ()

var funcMap = template.FuncMap{
	"N":     iter.N,
	"Title": strings.Title,
	"sub": func(x, y int) int {
		return x - y
	},
}

type (
	GridModel struct {
	}
	Corr struct {
		Left  string
		Right string
		Value float64
	}
)

func AssignGridModel(s live.Socket) *GridModel {
	m, ok := s.Assigns().(*GridModel)
	if !ok {
		return &GridModel{}
	}

	return m
}

func (h *Handler) Grid() live.Handler {
	t := template.Must(template.New("layout.html").Funcs(funcMap).ParseFiles(
		h.t+"layout.html",
		h.t+"grid.html",
		h.t+"grid_terms.html",
		h.t+"grid_triads.html",
		h.t+"grid_linking.html",
		h.t+"grid_result.html",
		h.t+"alerts.html",
	))

	lvh := live.NewHandler(live.WithTemplateRenderer(t))

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		return AssignGridModel(s), nil
	})

	return lvh
}
