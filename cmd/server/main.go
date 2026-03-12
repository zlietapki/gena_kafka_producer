package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/zlietapki/boilerplate/internal/config"
	"github.com/zlietapki/boilerplate/internal/kafka"
	"github.com/zlietapki/boilerplate/internal/usecase"
)

func main() {
	cfg := config.New()

	kafkaClient, err := kafka.NewClient(cfg.Kafka)
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaClient.Close()
	defer kafkaClient.Flush(context.Background())

	producer := kafka.NewProducer(kafkaClient)
	uc := usecase.New(usecase.Depends{
		EventPublisher: producer,
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err = uc.Example(ctx); err != nil {
		log.Fatal(err)
	}
}
