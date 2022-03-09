package schedule

import (
	"errors"
	"fmt"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/service/schedule"
)

type Handler interface {
	GetUpcomingShifts(channelID string) (slack.Attachment, error)
}

func NewHandler(manager schedule.Manager) Handler {
	return &handler{manager}
}

type handler struct {
	manager schedule.Manager
}

func (h *handler) GetUpcomingShifts(channelID string) (slack.Attachment, error) {
	shifts, err := h.manager.GetNextShifts(channelID)
	if err != nil {
		if errors.Is(err, schedule.ErrNotFound) {
			return slack.Attachment{
				Title: "Schedule not setup",
				Text:  ":warning: There is no on call schedule set up in this channel. You can get one setup by using the */schedule create* command",
			}, nil
		}

		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":cry: O wow this is embarrassing... I was unable to fetch the schedule overview. try again?",
		}, err
	}

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewSectionBlock(
				slack.NewTextBlockObject(
					slack.MarkdownType,
					":calendar: Upcoming On Call Shifts",
					false, false,
				),
				nil, nil,
			),
		},
	}

	if len(shifts) < 1 {
		blocks.BlockSet = append(blocks.BlockSet,
			slack.NewSectionBlock(
				slack.NewTextBlockObject(
					slack.MarkdownType,
					`:thinking_face: Hmm... doesn't look like there is anyone assigned to this schedule.

Use the */add* command to get team members in the rotation`,
					false, false,
				),
				nil, nil,
			),
		)
	}

	for _, shift := range shifts {
		blocks.BlockSet = append(blocks.BlockSet,
			slack.NewSectionBlock(
				nil,
				[]*slack.TextBlockObject{
					slack.NewTextBlockObject(
						slack.MarkdownType,
						shift.StartTime.Format("Monday Jan 2"),
						false, false,
					),
					slack.NewTextBlockObject(
						slack.MarkdownType,
						fmt.Sprintf("%s %s - %s", shift.FirstName, shift.LastName, shift.SlackHandle),
						false, false,
					),
				},
				nil,
			),
		)
	}

	return slack.Attachment{
		Title:  "Upcoming On Call Shifts",
		Blocks: blocks,
	}, nil
}
