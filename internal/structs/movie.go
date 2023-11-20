package structs

import "time"

type Movie struct {
	Name          string
	Timestamp     time.Time
	TimestampNext time.Time
}
