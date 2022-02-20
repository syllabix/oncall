package oncall

import (
	"context"

	"github.com/syllabix/oncall/api"
	"github.com/syllabix/oncall/client"
	"github.com/syllabix/oncall/common"
	observabilty "github.com/syllabix/oncall/common/observability"
	"github.com/syllabix/oncall/config"
	"github.com/syllabix/oncall/datastore"

	"github.com/syllabix/oncall/service"
	"github.com/syllabix/oncall/slack"
	"github.com/syllabix/oncall/slack/notifications"
	"github.com/syllabix/oncall/web"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Application is a web service that powers the
// on call backend. It handles application start,
// dependency injection, and graceful shutdowns
// It wraps an uber fx app, so for detailed documentation
// please check the go doc at https://pkg.go.dev/go.uber.org/fx
type Application struct {
	*fx.App
}

// NewApplication constucts a on call backend
// application ready to be run.
func NewApplication(options ...Option) Application {
	settings := apply(options...)

	app := fx.New(
		// fx configuration
		fx.StartTimeout(settings.startTimeout),
		fx.StopTimeout(settings.stopTimeout),

		// the default fx logger will print the dependency
		// graph for debugging purposes.
		// to view it, comment out the no logger option
		// below
		fx.NopLogger,

		// application dependencies
		fx.Provide(config.Load),
		api.Module,
		common.Module,
		client.Module,
		datastore.Module,
		service.Module,
		slack.Module,
		web.Module,

		// start the engines
		fx.Invoke(start),
		fx.Invoke(notifications.ScheduleAll),
		fx.Invoke(observabilty.StartProfiler),
	)

	return Application{app}
}

// start boots up the web server hooking it into application start
// and stop lifecycles
func start(lc fx.Lifecycle, server web.Server, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			log.Info("web server listening and ready to accept connections",
				zap.String("address", server.Addr))
			return nil
		},

		OnStop: func(ctx context.Context) error {
			log.Info("powering down the on call web service...")
			return server.Shutdown(ctx)
		},
	})
}
