package sessions

import (
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/domain/entity"
)

func mapWorkoutSessionToResponse(s *entity.WorkoutSession) dto.WorkoutSessionResponse {
	resp := dto.WorkoutSessionResponse{
		ID:           s.ID,
		RoutineID:    s.RoutineID,
		ScheduledDay: s.ScheduledDay,
		ActualDate:   s.ActualDate.Format(time.DateOnly),
		Status:       string(s.Status),
	}
	if s.StartedAt != nil {
		resp.StartedAt = s.StartedAt.Format(time.RFC3339)
	}
	if s.FinishedAt != nil {
		resp.FinishedAt = s.FinishedAt.Format(time.RFC3339)
	}
	return resp
}

func mapExerciseLogToResponse(l *entity.LogExerciseSession) dto.ExerciseLogResponse {
	return dto.ExerciseLogResponse{
		ID:               l.ID,
		WorkoutSessionID: l.WorkoutSessionID,
		RoutineDetailID:  l.RoutineDetailID,
		SetNumber:        l.SetNumber,
		RepsDone:         l.RepsDone,
		Weight:           l.Weight,
		Notes:            l.Notes,
	}
}
