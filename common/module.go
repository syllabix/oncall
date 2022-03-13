package common

import (
	observabilty "github.com/syllabix/oncall/common/observability"
	"go.uber.org/fx"
)

// Module contains dependencies that are used
// across the project
var Module = fx.Provide(
	// db.Setup,
	observabilty.NewLogger,
)
