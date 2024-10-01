package main

import (
	"log/slog"
	"url-shortener/internal/config"
	"url-shortener/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.Setup(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
}
