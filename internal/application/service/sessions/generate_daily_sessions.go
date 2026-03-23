package sessions

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
)

type GenerateDailySessionsUseCase struct {
	routineRepo        port.RoutineRepository
	workoutSessionRepo port.WorkoutSessionRepository
}

func NewGenerateDailySessionsUseCase(
	routineRepo port.RoutineRepository,
	workoutSessionRepo port.WorkoutSessionRepository,
) *GenerateDailySessionsUseCase {
	return &GenerateDailySessionsUseCase{
		routineRepo:        routineRepo,
		workoutSessionRepo: workoutSessionRepo,
	}
}

// goWeekdayToISO converts Go's time.Weekday (0=Sunday) to ISO 8601 (1=Monday..7=Sunday).
func goWeekdayToISO(wd time.Weekday) int16 {
	if wd == time.Sunday {
		return 7
	}
	return int16(wd)
}

func (uc *GenerateDailySessionsUseCase) Execute(ctx context.Context, date time.Time) (*dto.GenerateDailySessionsResponse, error) {
	isoDay := goWeekdayToISO(date.Weekday())

	// Active routines scheduled for this day that don't already have a session
	routines, err := uc.routineRepo.GetRoutinesToGenerateSessions(ctx, isoDay, date)
	if err != nil {
		return nil, err
	}

	if len(routines) == 0 {
		return &dto.GenerateDailySessionsResponse{SessionsCreated: 0}, nil
	}

	sessions := make([]*entity.WorkoutSession, len(routines))
	for i, r := range routines {
		sessions[i] = entity.NewWorkoutSession(r.ID, isoDay, date)
	}

	if err := uc.workoutSessionRepo.CreateBulk(ctx, sessions); err != nil {
		return nil, err
	}

	return &dto.GenerateDailySessionsResponse{
		SessionsCreated: len(sessions),
	}, nil
}
