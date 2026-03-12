package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/zlietapki/boilerplate/internal/config"
	"github.com/zlietapki/boilerplate/internal/kafka"
	"github.com/zlietapki/boilerplate/internal/usecase"
)

func main() {
	cfg := config.New()

	kafkaClient, err := kafka.NewClient(cfg.Kafka)
	if err != nil {
		slog.Error("kafka init failed", "err", err)
		os.Exit(1)
	}
	defer func() {
		err = kafkaClient.Stop()
		if err != nil {
			slog.Error("kafka client stop failed", "err", err)
		}
	}()

	producer := kafka.NewProducer(kafkaClient.Kgo())

	uc := usecase.New(usecase.Depends{
		EventPublisher: producer,
	})

	ctx := context.Background()

	if err = uc.Example(ctx); err != nil {
		slog.Error("usecase failed", "err", err)
		os.Exit(1)
	}
}
