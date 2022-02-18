package schedule

import (
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
		return slack.Attachment{
			Title: "Something went wrong",
			Text:  ":cry: O wow this is embarrassing... I was unable to fetch the schedule overview. try again?",
		}, nil
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
		Title: "Upcoming On Call Shifts",
		// Text:   ":calendar: Upcoming On Call Shifts",
		Blocks: blocks,
	}, nil
}
