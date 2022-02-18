package notifications

import "go.uber.org/fx"

// Module contains all of the dependencies
// that power slack scheduling notification
var Module = fx.Provide(
	NewService,
)
