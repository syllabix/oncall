package schedule

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/service/schedule/oncall"
	"github.com/syllabix/oncall/slack/forminput"
)

func scheduleFromState(state *slack.BlockActionStates) oncall.Schedule {
	name := state.Values[forminput.CreateScheduleName][forminput.CreateScheduleNameInput]
	fmt.Printf("\n schedule name input: %+v\n", name)

	start := state.Values[forminput.CreateScheduleStartTime][forminput.CreateScheduleStartTimeInput]
	fmt.Printf("\n schedule start input: %+v\n", start)

	// start := state.Values[forminput.CreateScheduleStartTime]
	// end := state.Values[forminput.CreateScheduleEndTime]
	// interval := state.Values[forminput.CreateScheduleInterval]

	return oncall.Schedule{}
}

func timeInUTC(clocktime, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{},
			fmt.Errorf("failed to load users timezone location: %w", err)
	}

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
