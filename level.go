package heiloger

import "log/slog"

// LogLevel defines the level of logging.
type LogLevel string

// LogLevel constants define the levels of logging, from debug to error.
const (
	LevelDebug LogLevel = "debug" // LevelDebug represents detailed debugging information.
	LevelInfo  LogLevel = "info"  // LevelInfo represents general informational messages.
	LevelWarn  LogLevel = "warn"  // LevelWarn indicates potentially harmful situations.
	LevelError LogLevel = "error" // LevelError represents error events that might still allow the application to continue running.
)

func slogLevel(level LogLevel) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
