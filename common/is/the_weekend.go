package is

import "time"

func TheWeekend() bool {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		return true

	default:
		return false
	}
}

func NotTheWeekend() bool {
	return !TheWeekend()
}
