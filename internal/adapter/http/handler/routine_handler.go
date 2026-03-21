package handler

import (
	"net/http"
	"strconv"

	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/dto"
	"kochappi/internal/application/service/routines"

	"github.com/gin-gonic/gin"
)

type RoutineHandler struct {
	getRoutinesUseCase        *routines.GetRoutinesUseCase
	getRoutineByIDUseCase     *routines.GetRoutineByIDUseCase
	createRoutineUseCase      *routines.CreateRoutineUseCase
	updateRoutineUseCase      *routines.UpdateRoutineUseCase
	activateRoutineUseCase    *routines.ActivateRoutineUseCase
	deactivateRoutineUseCase  *routines.DeactivateRoutineUseCase
	addRoutineDetailUseCase   *routines.AddRoutineDetailUseCase
	removeRoutineDetailUseCase *routines.RemoveRoutineDetailUseCase
	getRoutinePeriodsUseCase  *routines.GetRoutinePeriodsUseCase
}

func NewRoutineHandler(
	getRoutinesUseCase *routines.GetRoutinesUseCase,
	getRoutineByIDUseCase *routines.GetRoutineByIDUseCase,
	createRoutineUseCase *routines.CreateRoutineUseCase,
	updateRoutineUseCase *routines.UpdateRoutineUseCase,
	activateRoutineUseCase *routines.ActivateRoutineUseCase,
	deactivateRoutineUseCase *routines.DeactivateRoutineUseCase,
	addRoutineDetailUseCase *routines.AddRoutineDetailUseCase,
	removeRoutineDetailUseCase *routines.RemoveRoutineDetailUseCase,
	getRoutinePeriodsUseCase *routines.GetRoutinePeriodsUseCase,
) *RoutineHandler {
	return &RoutineHandler{
		getRoutinesUseCase:         getRoutinesUseCase,
		getRoutineByIDUseCase:      getRoutineByIDUseCase,
		createRoutineUseCase:       createRoutineUseCase,
		updateRoutineUseCase:       updateRoutineUseCase,
		activateRoutineUseCase:     activateRoutineUseCase,
		deactivateRoutineUseCase:   deactivateRoutineUseCase,
		addRoutineDetailUseCase:    addRoutineDetailUseCase,
		removeRoutineDetailUseCase: removeRoutineDetailUseCase,
		getRoutinePeriodsUseCase:   getRoutinePeriodsUseCase,
	}
}

// GetRoutines godoc
// @Summary      List routines
// @Description  Returns a list of routines, optionally filtered by customer ID
// @Tags         Routines
// @Produce      json
// @Security     BearerAuth
// @Param        customerId query int false "Customer ID"
// @Success      200 {array}  dto.RoutineResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /routines [get]
func (h *RoutineHandler) GetRoutines(c *gin.Context) {
	var customerID *int
	if cidStr := c.Query("customerId"); cidStr != "" {
		cid, err := strconv.Atoi(cidStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customerId", "code": "VALIDATION_ERROR"})
			return
		}
		customerID = &cid
	}

	response, err := h.getRoutinesUseCase.Execute(c.Request.Context(), customerID)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetRoutineByID godoc
// @Summary      Get a routine by ID
// @Description  Returns a routine with all its details
// @Tags         Routines
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Routine ID"
// @Success      200 {object} dto.RoutineWithDetailsResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /routines/{id} [get]
func (h *RoutineHandler) GetRoutineByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getRoutineByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateRoutine godoc
// @Summary      Create a new routine
// @Description  Creates a new routine, optionally including exercise details
// @Tags         Routines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateRoutineRequest true "Routine data"
// @Success      201 {object} dto.RoutineWithDetailsResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      409 {object} dto.ErrorResponse
// @Router       /routines [post]
func (h *RoutineHandler) CreateRoutine(c *gin.Context) {
	var req dto.CreateRoutineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.createRoutineUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateRoutine godoc
// @Summary      Update a routine
// @Description  Updates the name of a routine
// @Tags         Routines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Routine ID"
// @Param        request body dto.UpdateRoutineRequest true "Routine data"
// @Success      200 {object} dto.RoutineResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /routines/{id} [put]
func (h *RoutineHandler) UpdateRoutine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.UpdateRoutineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.updateRoutineUseCase.Execute(c.Request.Context(), id, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// ActivateRoutine godoc
// @Summary      Activate a routine
// @Description  Activates a routine and creates a new period
// @Tags         Routines
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Routine ID"
// @Success      200 {object} dto.RoutineResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      409 {object} dto.ErrorResponse
// @Router       /routines/{id}/activate [post]
func (h *RoutineHandler) ActivateRoutine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.activateRoutineUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeactivateRoutine godoc
// @Summary      Deactivate a routine
// @Description  Deactivates a routine and closes the ongoing period
// @Tags         Routines
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Routine ID"
// @Success      200 {object} dto.RoutineResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /routines/{id}/deactivate [post]
func (h *RoutineHandler) DeactivateRoutine(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.deactivateRoutineUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// AddRoutineDetail godoc
// @Summary      Add a detail to a routine
// @Description  Adds a new exercise detail to an existing routine
// @Tags         Routines
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Routine ID"
// @Param        request body dto.AddRoutineDetailRequest true "Detail data"
// @Success      201 {object} dto.RoutineDetailResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /routines/{id}/details [post]
func (h *RoutineHandler) AddRoutineDetail(c *gin.Context) {
	routineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.AddRoutineDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.addRoutineDetailUseCase.Execute(c.Request.Context(), routineID, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// RemoveRoutineDetail godoc
// @Summary      Remove a detail from a routine
// @Description  Removes a specific detail from a routine
// @Tags         Routines
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Routine ID"
// @Param        detailId path int true "Detail ID"
// @Success      204
// @Failure      404 {object} dto.ErrorResponse
// @Router       /routines/{id}/details/{detailId} [delete]
func (h *RoutineHandler) RemoveRoutineDetail(c *gin.Context) {
	routineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	detailID, err := strconv.Atoi(c.Param("detailId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid detailId", "code": "VALIDATION_ERROR"})
		return
	}

	if err := h.removeRoutineDetailUseCase.Execute(c.Request.Context(), routineID, detailID); err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// GetRoutinePeriods godoc
// @Summary      Get routine periods
// @Description  Returns all periods for a routine
// @Tags         Routines
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Routine ID"
// @Success      200 {array}  dto.RoutinePeriodResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /routines/{id}/periods [get]
func (h *RoutineHandler) GetRoutinePeriods(c *gin.Context) {
	routineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getRoutinePeriodsUseCase.Execute(c.Request.Context(), routineID)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
