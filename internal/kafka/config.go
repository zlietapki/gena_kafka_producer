package kafka

import "time"

type Config struct {
	Brokers         []string      `envconfig:"KAFKA_BROKERS"`
	SASLMechanism   string        `envconfig:"KAFKA_SASL_MECHANISM" default:"SCRAM-SHA-256"`
	User            string        `envconfig:"KAFKA_USER"`
	Password        string        `envconfig:"KAFKA_PASSWORD"`
	RequestRetries  int           `envconfig:"KAFKA_REQUEST_RETRIES" default:"3"`
	RetryTimeout    time.Duration `envconfig:"KAFKA_RETRY_TIMEOUT" default:"5s"`
	ConnIdleTimeout time.Duration `envconfig:"KAFKA_CONN_IDLE_TIMEOUT" default:"5s"`
	DialTimeout     time.Duration `envconfig:"KAFKA_DIAL_TIMEOUT" default:"5s"`
	Topic           string        `envconfig:"KAFKA_TOPIC"`
}
