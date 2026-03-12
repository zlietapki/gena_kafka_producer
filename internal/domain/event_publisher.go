package domain

import "context"

type IEventPublisher interface {
	Publish(ctx context.Context, events []Event) error
}
