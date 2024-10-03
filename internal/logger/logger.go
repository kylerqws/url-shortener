package logger

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
)

func Setup(env string) *slog.Logger {
	return slog.New(*getHandler(env))
}

func getHandler(env string) *slog.Handler {
	var handler slog.Handler

	switch env {
	case config.EnvLocal:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case config.EnvDev:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case config.EnvProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return &handler
}
