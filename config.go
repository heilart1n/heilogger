package heiloger

import "time"

type (
	// LogFormatType defines the output format of the logs (e.g., JSON, pretty-JSON, text).
	LogFormatType string

	// Config holds the configuration options for initializing the Logger.
	Config struct {
		// OutputFormat The format for log output (e.g., "json", "pretty-json", "text").
		OutputFormat LogFormatType `json:"output_format" yaml:"output_format"`
		// EnableFileOutput If true, logs will be saved to a file.
		EnableFileOutput bool `json:"enable_file_output" yaml:"enable_file_output"`
		// OutputLevel The minimum log level that will be output (e.g., info, debug).
		OutputLevel LogLevel `json:"output_level" yaml:"output_level"`
		// WithSource If true, includes the source file and line number in logs.
		WithSource bool `json:"with_source" yaml:"with_source"`
		// OutputToConsole If true, logs will also be printed to the console.
		OutputToConsole bool `json:"output_to_console" yaml:"output_to_console"`
		// Rotation
		Rotation *ConfigRotation `yaml:"rotation" json:"rotation"`
	}
	ConfigRotation struct {
		// RotateDaily If true, log file rotates daily.
		RotateDaily bool `json:"rotate_daily" yaml:"rotate_daily"`
		// MaxAge The maximum number of days to retain old log files.
		MaxAge time.Duration `json:"max_age" yaml:"max_age"`
		// RotationTime Time for rotation file (for example 24 hours)
		RotationTime time.Duration `json:"rotation_time" yaml:"rotation_time"`
		// OutputDirectory directory where creates logs files
		OutputDirectory string `json:"output_directory" yaml:"output_directory"`
		// FileName file name creates after rotate {file_name}_2020-11-26.log
		FileName string `json:"file_name" yaml:"file_name"`
	}
)

// LogFormatType Log Formats
const (
	LogFormatJSON       LogFormatType = "json"
	LogFormatPrettyJSON LogFormatType = "pretty-json"
	LogFormatText       LogFormatType = "text"
)

// Default values for Config settings.
const (
	DefaultLogFormat        LogFormatType = "text"    // Default output format is plain text.
	DefaultOutputDirectory                = "./logs"  // Default path for log files.
	DefaultFileName                       = "log.txt" // Default file name
	DefaultEnableFileOutput               = false     // File output is enabled by default.
	DefaultLogLevel                       = LevelInfo // Default log level is info.
	DefaultWithSource                     = true      // Source information is included by default.
	DefaultOutputToConsole                = true
	DefaultRotateDaily                    = false           // Rotate log files daily.
	DefaultMaxAge                         = 168 * time.Hour // Keep log files for 30 days.
	DefaultRotationTime                   = 24 * time.Hour
)

// DefaultConfig returns a Config instance with default settings.
func DefaultConfig() *Config {
	return &Config{
		OutputFormat:     DefaultLogFormat,
		EnableFileOutput: DefaultEnableFileOutput,
		OutputLevel:      DefaultLogLevel,
		WithSource:       DefaultWithSource,
		OutputToConsole:  DefaultOutputToConsole,
		Rotation:         DefaultRotation(),
	}
}

func DefaultRotation() *ConfigRotation {
	return &ConfigRotation{
		RotateDaily:     DefaultRotateDaily,
		MaxAge:          DefaultMaxAge,
		FileName:        DefaultFileName,
		RotationTime:    DefaultRotationTime,
		OutputDirectory: DefaultOutputDirectory,
	}
}

// Validate checks the Config fields and sets default values if necessary.
func (cfg *Config) Validate() {
	// Set default log format if not provided or invalid
	if cfg.OutputFormat == "" {
		cfg.OutputFormat = DefaultLogFormat
	} else {
		switch cfg.OutputFormat {
		case LogFormatJSON, LogFormatPrettyJSON, LogFormatText:
			// Valid formats, do nothing
		default:
			cfg.OutputFormat = DefaultLogFormat
		}
	}

	// Set default log level if invalid
	if cfg.OutputLevel < LevelDebug || cfg.OutputLevel > LevelError {
		cfg.OutputLevel = DefaultLogLevel
	}

	if cfg.Rotation == nil {
		cfg.Rotation = DefaultRotation()
	}
}

func (cfg *Config) SetOutputFormat(logFormat LogFormatType) *Config {
	cfg.OutputFormat = logFormat
	return cfg
}

func (cfg *Config) SetOutputLevel(logLevel LogLevel) *Config {
	cfg.OutputLevel = logLevel
	return cfg
}

func (cfg *Config) SetConfigRotation(rotation *ConfigRotation) *Config {
	cfg.Rotation = rotation
	return cfg
}

func (cfg *Config) SetWithSource(source bool) *Config {
	cfg.WithSource = source
	return cfg
}

func (cfg *Config) SetEnableFileOutput(enable bool) *Config {
	cfg.EnableFileOutput = enable
	return cfg
}

func (cfg *ConfigRotation) SetRotateDaily(rotateDaily bool) *ConfigRotation {
	cfg.RotateDaily = rotateDaily
	return cfg
}

func (cfg *ConfigRotation) SetMaxAge(maxAge time.Duration) *ConfigRotation {
	cfg.MaxAge = maxAge
	return cfg
}

func (cfg *ConfigRotation) SetRotationTime(rotationTime time.Duration) *ConfigRotation {
	cfg.RotationTime = rotationTime
	return cfg
}

func (cfg *ConfigRotation) SetOutputDirectory(outputDirectory string) *ConfigRotation {
	cfg.OutputDirectory = outputDirectory
	return cfg
}

func (cfg *ConfigRotation) SetOutputFileName(outputFileName string) *ConfigRotation {
	cfg.FileName = outputFileName
	return cfg
}
