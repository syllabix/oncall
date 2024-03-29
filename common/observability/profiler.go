package observabilty

import (
	"net/http"

	"github.com/syllabix/oncall/config"
	"go.uber.org/zap"

	// register application profiling
	_ "net/http/pprof"
)

// Start spins up a server with standard endpoints for running pprof
func StartProfiler(settings config.ServerSettings, log *zap.Logger) {
	log.Info("starting pprof profiler", zap.String("port", settings.ProfilerPort))
	go func() {
		err := http.ListenAndServe(settings.Host+":"+settings.ProfilerPort, nil)
		if err != nil {
			log.Warn("unable to start profiler", zap.String("reason", err.Error()))
		}
	}()
}
