package schedule

import "github.com/syllabix/oncall/slack/forminput"

var newForm = Form{
	Blocks: []Block{
		// schedule name text input
		{
			Type:    "input",
			BlockID: forminput.CreateScheduleName,
			Element: &FormElement{
				Type:     "plain_text_input",
				ActionID: forminput.CreateScheduleNameInput,
				Placeholder: Label{
					Type: "plain_text",
					Text: "Schedule Name",
				},
			},
			Label: &Label{
				Type: "plain_text",
				Text: "What is the name of the schedule?",
			},
		},
		// shift start time picker in put
		{
			Type:    "input",
			BlockID: forminput.CreateScheduleStartTime,
			Element: &FormElement{
				Type:     "timepicker",
				ActionID: forminput.CreateScheduleStartTimeInput,
				Placeholder: Label{
					Type: "plain_text",
					Text: "Select time",
				},
			},
			Label: &Label{
				Type: "plain_text",
				Text: "When does the shift start?",
			},
		},
		// shift end time picker in put
		{
			Type:    "input",
			BlockID: forminput.CreateScheduleEndTime,
			Element: &FormElement{
				Type:     "timepicker",
				ActionID: forminput.CreateScheduleEndTimeInput,
				Placeholder: Label{
					Type: "plain_text",
					Text: "Select time",
				},
			},
			Label: &Label{
				Type: "plain_text",
				Text: "When does the shift end?",
			},
		},
		// shift end time picker input
		{
			Type:    "input",
			BlockID: forminput.CreateScheduleInterval,
			Element: &FormElement{
				Type:     "static_select",
				ActionID: forminput.CreateScheduleIntervalInput,
				Placeholder: Label{
					Type: "plain_text",
					Text: "Select interval",
				},
				Options: []Option{
					{
						Text: Label{
							Type: "plain_text",
							Text: "Daily",
						},
						Value: "daily",
					},
					{
						Text: Label{
							Type: "plain_text",
							Text: "Weekly",
						},
						Value: "weekly",
					},
					{
						Text: Label{
							Type: "plain_text",
							Text: "Bi-Weekly",
						},
						Value: "bi-weekly",
					},
				},
			},
			Label: &Label{
				Type: "plain_text",
				Text: "When does the shift rotate?",
			},
		},
		// submit action block
		{
			Type:    "actions",
			BlockID: "create_new_schedule_submit_block",
			Actions: []ActionElement{
				{
					Type:     "button",
					ActionID: "create_new_schedule_action_submit",
					Text: Label{
						Type: "plain_text",
						Text: "Create",
					},
					Value: "create_new_schedule",
				},
			},
		},
	},
}
