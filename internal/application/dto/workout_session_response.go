package dto

type WorkoutSessionResponse struct {
	ID           int    `json:"id" example:"1"`
	RoutineID    int    `json:"routineId" example:"1"`
	ScheduledDay int16  `json:"scheduledDay" example:"1"`
	ActualDate   string `json:"actualDate" example:"2026-01-15"`
	Status       string `json:"status" example:"pending"`
	StartedAt    string `json:"startedAt,omitempty" example:"2026-01-15T10:00:00Z"`
	FinishedAt   string `json:"finishedAt,omitempty" example:"2026-01-15T11:00:00Z"`
}

type ExerciseLogResponse struct {
	ID               int     `json:"id" example:"1"`
	WorkoutSessionID int     `json:"workoutSessionId" example:"1"`
	RoutineDetailID  int     `json:"routineDetailId" example:"1"`
	SetNumber        int16   `json:"setNumber" example:"1"`
	RepsDone         int16   `json:"repsDone" example:"10"`
	Weight           float64 `json:"weight" example:"50.0"`
	Notes            string  `json:"notes" example:"Felt strong"`
}

type WorkoutSessionWithLogsResponse struct {
	ID           int                   `json:"id" example:"1"`
	RoutineID    int                   `json:"routineId" example:"1"`
	ScheduledDay int16                 `json:"scheduledDay" example:"1"`
	ActualDate   string                `json:"actualDate" example:"2026-01-15"`
	Status       string                `json:"status" example:"pending"`
	StartedAt    string                `json:"startedAt,omitempty" example:"2026-01-15T10:00:00Z"`
	FinishedAt   string                `json:"finishedAt,omitempty" example:"2026-01-15T11:00:00Z"`
	ExerciseLogs []ExerciseLogResponse `json:"exerciseLogs"`
}

type GenerateDailySessionsResponse struct {
	SessionsCreated int `json:"sessionsCreated" example:"3"`
}
