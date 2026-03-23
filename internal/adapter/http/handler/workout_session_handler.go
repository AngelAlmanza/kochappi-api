package handler

import (
	"net/http"
	"strconv"
	"time"

	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/dto"
	"kochappi/internal/application/service/sessions"

	"github.com/gin-gonic/gin"
)

type WorkoutSessionHandler struct {
	getWorkoutSessionsUseCase        *sessions.GetWorkoutSessionsUseCase
	getWorkoutSessionByIDUseCase     *sessions.GetWorkoutSessionByIDUseCase
	updateWorkoutSessionStatusUseCase *sessions.UpdateWorkoutSessionStatusUseCase
	createExerciseLogUseCase         *sessions.CreateExerciseLogUseCase
	updateExerciseLogUseCase         *sessions.UpdateExerciseLogUseCase
	deleteExerciseLogUseCase         *sessions.DeleteExerciseLogUseCase
	generateDailySessionsUseCase     *sessions.GenerateDailySessionsUseCase
}

func NewWorkoutSessionHandler(
	getWorkoutSessionsUseCase *sessions.GetWorkoutSessionsUseCase,
	getWorkoutSessionByIDUseCase *sessions.GetWorkoutSessionByIDUseCase,
	updateWorkoutSessionStatusUseCase *sessions.UpdateWorkoutSessionStatusUseCase,
	createExerciseLogUseCase *sessions.CreateExerciseLogUseCase,
	updateExerciseLogUseCase *sessions.UpdateExerciseLogUseCase,
	deleteExerciseLogUseCase *sessions.DeleteExerciseLogUseCase,
	generateDailySessionsUseCase *sessions.GenerateDailySessionsUseCase,
) *WorkoutSessionHandler {
	return &WorkoutSessionHandler{
		getWorkoutSessionsUseCase:         getWorkoutSessionsUseCase,
		getWorkoutSessionByIDUseCase:      getWorkoutSessionByIDUseCase,
		updateWorkoutSessionStatusUseCase: updateWorkoutSessionStatusUseCase,
		createExerciseLogUseCase:          createExerciseLogUseCase,
		updateExerciseLogUseCase:          updateExerciseLogUseCase,
		deleteExerciseLogUseCase:          deleteExerciseLogUseCase,
		generateDailySessionsUseCase:      generateDailySessionsUseCase,
	}
}

// GetWorkoutSessions godoc
// @Summary      List workout sessions
// @Description  Returns workout sessions for a routine, optionally filtered by status or date range
// @Tags         Workout Sessions
// @Produce      json
// @Security     BearerAuth
// @Param        routineId query int true "Routine ID"
// @Param        status query string false "Filter by status (pending, in_progress, completed, skipped)"
// @Param        from query string false "Filter from date (YYYY-MM-DD)"
// @Param        to query string false "Filter to date (YYYY-MM-DD)"
// @Success      200 {array}  dto.WorkoutSessionResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /workout-sessions [get]
func (h *WorkoutSessionHandler) GetWorkoutSessions(c *gin.Context) {
	routineIDStr := c.Query("routineId")
	if routineIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "routineId is required", "code": "VALIDATION_ERROR"})
		return
	}
	routineID, err := strconv.Atoi(routineIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid routineId", "code": "VALIDATION_ERROR"})
		return
	}

	var status *string
	if s := c.Query("status"); s != "" {
		status = &s
	}

	var from, to *time.Time
	if fromStr := c.Query("from"); fromStr != "" {
		parsed, err := time.Parse(time.DateOnly, fromStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from date format, use YYYY-MM-DD", "code": "VALIDATION_ERROR"})
			return
		}
		from = &parsed
	}
	if toStr := c.Query("to"); toStr != "" {
		parsed, err := time.Parse(time.DateOnly, toStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to date format, use YYYY-MM-DD", "code": "VALIDATION_ERROR"})
			return
		}
		to = &parsed
	}

	response, err := h.getWorkoutSessionsUseCase.Execute(c.Request.Context(), routineID, status, from, to)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetWorkoutSessionByID godoc
// @Summary      Get a workout session by ID
// @Description  Returns a workout session with its exercise logs
// @Tags         Workout Sessions
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Workout Session ID"
// @Success      200 {object} dto.WorkoutSessionWithLogsResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /workout-sessions/{id} [get]
func (h *WorkoutSessionHandler) GetWorkoutSessionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getWorkoutSessionByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// UpdateWorkoutSessionStatus godoc
// @Summary      Update workout session status
// @Description  Transitions a workout session to a new status (in_progress, completed, skipped)
// @Tags         Workout Sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Workout Session ID"
// @Param        request body dto.UpdateWorkoutSessionStatusRequest true "New status"
// @Success      200 {object} dto.WorkoutSessionResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      422 {object} dto.ErrorResponse
// @Router       /workout-sessions/{id}/status [patch]
func (h *WorkoutSessionHandler) UpdateWorkoutSessionStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.UpdateWorkoutSessionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.updateWorkoutSessionStatusUseCase.Execute(c.Request.Context(), id, req.Status)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateExerciseLog godoc
// @Summary      Log an exercise set
// @Description  Creates an exercise log for an in_progress workout session
// @Tags         Workout Sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Workout Session ID"
// @Param        request body dto.CreateExerciseLogRequest true "Exercise log data"
// @Success      201 {object} dto.ExerciseLogResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      422 {object} dto.ErrorResponse
// @Router       /workout-sessions/{id}/logs [post]
func (h *WorkoutSessionHandler) CreateExerciseLog(c *gin.Context) {
	sessionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.CreateExerciseLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.createExerciseLogUseCase.Execute(c.Request.Context(), sessionID, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateExerciseLog godoc
// @Summary      Update an exercise log
// @Description  Updates an existing exercise log
// @Tags         Workout Sessions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Workout Session ID"
// @Param        logId path int true "Exercise Log ID"
// @Param        request body dto.UpdateExerciseLogRequest true "Updated exercise log data"
// @Success      200 {object} dto.ExerciseLogResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /workout-sessions/{id}/logs/{logId} [put]
func (h *WorkoutSessionHandler) UpdateExerciseLog(c *gin.Context) {
	logID, err := strconv.Atoi(c.Param("logId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid logId", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.UpdateExerciseLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.updateExerciseLogUseCase.Execute(c.Request.Context(), logID, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteExerciseLog godoc
// @Summary      Delete an exercise log
// @Description  Deletes an exercise log
// @Tags         Workout Sessions
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Workout Session ID"
// @Param        logId path int true "Exercise Log ID"
// @Success      204
// @Failure      404 {object} dto.ErrorResponse
// @Router       /workout-sessions/{id}/logs/{logId} [delete]
func (h *WorkoutSessionHandler) DeleteExerciseLog(c *gin.Context) {
	logID, err := strconv.Atoi(c.Param("logId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid logId", "code": "VALIDATION_ERROR"})
		return
	}

	if err := h.deleteExerciseLogUseCase.Execute(c.Request.Context(), logID); err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GenerateDailySessions godoc
// @Summary      Generate daily workout sessions
// @Description  Manually triggers daily session generation for all active routines
// @Tags         Workout Sessions
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} dto.GenerateDailySessionsResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /workout-sessions/generate [post]
func (h *WorkoutSessionHandler) GenerateDailySessions(c *gin.Context) {
	today := time.Now().Truncate(24 * time.Hour)

	response, err := h.generateDailySessionsUseCase.Execute(c.Request.Context(), today)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
