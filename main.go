package main

import (
	"github.com/syllabix/logger"
	"github.com/syllabix/oncall/oncall"
	"github.com/syllabix/oncall/util/banner"
	"go.uber.org/zap"
)

func main() {
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
