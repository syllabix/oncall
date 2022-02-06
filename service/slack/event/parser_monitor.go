package event

import (
	"io"

	"github.com/slack-go/slack/slackevents"
	"github.com/syllabix/oncall/config"
	"go.uber.org/zap"
)

func NewParser(settings config.SlackSettings, log *zap.Logger) Parser {
	return &parsemonitor{
		parser: &parser{
			parse: slackevents.ParseEvent,
			option: slackevents.OptionVerifyToken(
				slackevents.TokenComparator{
					VerificationToken: settings.VerificationToken,
				},
			),
		},
		log: log,
	}
}

// parsemonitor is an observability wrapper around
// an event Parser. It's primary purpose is to decouple
// monitoring code from actual parsing code
type parsemonitor struct {
	parser Parser
	log    *zap.Logger
}

func (p *parsemonitor) Parse(r io.ReadCloser) (Data, error) {
	data, err := p.parser.Parse(r)
	if err != nil {
		p.log.Error("event parsing failed", zap.Error(err))
	}
	return data, err
}
