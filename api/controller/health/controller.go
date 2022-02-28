package health

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syllabix/oncall/api/rest"
	"go.uber.org/zap"
)

type Controller struct {
	log *zap.Logger
}

func NewController() rest.Controller {
	return rest.MakeController(
		&Controller{},
	)
}

func (ctrl *Controller) Register(router *httprouter.Router) {
	router.GET("/healthz", ctrl.HealthCheck)
}

func (ctrl *Controller) HealthCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("OK"))
}
