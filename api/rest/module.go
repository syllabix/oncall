package rest

import "go.uber.org/fx"

// Module contains all the dependencies that power
// rest api routing
var Module = fx.Provide(
	NewRouter,
)
