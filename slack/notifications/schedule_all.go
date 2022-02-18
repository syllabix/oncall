package notifications

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ScheduleAll(lc fx.Lifecycle, service Service, log *zap.Logger) error {

	err := service.ScheduleAll()
	if err != nil {
		return fmt.Errorf("failed to start notification and job scheduling: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			service.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			service.Stop()
			return nil
		},
	})

	return nil
}
