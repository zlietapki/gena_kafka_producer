package usecase

import "github.com/zlietapki/boilerplate/internal/domain"

var _ domain.IUsecase = (*Usecase)(nil) // compile-time check

type Usecase struct {
	eventPublisher domain.IEventPublisher
}

type Depends struct {
	EventPublisher domain.IEventPublisher
}

func New(deps Depends) *Usecase {
	return &Usecase{
		eventPublisher: deps.EventPublisher,
	}
}
