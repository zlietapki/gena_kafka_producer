package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/zlietapki/boilerplate/internal/kafka"
)

type Config struct {
	Env   string `envconfig:"ENV"`
	Kafka kafka.Config
}

func New() Config {
	godotenv.Load()

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
