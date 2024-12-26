package heiloger

import (
	"context"
	"log/slog"
	"time"
)

type (
	// Logger is the main struct for managing logging configurations and state
	Logger struct {
		currentDate time.Time
		*slog.Logger
		cfg *Config
	}
)

// NewLogger creates a new logger instance with default configurations
func NewLogger() Interface {
	return build(DefaultConfig())
}

// NewLoggerWithConfig creates a new logger instance with custom configurations
func NewLoggerWithConfig(cfg *Config) Interface {
	cfg.Validate()
	return build(cfg)
}

// build create logger with cfg
func build(cfg *Config) Interface {
	return &Logger{
		currentDate: time.Now(),
		cfg:         cfg,
		Logger:      slog.New(GetHandler(cfg)),
	}
}

type Interface interface {
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}
