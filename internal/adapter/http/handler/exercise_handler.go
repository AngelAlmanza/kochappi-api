package handler

import (
	"net/http"
	"strconv"

	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/dto"
	"kochappi/internal/application/service/exercises"

	"github.com/gin-gonic/gin"
)

type ExerciseHandler struct {
	getExercisesUseCase    *exercises.GetExercisesUseCase
	getExerciseByIDUseCase *exercises.GetExerciseByIDUseCase
	createExerciseUseCase  *exercises.CreateExerciseUseCase
	updateExerciseUseCase  *exercises.UpdateExerciseUseCase
	deleteExerciseUseCase  *exercises.DeleteExerciseUseCase
}

func NewExerciseHandler(
	getExercisesUseCase *exercises.GetExercisesUseCase,
	getExerciseByIDUseCase *exercises.GetExerciseByIDUseCase,
	createExerciseUseCase *exercises.CreateExerciseUseCase,
	updateExerciseUseCase *exercises.UpdateExerciseUseCase,
	deleteExerciseUseCase *exercises.DeleteExerciseUseCase,
) *ExerciseHandler {
	return &ExerciseHandler{
		getExercisesUseCase:    getExercisesUseCase,
		getExerciseByIDUseCase: getExerciseByIDUseCase,
		createExerciseUseCase:  createExerciseUseCase,
		updateExerciseUseCase:  updateExerciseUseCase,
		deleteExerciseUseCase:  deleteExerciseUseCase,
	}
}

// GetExercises godoc
// @Summary      List all exercises
// @Description  Returns a list of all exercises
// @Tags         Exercises
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array}  dto.ExerciseResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /exercises [get]
func (h *ExerciseHandler) GetExercises(c *gin.Context) {
	response, err := h.getExercisesUseCase.Execute(c.Request.Context())
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetExerciseByID godoc
// @Summary      Get an exercise by ID
// @Description  Returns a single exercise by its ID
// @Tags         Exercises
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Exercise ID"
// @Success      200 {object} dto.ExerciseResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /exercises/{id} [get]
func (h *ExerciseHandler) GetExerciseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getExerciseByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateExercise godoc
// @Summary      Create a new exercise
// @Description  Creates a new exercise and returns it
// @Tags         Exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateExerciseRequest true "Exercise data"
// @Success      201 {object} dto.ExerciseResponse
// @Failure      400 {object} dto.ErrorResponse
// @Router       /exercises [post]
func (h *ExerciseHandler) CreateExercise(c *gin.Context) {
	var req dto.CreateExerciseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.createExerciseUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateExercise godoc
// @Summary      Update an exercise
// @Description  Updates an existing exercise by its ID
// @Tags         Exercises
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Exercise ID"
// @Param        request body dto.UpdateExerciseRequest true "Exercise data"
// @Success      200 {object} dto.ExerciseResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /exercises/{id} [put]
func (h *ExerciseHandler) UpdateExercise(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.UpdateExerciseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.updateExerciseUseCase.Execute(c.Request.Context(), id, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteExercise godoc
// @Summary      Delete an exercise
// @Description  Deletes an exercise by its ID
// @Tags         Exercises
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Exercise ID"
// @Success      204
// @Failure      404 {object} dto.ErrorResponse
// @Router       /exercises/{id} [delete]
func (h *ExerciseHandler) DeleteExercise(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	if err := h.deleteExerciseUseCase.Execute(c.Request.Context(), id); err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
