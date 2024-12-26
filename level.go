package heiloger

import "log/slog"

// LogLevel defines the level of logging.
type LogLevel int

// LogLevel constants define the levels of logging, from debug to error.
const (
	LevelDebug LogLevel = iota + 4 // LevelDebug represents detailed debugging information.
	LevelInfo                      // LevelInfo represents general informational messages.
	LevelWarn                      // LevelWarn indicates potentially harmful situations.
	LevelError                     // LevelError represents error events that might still allow the application to continue running.
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
