// start name:top
package usecase

import "github.com/zlietapki/gena/internal/domain"

var _ domain.IUsecase = (*Usecase)(nil) // compile-time check

type Usecase struct {
	// start name:uc_struct type:add
	eventPublisher domain.IEventPublisher
	// start name:post_uc_struct
}

type Depends struct {
	// start name:deps type:add
	EventPublisher domain.IEventPublisher
	// start name:post_deps
}

func New(deps Depends) *Usecase {
	return &Usecase{
		// start name:new type:add
		eventPublisher: deps.EventPublisher,
		// start name:post_new
	}
}
