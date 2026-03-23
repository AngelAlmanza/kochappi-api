package cron

import (
	"context"
	"time"

	"kochappi/internal/application/service/sessions"
	"kochappi/internal/shared/logger"

	"github.com/robfig/cron/v3"
)

type DailySessionsCron struct {
	cron    *cron.Cron
	useCase *sessions.GenerateDailySessionsUseCase
}

func NewDailySessionsCron(useCase *sessions.GenerateDailySessionsUseCase) *DailySessionsCron {
	return &DailySessionsCron{
		cron:    cron.New(cron.WithSeconds()),
		useCase: useCase,
	}
}

func (d *DailySessionsCron) Start() {
	// 10:00 UTC = 3:00 AM America/Mazatlan (UTC-7, no DST)
	_, err := d.cron.AddFunc("0 0 10 * * *", func() {
		today := time.Now().UTC().Truncate(24 * time.Hour)
		logger.Info.Printf("Generating daily workout sessions for %s", today.Format(time.DateOnly))

		result, err := d.useCase.Execute(context.Background(), today)
		if err != nil {
			logger.Error.Printf("Failed to generate daily sessions: %v", err)
			return
		}

		logger.Info.Printf("Generated %d workout sessions", result.SessionsCreated)
	})
	if err != nil {
		logger.Error.Printf("Failed to schedule daily sessions cron: %v", err)
		return
	}

	d.cron.Start()
	logger.Info.Println("Daily sessions cron started (10:00 UTC / 3:00 AM Mazatlan)")
}

func (d *DailySessionsCron) Stop() {
	d.cron.Stop()
}
