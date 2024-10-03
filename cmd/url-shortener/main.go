package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/slg"
	"url-shortener/internal/logger"
	"url-shortener/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	log := logger.Setup(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("filed to init storage", slg.Err(err))
		os.Exit(1)
	}

	_ = storage
}
