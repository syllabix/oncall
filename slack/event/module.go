package event

import (
	"go.uber.org/fx"
)

// Module contains all services that power handling
// interactions with Slack Event api calls
var Module = fx.Provide(
	NewParser,
	NewHandler,
)
