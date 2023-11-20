package utils

import (
	"fmt"
	"time"
)

func FormatDuration(duration time.Duration) string {
	duration = duration.Round(time.Minute)
	hours := duration / time.Hour
	duration -= hours * time.Hour
	minutes := duration / time.Minute

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
