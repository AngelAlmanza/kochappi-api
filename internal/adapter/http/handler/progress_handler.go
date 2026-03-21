package handler

import (
	"net/http"
	"strconv"

	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/dto"
	"kochappi/internal/application/service/progress"
	"kochappi/internal/domain/value_object"

	"github.com/gin-gonic/gin"
)

type ProgressHandler struct {
	getProgressLogsUseCase    *progress.GetProgressLogsUseCase
	getProgressLogByIDUseCase *progress.GetProgressLogByIDUseCase
	createProgressLogUseCase  *progress.CreateProgressLogUseCase
	deleteProgressLogUseCase  *progress.DeleteProgressLogUseCase
	uploadProgressPhotoUseCase *progress.UploadProgressPhotoUseCase
	deleteProgressPhotoUseCase *progress.DeleteProgressPhotoUseCase
}

func NewProgressHandler(
	getProgressLogsUseCase *progress.GetProgressLogsUseCase,
	getProgressLogByIDUseCase *progress.GetProgressLogByIDUseCase,
	createProgressLogUseCase *progress.CreateProgressLogUseCase,
	deleteProgressLogUseCase *progress.DeleteProgressLogUseCase,
	uploadProgressPhotoUseCase *progress.UploadProgressPhotoUseCase,
	deleteProgressPhotoUseCase *progress.DeleteProgressPhotoUseCase,
) *ProgressHandler {
	return &ProgressHandler{
		getProgressLogsUseCase:     getProgressLogsUseCase,
		getProgressLogByIDUseCase:  getProgressLogByIDUseCase,
		createProgressLogUseCase:   createProgressLogUseCase,
		deleteProgressLogUseCase:   deleteProgressLogUseCase,
		uploadProgressPhotoUseCase: uploadProgressPhotoUseCase,
		deleteProgressPhotoUseCase: deleteProgressPhotoUseCase,
	}
}

// GetProgressLogs godoc
// @Summary      List progress logs for a customer
// @Description  Returns all progress logs for a given customer
// @Tags         Progress
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Success      200 {array}  dto.ProgressLogResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /customers/{id}/log_customer_progress [get]
func (h *ProgressHandler) GetProgressLogs(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getProgressLogsUseCase.Execute(c.Request.Context(), customerID)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetProgressLogByID godoc
// @Summary      Get a progress log with photos
// @Description  Returns a single progress log with its associated photos
// @Tags         Progress
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Param        logId path int true "Progress Log ID"
// @Success      200 {object} dto.ProgressLogWithPhotosResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /customers/{id}/log_customer_progress/{logId} [get]
func (h *ProgressHandler) GetProgressLogByID(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	logID, err := strconv.Atoi(c.Param("logId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid logId", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getProgressLogByIDUseCase.Execute(c.Request.Context(), customerID, logID)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateProgressLog godoc
// @Summary      Create a progress log
// @Description  Creates a new progress log for a customer with check date and weight
// @Tags         Progress
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Param        request body dto.CreateProgressLogRequest true "Progress log data"
// @Success      201 {object} dto.ProgressLogResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /customers/{id}/log_customer_progress [post]
func (h *ProgressHandler) CreateProgressLog(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.CreateProgressLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.createProgressLogUseCase.Execute(c.Request.Context(), customerID, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// DeleteProgressLog godoc
// @Summary      Delete a progress log
// @Description  Deletes a progress log and all its associated photos and files
// @Tags         Progress
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Param        logId path int true "Progress Log ID"
// @Success      204
// @Failure      404 {object} dto.ErrorResponse
// @Router       /customers/{id}/log_customer_progress/{logId} [delete]
func (h *ProgressHandler) DeleteProgressLog(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	logID, err := strconv.Atoi(c.Param("logId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid logId", "code": "VALIDATION_ERROR"})
		return
	}

	if err := h.deleteProgressLogUseCase.Execute(c.Request.Context(), customerID, logID); err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// UploadProgressPhoto godoc
// @Summary      Upload a progress photo
// @Description  Uploads a photo for a progress log (multipart form: file + pictureType)
// @Tags         Progress
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Param        logId path int true "Progress Log ID"
// @Param        file formData file true "Photo file"
// @Param        pictureType formData string true "Picture type (front, side, back)"
// @Success      201 {object} dto.ProgressPhotoResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /customers/{id}/log_customer_progress/{logId}/photos [post]
func (h *ProgressHandler) UploadProgressPhoto(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	logID, err := strconv.Atoi(c.Param("logId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid logId", "code": "VALIDATION_ERROR"})
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required", "code": "VALIDATION_ERROR"})
		return
	}

	pictureTypeStr := c.PostForm("pictureType")
	pictureType, err := value_object.NewPictureType(pictureTypeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file", "code": "INTERNAL_ERROR"})
		return
	}
	defer file.Close()

	response, err := h.uploadProgressPhotoUseCase.Execute(c.Request.Context(), customerID, logID, pictureType, fileHeader.Filename, file)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// DeleteProgressPhoto godoc
// @Summary      Delete a progress photo
// @Description  Deletes a single photo from a progress log
// @Tags         Progress
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Param        logId path int true "Progress Log ID"
// @Param        photoId path int true "Photo ID"
// @Success      204
// @Failure      404 {object} dto.ErrorResponse
// @Router       /customers/{id}/log_customer_progress/{logId}/photos/{photoId} [delete]
func (h *ProgressHandler) DeleteProgressPhoto(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	logID, err := strconv.Atoi(c.Param("logId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid logId", "code": "VALIDATION_ERROR"})
		return
	}

	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid photoId", "code": "VALIDATION_ERROR"})
		return
	}

	if err := h.deleteProgressPhotoUseCase.Execute(c.Request.Context(), customerID, logID, photoID); err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
