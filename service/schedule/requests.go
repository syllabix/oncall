package schedule

type CreateFormRequest struct {
	Type                string    `json:"type"`
	User                User      `json:"user"`
	APIAppID            string    `json:"api_app_id"`
	Token               string    `json:"token"`
	Container           Container `json:"container"`
	TriggerID           string    `json:"trigger_id"`
	Team                Team      `json:"team"`
	IsEnterpriseInstall bool      `json:"is_enterprise_install"`
	Channel             Channel   `json:"channel"`
	State               State     `json:"state"`
	ResponseURL         string    `json:"response_url"`
	Actions             []Action  `json:"actions"`
}

type Action struct {
	ActionID string `json:"action_id"`
	BlockID  string `json:"block_id"`
	Text     Text   `json:"text"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	ActionTs string `json:"action_ts"`
}

type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Container struct {
	Type        string `json:"type"`
	MessageTs   string `json:"message_ts"`
	ChannelID   string `json:"channel_id"`
	IsEphemeral bool   `json:"is_ephemeral"`
}

type State struct {
	Values Values `json:"values"`
}

type Values struct {
	ScheduleNameInputBlock   ScheduleNameInputBlock   `json:"schedule_name_input_block"`
	StartTimeInputBlock      StartTimeInputBlock      `json:"start_time_input_block"`
	EndTimeInputBlock        EndTimeInputBlock        `json:"end_time_input_block"`
	IntervalSelectInputBlock IntervalSelectInputBlock `json:"interval_select_input_block"`
}

type EndTimeInputBlock struct {
	EndTimeInput TimeInput `json:"end_time_input"`
}

type TimeInput struct {
	Type         string `json:"type"`
	SelectedTime string `json:"selected_time"`
}

type IntervalSelectInputBlock struct {
	IntervalSelectInput IntervalSelectInput `json:"interval_select_input"`
}

type IntervalSelectInput struct {
	Type           string         `json:"type"`
	SelectedOption SelectedOption `json:"selected_option"`
}

type SelectedOption struct {
	Text  Text   `json:"text"`
	Value string `json:"value"`
}

type ScheduleNameInputBlock struct {
	ScheduleNameInput ScheduleNameInput `json:"schedule_name_input"`
}

type ScheduleNameInput struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type StartTimeInputBlock struct {
	StartTimeInput TimeInput `json:"start_time_input"`
}

type Team struct {
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	TeamID   string `json:"team_id"`
}
