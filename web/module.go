package web

import (
	"github.com/syllabix/oncall/web/home"
	"go.uber.org/fx"
)

// Module contains all the dependencies for handling web requests
var Module = fx.Provide(
	home.NewPage,
	NewServer,
)
