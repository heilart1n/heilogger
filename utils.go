package heiloger

import (
	"encoding/json"
	"errors"
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
			MaxAge       string `json:"max_age"`
			RotationTime string `json:"rotation_time"`
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
			return errors.New("invalid max_age duration format")
		}
		cfg.Rotation.MaxAge = maxAge
	}
	if aux.Rotation.RotationTime != "" {
		rotationTime, err := time.ParseDuration(aux.Rotation.RotationTime)
		if err != nil {
			return errors.New("invalid rotation_time duration format")
		}
		cfg.Rotation.RotationTime = rotationTime
	}

	return nil
}
