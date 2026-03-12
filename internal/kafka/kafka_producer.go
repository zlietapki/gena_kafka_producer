package kafka

import (
	"context"
	"log/slog"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/zlietapki/boilerplate/internal/domain"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(client *kgo.Client) *Producer {
	return &Producer{
		client: client,
	}
}

func (p Producer) Publish(ctx context.Context, events []domain.Event) error {
	var wg sync.WaitGroup

	for _, e := range events {
		record, err := toKgoRecord(e)
		if err != nil {
			return err
		}

		wg.Add(1)
		p.client.Produce(ctx, record, func(r *kgo.Record, err error) {
			defer wg.Done()

			if err != nil {
				slog.Error("Kafka produce err", "key", r.Key, "err", err)
				return
			}

			slog.Info("Delivered", "topic", r.Topic, "partition", r.Partition, "offset", r.Offset, "key", r.Key)
		})

	}
	wg.Wait()

	return nil
}
