package schedule

type CreateForm struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type    string          `json:"type"`
	BlockID string          `json:"block_id"`
	Element *FormElement    `json:"element,omitempty"`
	Label   *Label          `json:"label,omitempty"`
	Actions []ActionElement `json:"elements,omitempty"`
}

type FormElement struct {
	Type        string   `json:"type"`
	ActionID    string   `json:"action_id"`
	Placeholder Label    `json:"placeholder"`
	Options     []Option `json:"options,omitempty"`
}

type Option struct {
	Text  Label  `json:"text"`
	Value string `json:"value"`
}

type Label struct {
	Type  Type   `json:"type"`
	Text  string `json:"text"`
	Emoji *bool  `json:"emoji,omitempty"`
}

type ActionElement struct {
	Type     string `json:"type"`
	Text     Label  `json:"text"`
	Value    string `json:"value"`
	ActionID string `json:"action_id"`
}

type Type string

const (
	PlainText Type = "plain_text"
)

var newScheduleForm = CreateForm{
	Blocks: []Block{
		// schedule name text input
		{
			Type:    "input",
			BlockID: "schedule_name_input_block",
			Element: &FormElement{
				Type:     "plain_text_input",
				ActionID: "schedule_name_input",
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
			BlockID: "start_time_input_block",
			Element: &FormElement{
				Type:     "timepicker",
				ActionID: "start_time_input",
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
			BlockID: "end_time_input_block",
			Element: &FormElement{
				Type:     "timepicker",
				ActionID: "end_time_input",
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
		// shift end time picker in put
		{
			Type:    "input",
			BlockID: "interval_select_input_block",
			Element: &FormElement{
				Type:     "static_select",
				ActionID: "interval_select_input",
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
