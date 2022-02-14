package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/DanielTitkov/dashboars/cmd/app/prepare"
	"github.com/DanielTitkov/dashboars/internal/app"
	"github.com/DanielTitkov/dashboars/internal/configs"
	"github.com/DanielTitkov/dashboars/internal/handler"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent"
	"github.com/DanielTitkov/dashboars/logger"
	"github.com/jfyne/live"

	_ "github.com/lib/pq"
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

	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	err = prepare.Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	repo := entgo.NewEntgoRepository(db, logger)

	a, err := app.New(cfg, logger, repo)
	if err != nil {
		logger.Fatal("failed to init app", err)
	}

	h := handler.NewHandler(
		a,
		logger,
		"templates/",
	)

	// Run the server.
	http.Handle("/", live.NewHttpHandler(live.NewCookieStore("session-name", []byte("weak-secret")), h.Home()))
	http.Handle("/summary", live.NewHttpHandler(live.NewCookieStore("session-name", []byte("weak-secret")), h.SystemSummary()))
	// live scripts
	http.Handle("/live.js", live.Javascript{})
	http.Handle("/auto.js.map", live.JavascriptMap{})
	// favicon
	http.HandleFunc("/favicon.ico", faviconHandler)
	// serve
	log.Fatal(http.ListenAndServe(cfg.Server.GetAddress(), nil))
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/favicon.ico")
}
