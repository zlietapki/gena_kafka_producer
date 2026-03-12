package kafka

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
)

func NewClient(cfg Config) (*kgo.Client, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(cfg.Brokers...),
		kgo.RequestRetries(cfg.RequestRetries),
		kgo.RetryTimeout(cfg.RetryTimeout),
		kgo.DialTimeout(cfg.DialTimeout),
		kgo.DefaultProduceTopic(cfg.Topic),
	}

	if cfg.User != "" {
		switch cfg.SASLMechanism {
		case "SCRAM-SHA-256":
			opts = append(opts, kgo.SASL(
				scram.Auth{
					User: cfg.User,
					Pass: cfg.Password,
				}.AsSha256Mechanism(),
			))
		case "SCRAM-SHA-512":
			opts = append(opts, kgo.SASL(
				scram.Auth{
					User: cfg.User,
					Pass: cfg.Password,
				}.AsSha512Mechanism()),
			)
		}
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, fmt.Errorf("create kafka client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	defer cancel()
	err = client.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping kafka client: %w", err)
	}

	return client, nil
}
