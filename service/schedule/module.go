package schedule

import (
	"go.uber.org/fx"
)

// Module contains all services that power handling
// on call schedules
var Module = fx.Provide(
	NewManager,
)
