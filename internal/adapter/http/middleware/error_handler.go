package middleware

import (
	"errors"
	"net/http"

	domainerror "kochappi/internal/domain/error"
	"kochappi/internal/domain/value_object"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var domainErr domainerror.DomainError
	if errors.As(err, &domainErr) {
		statusCode := mapDomainErrorToHTTPStatus(domainErr)
		c.JSON(statusCode, gin.H{
			"error": domainErr.Error(),
			"code":  domainErr.Code(),
		})
		return
	}

	var validationErr *value_object.ValidationError
	if errors.As(err, &validationErr) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErr.Error(),
			"code":  "VALIDATION_ERROR",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
		"code":  "INTERNAL_ERROR",
	})
}

func mapDomainErrorToHTTPStatus(err domainerror.DomainError) int {
	switch err.Code() {
	case "USER_NOT_FOUND":
		return http.StatusNotFound
	case "EMAIL_ALREADY_EXISTS":
		return http.StatusConflict
	case "INVALID_CREDENTIALS":
		return http.StatusUnauthorized
	case "INVALID_OTP":
		return http.StatusBadRequest
	case "INVALID_TOKEN":
		return http.StatusUnauthorized
	case "UNAUTHORIZED":
		return http.StatusForbidden
	case "EXERCISE_NOT_FOUND":
		return http.StatusNotFound
	case "CUSTOMER_NOT_FOUND":
		return http.StatusNotFound
	case "CUSTOMER_ALREADY_EXISTS":
		return http.StatusConflict
	case "USER_NOT_CUSTOMER":
		return http.StatusUnprocessableEntity
	case "INVALID_BIRTHDATE":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
