package schedule

import (
	"errors"
	"fmt"
	"time"

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
	schedule, err := c.mapToSchedule(callback)
	if err != nil {
		return c.respond(callback.Channel.ID, callback.ResponseURL,
			":warning: Please provide a name, start time, end time and rotation interval to set up a new schedule",
		)
	}

	schedule, err = c.scheduler.Create(schedule)
	if err != nil {
		return c.respond(callback.Channel.ID, callback.ResponseURL,
			":cry: O no... something did not go quite right when setting up your new schedule",
		)
	}

	_, err = c.client.SetTopicOfConversation(callback.Channel.ID, "Active on call schedule: "+schedule.Name)
	if err != nil {
		fmt.Println(err)
	}

	return c.respond(callback.Channel.ID, callback.ResponseURL,
		":white_check_mark: Your schedule has been created. Add team members to the rotation with */add* command",
	)
}

func (c *creator) respond(channelID, url, msg string) error {
	_, _, err := c.client.PostMessage(
		channelID,
		slack.MsgOptionReplaceOriginal(url),
		slack.MsgOptionText(msg, false),
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
	if len(name.Value) < 1 {
		return oncall.Schedule{}, errors.New("name required")
	}

	start := callback.
		BlockActionState.Values[forminput.CreateScheduleStartTime][forminput.CreateScheduleStartTimeInput]
	if len(start.SelectedTime) < 1 {
		return oncall.Schedule{}, errors.New("start time required")
	}

	end := callback.
		BlockActionState.Values[forminput.CreateScheduleEndTime][forminput.CreateScheduleEndTimeInput]
	if len(end.SelectedTime) < 1 {
		return oncall.Schedule{}, errors.New("end time required")
	}

	interval := callback.
		BlockActionState.Values[forminput.CreateScheduleInterval][forminput.CreateScheduleIntervalInput]
	if len(interval.SelectedOption.Value) < 1 {
		return oncall.Schedule{}, errors.New("interval is required")
	}

	location, err := time.LoadLocation(user.TZ)
	if err != nil {
		return oncall.Schedule{},
			fmt.Errorf("failed to load users timezone location: %w", err)
	}

	startTime, err := timeInUTC(start.SelectedTime, location)
	if err != nil {
		return oncall.Schedule{}, fmt.Errorf("start time invalid: %w", err)
	}

	endTime, err := timeInUTC(end.SelectedTime, location)
	if err != nil {
		return oncall.Schedule{}, fmt.Errorf("start time invalid: %w", err)
	}

	return oncall.Schedule{
		ChannelID: callback.Channel.ID,
		TeamID:    callback.Team.ID,
		Name:      name.Value,
		StartTime: startTime,
		EndTime:   endTime,
		Interval:  oncall.Interval(interval.SelectedOption.Value),
	}, nil
}
