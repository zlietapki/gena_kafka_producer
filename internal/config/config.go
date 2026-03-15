// start name:top
package config

// start name:import type:merge
import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/zlietapki/gena/internal/kafka"
)

// start name:config_struct
type Config struct {
	Env string `envconfig:"ENV"`
	//start name:config_fields type:add
	Kafka kafka.Config
	//start name:post_config_fields
}

func New() Config {
	godotenv.Load()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		slog.Error("config parse failed", "err", err)
		os.Exit(1)
	}

	return cfg
}
