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

func NewController(log *zap.Logger) rest.Controller {
	return rest.MakeController(
		&Controller{log: log},
	)
}

func (ctrl *Controller) Register(router *httprouter.Router) {
	router.GET("/healthz", ctrl.HealthCheck)
}

func (ctrl *Controller) HealthCheck(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctrl.log.Info("health check pinged")
	w.Write([]byte("OK"))
}
