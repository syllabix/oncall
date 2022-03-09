package event

import (
	"io"

	"go.uber.org/zap"
)

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
