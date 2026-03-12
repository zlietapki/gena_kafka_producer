package domain

import "time"

type Event struct {
	Key       string
	Payload   any
	Timestamp time.Time
}
