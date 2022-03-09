package event

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/slack-go/slack/slackevents"
	"github.com/syllabix/oncall/config"
	"go.uber.org/zap"
)

var (
	ErrParseError = errors.New("failed to parse slack event")
)

// Data contains all of the information provided
// in a slack event request
type Data struct {
	// The raw representation of the api event
	RawMessage json.RawMessage

	slackevents.EventsAPIEvent
}

// Parser attempts to parse a Slack event from the provided
// read closer
type Parser interface {
	Parse(io.ReadCloser) (Data, error)
}

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

type parser struct {
	option slackevents.Option
	parse  func(
		rawEvent json.RawMessage,
		opts ...slackevents.Option,
	) (slackevents.EventsAPIEvent, error)
}

func (p *parser) Parse(req io.ReadCloser) (Data, error) {
	var payload bytes.Buffer
	_, err := io.Copy(&payload, req)
	if err != nil {
		return Data{}, fmt.Errorf("%w: %v", ErrParseError, err)
	}

	event, err := p.parse(payload.Bytes(), p.option)
	if err != nil {
		return Data{}, fmt.Errorf("%w: %v", ErrParseError, err)
	}

	return Data{
		RawMessage:     payload.Bytes(),
		EventsAPIEvent: event,
	}, nil
}
