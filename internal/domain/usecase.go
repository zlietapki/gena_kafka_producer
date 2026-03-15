// start name:top
package domain

//start name:imports type:merge
import (
	"context"
)

// start name:usecase

//go:generate minimock -g -i IUsecase -o mocks -s _mock.gen.go
type IUsecase interface {
	//start name:methods type:merge
	Example(ctx context.Context) error
	//start name:post_methods
}
