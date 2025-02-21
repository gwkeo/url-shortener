package main

import (
	"github.com/gwkeo/url-shortener/internal/config"
	"github.com/gwkeo/url-shortener/internal/http-server/router"
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
	router := router.New(log)
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
