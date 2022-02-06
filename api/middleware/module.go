package middleware

import (
	"go.uber.org/fx"
)

// Module contains all the dependencies
// that power api middleware
var Module = fx.Provide(
	NewSlackVerifier,
)
