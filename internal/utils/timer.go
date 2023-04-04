package utils

import (
	"fmt"
	"time"
)

func FormatTime(duration time.Duration) string {
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
