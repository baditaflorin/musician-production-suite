package config

import (
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIAddr            string        `envconfig:"API_ADDR" default:":8080"`
	StorageDir         string        `envconfig:"STORAGE_DIR" default:"./tmp/jobs"`
	PublicBaseURL      string        `envconfig:"PUBLIC_BASE_URL" default:"http://localhost:8080"`
	CORSAllowedOrigins []string      `envconfig:"CORS_ALLOWED_ORIGINS" default:"http://localhost:5173,http://127.0.0.1:5173"`
	MaxUploadMB        int64         `envconfig:"MAX_UPLOAD_MB" default:"250"`
	JobTTLHours        time.Duration `envconfig:"JOB_TTL_HOURS" default:"24h"`
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("MPS", &cfg); err != nil {
		return Config{}, err
	}
	if cfg.PublicBaseURL != "" {
		cfg.PublicBaseURL = strings.TrimRight(cfg.PublicBaseURL, "/")
	}
	return cfg, nil
}
