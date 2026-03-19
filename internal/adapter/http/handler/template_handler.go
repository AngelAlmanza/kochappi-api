package handler

import (
	"net/http"
	"strconv"

	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/dto"
	"kochappi/internal/application/service/templates"

	"github.com/gin-gonic/gin"
)

type TemplateHandler struct {
	getTemplatesUseCase        *templates.GetTemplatesUseCase
	getTemplateByIDUseCase     *templates.GetTemplateByIDUseCase
	createTemplateUseCase      *templates.CreateTemplateUseCase
	updateTemplateUseCase      *templates.UpdateTemplateUseCase
	deleteTemplateUseCase      *templates.DeleteTemplateUseCase
	addTemplateDetailUseCase   *templates.AddTemplateDetailUseCase
	removeTemplateDetailUseCase *templates.RemoveTemplateDetailUseCase
}

func NewTemplateHandler(
	getTemplatesUseCase *templates.GetTemplatesUseCase,
	getTemplateByIDUseCase *templates.GetTemplateByIDUseCase,
	createTemplateUseCase *templates.CreateTemplateUseCase,
	updateTemplateUseCase *templates.UpdateTemplateUseCase,
	deleteTemplateUseCase *templates.DeleteTemplateUseCase,
	addTemplateDetailUseCase *templates.AddTemplateDetailUseCase,
	removeTemplateDetailUseCase *templates.RemoveTemplateDetailUseCase,
) *TemplateHandler {
	return &TemplateHandler{
		getTemplatesUseCase:         getTemplatesUseCase,
		getTemplateByIDUseCase:      getTemplateByIDUseCase,
		createTemplateUseCase:       createTemplateUseCase,
		updateTemplateUseCase:       updateTemplateUseCase,
		deleteTemplateUseCase:       deleteTemplateUseCase,
		addTemplateDetailUseCase:    addTemplateDetailUseCase,
		removeTemplateDetailUseCase: removeTemplateDetailUseCase,
	}
}

// GetTemplates godoc
// @Summary      List all templates
// @Description  Returns a list of all templates (without details)
// @Tags         Templates
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array}  dto.TemplateResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /templates [get]
func (h *TemplateHandler) GetTemplates(c *gin.Context) {
	response, err := h.getTemplatesUseCase.Execute(c.Request.Context())
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetTemplateByID godoc
// @Summary      Get a template by ID
// @Description  Returns a template with all its details
// @Tags         Templates
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Template ID"
// @Success      200 {object} dto.TemplateWithDetailsResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /templates/{id} [get]
func (h *TemplateHandler) GetTemplateByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getTemplateByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateTemplate godoc
// @Summary      Create a new template
// @Description  Creates a new template, optionally including exercise details
// @Tags         Templates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateTemplateRequest true "Template data"
// @Success      201 {object} dto.TemplateWithDetailsResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /templates [post]
func (h *TemplateHandler) CreateTemplate(c *gin.Context) {
	var req dto.CreateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.createTemplateUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateTemplate godoc
// @Summary      Update a template
// @Description  Updates the name and description of a template (does not affect details)
// @Tags         Templates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Template ID"
// @Param        request body dto.UpdateTemplateRequest true "Template data"
// @Success      200 {object} dto.TemplateResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /templates/{id} [put]
func (h *TemplateHandler) UpdateTemplate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.UpdateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.updateTemplateUseCase.Execute(c.Request.Context(), id, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteTemplate godoc
// @Summary      Delete a template
// @Description  Deletes a template. Associated details are cascade-deleted and routines referencing this template will have template_id set to NULL
// @Tags         Templates
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Template ID"
// @Success      204
// @Failure      404 {object} dto.ErrorResponse
// @Router       /templates/{id} [delete]
func (h *TemplateHandler) DeleteTemplate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	if err := h.deleteTemplateUseCase.Execute(c.Request.Context(), id); err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// AddTemplateDetail godoc
// @Summary      Add a detail to a template
// @Description  Adds a new exercise detail to an existing template
// @Tags         Templates
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Template ID"
// @Param        request body dto.AddTemplateDetailRequest true "Detail data"
// @Success      201 {object} dto.TemplateDetailResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /templates/{id}/details [post]
func (h *TemplateHandler) AddTemplateDetail(c *gin.Context) {
	templateID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.AddTemplateDetailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.addTemplateDetailUseCase.Execute(c.Request.Context(), templateID, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// RemoveTemplateDetail godoc
// @Summary      Remove a detail from a template
// @Description  Removes a specific detail from a template
// @Tags         Templates
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Template ID"
// @Param        detailId path int true "Detail ID"
// @Success      204
// @Failure      404 {object} dto.ErrorResponse
// @Router       /templates/{id}/details/{detailId} [delete]
func (h *TemplateHandler) RemoveTemplateDetail(c *gin.Context) {
	templateID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	detailID, err := strconv.Atoi(c.Param("detailId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid detailId", "code": "VALIDATION_ERROR"})
		return
	}

	if err := h.removeTemplateDetailUseCase.Execute(c.Request.Context(), templateID, detailID); err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
