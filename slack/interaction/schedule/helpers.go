package schedule

import (
	"fmt"
	"time"
)

func timeInUTC(clocktime string, location *time.Location) (time.Time, error) {
	now := time.Now()
	localtime, err := time.ParseInLocation("15:04", clocktime, location)
	if err != nil {
		return time.Time{},
			fmt.Errorf("the provided clock time is not valid: %w", err)
	}
	// ensure timezone information is considered with the current year
	localtime = localtime.AddDate(now.Year(), 1, 1)
	return localtime.UTC(), nil
}
