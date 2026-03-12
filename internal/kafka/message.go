package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/zlietapki/boilerplate/internal/domain"
)

func toKgoRecord(e domain.Event) (*kgo.Record, error) {
	value, err := json.Marshal(e.Payload)
	if err != nil {
		return nil, fmt.Errorf("serialize event: %w", err)
	}
	return &kgo.Record{
		Key:       []byte(e.Key),
		Value:     value,
		Timestamp: e.Timestamp,
	}, nil
}
