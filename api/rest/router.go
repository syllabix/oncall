package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syllabix/oncall/api/middleware"
)

// Router is responsible for routing API requests
type Router struct {
	http.Handler
}

// NewRouter returns an http.Handler for the On Call
// rest API
func NewRouter(registered Controllers, logger *middleware.Logger) Router {

	router := httprouter.New()

	for _, r := range registered.Controllers {
		r.Register(router)
	}

	handler := middleware.Recover(
		logger.LogRequests(router),
	)

	return Router{Handler: handler}
}
