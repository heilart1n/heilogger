package heiloger

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"
)

func Error(err error) slog.Attr {
	return slog.String("err", err.Error())
}

// UnmarshalJSON Custom UnmarshalJSON for Config
func (cfg *Config) UnmarshalJSON(data []byte) error {
	type Alias Config // Use an alias to avoid infinite recursion
	aux := &struct {
		Rotation struct {
			MaxAge          string `json:"max_age"`
			RotationTime    string `json:"rotation_time"`
			OutputDirectory string `json:"output_directory"`
			FileName        string `json:"file_name"`
			RotateDaily     bool   `json:"rotate_daily"`
		} `json:"rotation"`
		*Alias
	}{
		Alias: (*Alias)(cfg),
	}

	// Unmarshal into the auxiliary struct
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Ensure cfg.Rotation is not nil
	if cfg.Rotation == nil {
		cfg.Rotation = &ConfigRotation{}
	}

	// Parse time.Duration fields from string
	if aux.Rotation.MaxAge != "" {
		maxAge, err := time.ParseDuration(aux.Rotation.MaxAge)
		if err != nil {
			return fmt.Errorf("invalid max_age duration format: %w", err)
		}
		cfg.Rotation.MaxAge = maxAge
	}
	if aux.Rotation.RotationTime != "" {
		rotationTime, err := time.ParseDuration(aux.Rotation.RotationTime)
		if err != nil {
			return fmt.Errorf("invalid rotation_time duration format: %w", err)
		}
		cfg.Rotation.RotationTime = rotationTime
	}

	// Set other fields in Rotation
	cfg.Rotation.OutputDirectory = aux.Rotation.OutputDirectory
	cfg.Rotation.FileName = aux.Rotation.FileName
	cfg.Rotation.RotateDaily = aux.Rotation.RotateDaily

	return nil
}
