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
	// Проверяем, указан ли путь к директории
	if cfg.Rotation.OutputDirectory == "" {
		log.Fatalf("Output directory is not specified")
	}

	// Проверяем, создана ли директория, если нет - создаем
	if _, err := os.Stat(cfg.Rotation.OutputDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(cfg.Rotation.OutputDirectory, 0755); err != nil {
			log.Fatalf("Error creating log directory: %s", err)
		}
	}

	// Если RotateDaily выключен, пишем все в один файл
	if !cfg.Rotation.RotateDaily {
		// Создаем путь для логов
		path := cfg.Rotation.OutputDirectory + "/" + cfg.Rotation.FileName

		// Проверяем, существует ли файл, если нет - создаем
		if _, err := os.Stat(path); os.IsNotExist(err) {
			file, err := os.Create(path)
			if err != nil {
				log.Fatalf("Error creating log file: %s", err)
			}
			_ = file.Close() // Закрываем файл, чтобы потом открыть его для записи
		}

		// Открываем файл для записи (с возможностью добавления данных)
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening log file: %s", err)
		}

		// Если вывод в консоль включен, комбинируем запись в файл и консоль
		if cfg.OutputToConsole {
			return io.MultiWriter(os.Stdout, file)
		}

		// Если только в файл
		return file
	}

	// Если RotateDaily включен, настраиваем ротацию по времени
	path := cfg.Rotation.OutputDirectory + "/%Y-%m-%d_" + cfg.Rotation.FileName
	writer, err := rotatelogs.New(
		path,
		rotatelogs.WithMaxAge(cfg.Rotation.MaxAge),
		rotatelogs.WithRotationTime(cfg.Rotation.RotationTime),
	)
	if err != nil {
		log.Fatalf("Error setting rotation: %s", err)
	}

	// Если вывод в консоль включен, комбинируем запись в файл и консоль
	if cfg.OutputToConsole {
		return io.MultiWriter(os.Stdout, writer)
	}

	// Если только в файл
	return writer
}
