package usecase

import (
	"context"
	"time"

	"github.com/zlietapki/gena/internal/domain"
)

func (u Usecase) Example(ctx context.Context) error {
	event := domain.Event{
		Key:       "my_key",
		Payload:   "some payload",
		Timestamp: time.Time{},
	}

	err := u.eventPublisher.Publish(ctx, []domain.Event{event})
	if err != nil {
		return err
	}

	return nil
}
