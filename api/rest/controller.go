package rest

import (
	"github.com/julienschmidt/httprouter"
	"go.uber.org/fx"
)

// Registerable has a single Register method which takes an http.Handler as
// a single argument - intended to be used by the implementor to register
// handlers to a router
type Registerable interface {
	Register(*httprouter.Router)
}

// Controller is an entity used to wrap Registerables and
// make them available for dependency injection
type Controller struct {
	fx.Out
	Ctrl Registerable `group:"controllers"`
}

// Controllers are a grouping of all Registerables
// intended to be used as handlers for the API
type Controllers struct {
	fx.In
	Controllers []Registerable `group:"controllers"`
}

// MakeController is a convenience utility function
// to make an api controller from a Registerable
func MakeController(r Registerable) Controller {
	return Controller{Ctrl: r}
}
