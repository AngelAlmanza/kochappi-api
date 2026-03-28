package handler

import (
	"net/http"
	"strconv"

	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/dto"
	"kochappi/internal/application/service/users"
	"kochappi/internal/domain/entity"
	"kochappi/internal/shared/logger"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	getUsersUseCase        *users.GetUsersUseCase
	getUserByIDUseCase     *users.GetUserByIDUseCase
	createUserUseCase      *users.CreateUserUseCase
	updateUserEmailUseCase *users.UpdateUserEmailUseCase
}

func NewUserHandler(
	getUsersUseCase *users.GetUsersUseCase,
	getUserByIDUseCase *users.GetUserByIDUseCase,
	createUserUseCase *users.CreateUserUseCase,
	updateUserEmailUseCase *users.UpdateUserEmailUseCase,
) *UserHandler {
	return &UserHandler{
		getUsersUseCase:        getUsersUseCase,
		getUserByIDUseCase:     getUserByIDUseCase,
		createUserUseCase:      createUserUseCase,
		updateUserEmailUseCase: updateUserEmailUseCase,
	}
}

// GetUsers godoc
// @Summary      List users
// @Description  Returns a list of users, optionally filtered by role
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        role query string false "Role filter" Enums(trainer, client)
// @Success      200 {array}  dto.UserDetailResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	var role *entity.Role

	if roleStr := c.Query("role"); roleStr != "" {
		if roleStr != string(entity.ROLE_TRAINER) && roleStr != string(entity.ROLE_CLIENT) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role must be 'trainer' or 'client'", "code": "VALIDATION_ERROR"})
			return
		}
		r := entity.Role(roleStr)
		role = &r
	}

	includeWithCustomers := true
	if includeWithCustomersStr := c.Query("includeWithCustomers"); includeWithCustomersStr != "" {
		var err error
		includeWithCustomers, err = strconv.ParseBool(includeWithCustomersStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "includeWithCustomers must be a boolean", "code": "VALIDATION_ERROR"})
			return
		}
	}

	logger.Info.Printf("Getting with params: role=%s, includeWithCustomers=%t", *role, includeWithCustomers)

	response, err := h.getUsersUseCase.Execute(c.Request.Context(), role, includeWithCustomers)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetUserByID godoc
// @Summary      Get a user by ID
// @Description  Returns a single user by their ID
// @Tags         Users
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "User ID"
// @Success      200 {object} dto.UserDetailResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getUserByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Creates a new user with the specified role
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateUserRequest true "User data"
// @Success      201 {object} dto.UserDetailResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      409 {object} dto.ErrorResponse
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.createUserUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateUserEmail godoc
// @Summary      Update user email
// @Description  Updates the email address of a user, checking for uniqueness
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "User ID"
// @Param        request body dto.UpdateUserEmailRequest true "New email"
// @Success      200 {object} dto.UserDetailResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      409 {object} dto.ErrorResponse
// @Router       /users/{id}/email [patch]
func (h *UserHandler) UpdateUserEmail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.UpdateUserEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.updateUserEmailUseCase.Execute(c.Request.Context(), id, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
