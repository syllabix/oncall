package schedule

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/service/schedule"
	"github.com/syllabix/oncall/service/schedule/oncall"
	"github.com/syllabix/oncall/slack/forminput"
)

type Creator interface {
	Create(callback slack.InteractionCallback) error
}

func NewCreator(
	client *slack.Client,
	scheduler schedule.Manager,
) Creator {
	return &creator{client, scheduler}
}

type creator struct {
	client    *slack.Client
	scheduler schedule.Manager
}

func (c *creator) Create(callback slack.InteractionCallback) error {

	_, err := c.mapToSchedule(callback)
	if err != nil {
		fmt.Println(err)
	}

	_, _, err = c.client.PostMessage(
		callback.Channel.ID,
		slack.MsgOptionReplaceOriginal(callback.ResponseURL),
		slack.MsgOptionText(":white_check_mark: Your schedule has been created. Add team members to the rotation with */add* command", false),
	)
	if err != nil {
		return fmt.Errorf("something went wrong posting the message: %v", err)
	}

	return nil
}

func (c *creator) mapToSchedule(callback slack.InteractionCallback) (oncall.Schedule, error) {

	user, err := c.client.GetUserInfo(callback.User.ID)
	if err != nil {
		return oncall.Schedule{}, err
	}

	name := callback.
		BlockActionState.Values[forminput.CreateScheduleName][forminput.CreateScheduleNameInput]

	start := callback.
		BlockActionState.Values[forminput.CreateScheduleStartTime][forminput.CreateScheduleStartTimeInput]

	end := callback.
		BlockActionState.Values[forminput.CreateScheduleEndTime][forminput.CreateScheduleEndTimeInput]

	interval := callback.
		BlockActionState.Values[forminput.CreateScheduleInterval][forminput.CreateScheduleIntervalInput]

	startTime, err := timeInUTC(start.SelectedTime, user.TZ)
	if err != nil {
		return oncall.Schedule{}, fmt.Errorf("start time invalid: %w", err)
	}

	endTime, err := timeInUTC(end.SelectedTime, user.TZ)
	if err != nil {
		return oncall.Schedule{}, fmt.Errorf("start time invalid: %w", err)
	}

	return oncall.Schedule{
		TeamID:    callback.Team.ID,
		Name:      name.Value,
		StartTime: startTime,
		EndTime:   endTime,
		Interval:  oncall.Interval(interval.SelectedOption.Value),
	}, nil
}
