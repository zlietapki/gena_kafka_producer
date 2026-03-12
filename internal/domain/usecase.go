package domain

import "context"

type IUsecase interface {
	Example(ctx context.Context) error
}
