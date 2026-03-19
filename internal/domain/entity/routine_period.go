package entity

import "time"

type RoutinePeriod struct {
	ID        int
	RoutineID int
	StartedAt time.Time
	EndedAt   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRoutinePeriod(routineID int, startedAt time.Time) *RoutinePeriod {
	now := time.Now()
	return &RoutinePeriod{
		RoutineID: routineID,
		StartedAt: startedAt,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (rp *RoutinePeriod) End(endedAt time.Time) {
	rp.EndedAt = &endedAt
	rp.UpdatedAt = time.Now()
}

func (rp *RoutinePeriod) IsOngoing() bool {
	return rp.EndedAt == nil
}
