package main

import (
	"flag"
	"net/http"

	"github.com/syllabix/logger"
	"github.com/syllabix/oncall/common/banner"
	"github.com/syllabix/oncall/oncall"
	"go.uber.org/zap"
)

func main() {
	ping := flag.Bool("ping", false, "pings the healthz endpoint")
	flag.Parse()

	if *ping {
		log := logger.New()
		res, err := http.Get("https://biller-oncall.herokuapp.com/healthz")
		if err != nil {
			log.Error("failed to ping", zap.Error(err))
		}
		log.Info("ping completed", zap.Int("status-code", res.StatusCode))
		return
	}

	banner.Print()

	app := oncall.NewApplication()
	if app.Err() != nil {
		logger.New().
			Error("unable to start backend application",
				zap.Error(app.Err()),
			)
	}

	app.Run()
}
