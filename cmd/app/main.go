package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/DanielTitkov/dashboars/internal/app"
	"github.com/DanielTitkov/dashboars/internal/configs"
	"github.com/DanielTitkov/dashboars/internal/handler"
	"github.com/DanielTitkov/dashboars/logger"
	"github.com/jfyne/live"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	log.Println("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	log.Println("loaded config")

	logger := logger.NewLogger(cfg.Env)
	defer logger.Sync()
	logger.Info("starting service", "")

	a, err := app.New(cfg, logger)
	if err != nil {
		logger.Fatal("failed to init app", err)
	}

	err = a.LoadTasksPresets()
	if err != nil {
		logger.Fatal("failed load tasks", err)
	}

	h := handler.NewHandler(
		a,
		logger,
		"templates/",
	)

	// Run the server.
	http.Handle("/", live.NewHttpHandler(live.NewCookieStore("session-name", []byte("weak-secret")), h.Home()))
	// live scripts
	http.Handle("/live.js", live.Javascript{})
	http.Handle("/auto.js.map", live.JavascriptMap{})
	// favicon
	http.HandleFunc("/favicon.ico", faviconHandler)
	// serve
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/favicon.ico")
}
