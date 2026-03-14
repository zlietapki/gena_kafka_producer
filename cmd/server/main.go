// start name:top
package main

// start name:imports type:merge
import (
	"context"
	"log/slog"
	"os"

	"github.com/zlietapki/gena/internal/config"
	"github.com/zlietapki/gena/internal/kafka"
	"github.com/zlietapki/gena/internal/usecase"
)

// start name:main
func main() {
	cfg := config.New()

	// start name:usecase_deps type:add
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

	// start name:new_usecase
	uc := usecase.New(usecase.Depends{
		// start name:usecase_deps_objs type:add
		EventPublisher: producer,
		// start name:post_usecase_deps_objs
	})

	//start name:after_usecase type:add
	ctx := context.Background()

	if err = uc.Example(ctx); err != nil {
		slog.Error("usecase failed", "err", err)
		os.Exit(1)
	}

	// start name:bottom
}
