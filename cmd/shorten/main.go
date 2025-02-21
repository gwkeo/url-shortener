package main

import (
	"github.com/gwkeo/url-shortener/internal/config"
	"github.com/gwkeo/url-shortener/internal/http-server/server"
	"github.com/gwkeo/url-shortener/internal/repo/sqlite"
	"log/slog"
	"os"
)

const (
	envLocal = "LOCAL"
	envProd  = "PROD"
	envDebug = "DEBUG"
)

func main() {
	cfg := config.MustLoad()
	log := setUpLogger(cfg.Env)
	log.Info("starting shortener")

	repo, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("error opening sqlite repo", "error", err)
	}
	router := server.New(repo, log, cfg)
	if err = router.Start(); err != nil {
		log.Error("error starting shortener", "error", err)
	}
}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDebug:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}
