package schedule

import "github.com/slack-go/slack"

func setupScheduleForm() slack.Attachment {
	attachment := slack.Attachment{}

	nameInput := slack.NewInputBlock(
		"schedule_name_block",
		slack.NewTextBlockObject(slack.PlainTextType, "What is the on call schedule name?", false, false),
		slack.NewPlainTextInputBlockElement(
			slack.NewTextBlockObject(slack.PlainTextType, "Schedule Name", false, false),
			"create_schedule_name",
		),
	)

	startTime := slack.NewAccessory(
		slack.NewTimePickerBlockElement("create_schedule_start_time"),
	)

	endTime := slack.NewAccessory(
		slack.NewTimePickerBlockElement("create_schedule_end_time"),
	)

	intervals := slack.NewOptionsSelectBlockElement(
		slack.OptTypeStatic,
		slack.NewTextBlockObject(slack.PlainTextType, "Select Interval", false, false),
		"create_schedule_rotation_interval",
		slack.NewOptionBlockObject(
			"daily",
			slack.NewTextBlockObject(slack.PlainTextType, "Daily", false, false),
			slack.NewTextBlockObject(slack.PlainTextType, "rotates team members daily", false, false),
		),
		slack.NewOptionBlockObject(
			"weekly",
			slack.NewTextBlockObject(slack.PlainTextType, "Weekly", false, false),
			slack.NewTextBlockObject(slack.PlainTextType, "rotates team members weekly", false, false),
		),
		slack.NewOptionBlockObject(
			"bi-weekly",
			slack.NewTextBlockObject(slack.PlainTextType, "Bi-Weekly", false, false),
			slack.NewTextBlockObject(slack.PlainTextType, "rotates team members every two weeks", false, false),
		),
	)

	attachment.Blocks = slack.Blocks{
		BlockSet: []slack.Block{
			nameInput,
			slack.NewSectionBlock(
				&slack.TextBlockObject{
					Type: slack.MarkdownType,
					Text: "What time does the shift start?",
				},
				nil,
				startTime,
			),
			slack.NewSectionBlock(
				&slack.TextBlockObject{
					Type: slack.MarkdownType,
					Text: "What time does the shift end?",
				},
				nil,
				endTime,
			),
			slack.NewInputBlock(
				"interval_select_block",
				slack.NewTextBlockObject(slack.PlainTextType, "How often do shifts rotate?", false, false),
				intervals,
			),
		},
	}

	return attachment
}
