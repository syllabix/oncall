package api

import (
	"github.com/syllabix/oncall/api/controller"
	"github.com/syllabix/oncall/api/middleware"
	"github.com/syllabix/oncall/api/rest"
	"go.uber.org/fx"
)

// Module contains all the dependencies for routing
// and handling API requests
var Module = fx.Options(
	controller.Module,
	middleware.Module,
	rest.Module,
)
