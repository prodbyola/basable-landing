package logger

import (
	"log/slog"
	"os"
	"sync"
)

type appLogger struct {
	Slog *slog.Logger
}

var (
	logger *appLogger
	once   sync.Once
)

func Slog() *appLogger {
	once.Do(func() {
		logger = &appLogger{
			Slog: slog.New(slog.NewJSONHandler(os.Stderr, nil)),
		}
	})
	return logger
}

func (l *appLogger) Info(msg string) {
	l.Slog.Info(msg)
}

func (l *appLogger) Error(msg string) {
	l.Slog.Error(msg)
}

func (l *appLogger) Debug(msg string) {
	l.Slog.Debug(msg)
}
