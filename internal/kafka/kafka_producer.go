package kafka

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/zlietapki/boilerplate/internal/domain"
	"golang.org/x/sync/errgroup"
)

type Producer struct {
	producer *kgo.Client
}

func NewProducer(producer *kgo.Client) *Producer {
	return &Producer{
		producer: producer,
	}
}

func (p Producer) Publish(ctx context.Context, events []domain.Event) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, e := range events {
		record, err := toKgoRecord(e)
		if err != nil {
			return err
		}

		g.Go(func() error {
			errChan := make(chan error, 1)
			p.producer.Produce(ctx, record, func(_ *kgo.Record, err error) {
				errChan <- err
			})

			select {
			case err := <-errChan:
				return err
			case <-ctx.Done():
				return ctx.Err()
			}
		})
	}

	return g.Wait()
}

func (p Producer) Close() {
	p.producer.Close()
}
