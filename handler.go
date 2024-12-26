package heiloger

import (
	"github.com/dusted-go/logging/prettylog"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/phsym/console-slog"
	"io"
	"log"
	"log/slog"
	"os"
)

func GetHandler(cfg *Config) slog.Handler {
	switch cfg.OutputFormat {
	case LogFormatJSON:
		return HandlerJSON(cfg)
	case LogFormatPrettyJSON:
		return HandlerPrettyJSON(cfg)
	case LogFormatText:
		return HandlerText(cfg)
	default:
		return HandlerText(cfg)
	}
}

func handlerOptions(cfg *Config) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		AddSource: cfg.WithSource,
		Level:     slogLevel(cfg.OutputLevel),
	}
}

func HandlerJSON(cfg *Config) slog.Handler {
	return slog.NewJSONHandler(withRotation(cfg), handlerOptions(cfg))
}

func HandlerPrettyJSON(cfg *Config) slog.Handler {
	handler := prettylog.New(
		handlerOptions(cfg),
		prettylog.WithDestinationWriter(withRotation(cfg)),
		prettylog.WithOutputEmptyAttrs(),
	)
	if !cfg.EnableFileOutput {
		prettylog.WithColor()(handler)
	}
	return handler
}

func HandlerText(cfg *Config) slog.Handler {
	options := &console.HandlerOptions{
		Level:      slogLevel(cfg.OutputLevel),
		TimeFormat: "3:04:05PM",
		AddSource:  cfg.WithSource,
		Theme:      console.NewDefaultTheme(),
	}
	if cfg.EnableFileOutput {
		options.NoColor = true
	}
	return console.NewHandler(withRotation(cfg), options)
}

func withRotation(cfg *Config) io.Writer {
	if !cfg.EnableFileOutput {
		return os.Stdout
	}
	// Создаём файл с ротацией
	path := cfg.OutputDirectory + "/%Y-%m-%d_" + cfg.OutputFileName

	// Настраиваем ротацию по времени
	writer, err := rotatelogs.New(
		path,
		rotatelogs.WithMaxAge(cfg.Rotation.MaxAge),
		rotatelogs.WithRotationTime(cfg.Rotation.RotationTime),
	)
	if err != nil {
		log.Fatalf("Error setting rotation: %s", err)
	}

	return io.MultiWriter(os.Stdout, writer)
}
