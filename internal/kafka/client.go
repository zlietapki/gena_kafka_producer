package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"
)

type Client struct {
	kgoClient   *kgo.Client
	stopTimeout time.Duration
}

func (c Client) Kgo() *kgo.Client {
	return c.kgoClient
}

func NewClient(cfg Config) (*Client, error) {
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

	kgoClient, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, fmt.Errorf("create kafka client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	defer cancel()
	err = kgoClient.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping kafka client: %w", err)
	}

	return &Client{
		kgoClient:   kgoClient,
		stopTimeout: cfg.StopTimeout,
	}, nil
}

func (c Client) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.stopTimeout)
	defer cancel()

	if err := c.kgoClient.Flush(ctx); err != nil {
		return fmt.Errorf("flush timeout, some records may be lost: %w", err)
	}

	c.kgoClient.Close()
	return nil
}
