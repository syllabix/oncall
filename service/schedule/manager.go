package schedule

import (
	"github.com/slack-go/slack"
)

type Manager interface {
	Prepare() slack.Attachment
}

func NewManager() Manager {
	return &manager{
		attachment: setupScheduleForm(),
	}
}

type manager struct {
	attachment slack.Attachment
}

func (m *manager) Prepare() slack.Attachment {
	return m.attachment
}
