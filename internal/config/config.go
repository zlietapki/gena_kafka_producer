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

// start name:post_import
type Config struct {
	Env string `envconfig:"ENV"`
	//start name:conf type:add
	Kafka kafka.Config
	//start name:rest
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
