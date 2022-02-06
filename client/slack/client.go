package slack

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/syllabix/oncall/config"
)

func NewClient(settings config.SlackSettings) (*slack.Client, error) {
	client := slack.New(
		settings.AuthToken,
		slack.OptionDebug(settings.DebugMode),
		slack.OptionAppLevelToken(settings.AppToken),
	)

	_, err := client.AuthTest()
	if err != nil {
		return nil, fmt.Errorf("slack client failed to authorize: %w", err)
	}

	return client, nil
}
