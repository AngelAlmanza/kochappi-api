package handler

import (
	"net/http"

	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/dto"
	"kochappi/internal/application/service/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	registerUseCase       *auth.RegisterUseCase
	loginUseCase          *auth.LoginUseCase
	forgotPasswordUseCase *auth.ForgotPasswordUseCase
	resetPasswordUseCase  *auth.ResetPasswordUseCase
	refreshTokenUseCase   *auth.RefreshTokenUseCase
}

func NewAuthHandler(
	registerUseCase *auth.RegisterUseCase,
	loginUseCase *auth.LoginUseCase,
	forgotPasswordUseCase *auth.ForgotPasswordUseCase,
	resetPasswordUseCase *auth.ResetPasswordUseCase,
	refreshTokenUseCase *auth.RefreshTokenUseCase,
) *AuthHandler {
	return &AuthHandler{
		registerUseCase:       registerUseCase,
		loginUseCase:          loginUseCase,
		forgotPasswordUseCase: forgotPasswordUseCase,
		resetPasswordUseCase:  resetPasswordUseCase,
		refreshTokenUseCase:   refreshTokenUseCase,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Creates a new user account and returns authentication tokens
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterRequest true "Registration data"
// @Success      201 {object} dto.AuthResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      409 {object} dto.ErrorResponse
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.registerUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Login godoc
// @Summary      Login
// @Description  Authenticates a user and returns access and refresh tokens
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login credentials"
// @Success      200 {object} dto.AuthResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.loginUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// ForgotPassword godoc
// @Summary      Request password reset
// @Description  Sends a password reset OTP code to the user's email
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.ForgotPasswordRequest true "Email address"
// @Success      200 {object} dto.MessageResponse
// @Failure      400 {object} dto.ErrorResponse
// @Router       /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req dto.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.forgotPasswordUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// ResetPassword godoc
// @Summary      Reset password
// @Description  Resets the user's password using the OTP code received via email
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.ResetPasswordRequest true "Reset password data"
// @Success      200 {object} dto.MessageResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /auth/reset-password [post]
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.resetPasswordUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// RefreshToken godoc
// @Summary      Refresh tokens
// @Description  Generates a new pair of access and refresh tokens using a valid refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RefreshTokenRequest true "Refresh token"
// @Success      200 {object} dto.TokenResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Router       /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.refreshTokenUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Me godoc
// @Summary      Get current user info
// @Description  Returns the authenticated user's ID and role
// @Tags         Auth
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} dto.MeResponse
// @Failure      401 {object} dto.ErrorResponse
// @Router       /auth/me [get]
func (h *AuthHandler) Me(c *gin.Context) {
	userID := middleware.GetUserIDFromContext(c)
	role := middleware.GetRoleFromContext(c)

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"role":    role,
	})
}
