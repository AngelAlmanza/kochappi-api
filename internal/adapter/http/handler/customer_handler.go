package handler

import (
	"net/http"
	"strconv"

	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/dto"
	"kochappi/internal/application/service/customers"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	getCustomersUseCase    *customers.GetCustomersUseCase
	getCustomerByIDUseCase *customers.GetCustomerByIDUseCase
	createCustomerUseCase  *customers.CreateCustomerUseCase
	updateCustomerUseCase  *customers.UpdateCustomerUseCase
	deleteCustomerUseCase  *customers.DeleteCustomerUseCase
}

func NewCustomerHandler(
	getCustomersUseCase *customers.GetCustomersUseCase,
	getCustomerByIDUseCase *customers.GetCustomerByIDUseCase,
	createCustomerUseCase *customers.CreateCustomerUseCase,
	updateCustomerUseCase *customers.UpdateCustomerUseCase,
	deleteCustomerUseCase *customers.DeleteCustomerUseCase,
) *CustomerHandler {
	return &CustomerHandler{
		getCustomersUseCase:    getCustomersUseCase,
		getCustomerByIDUseCase: getCustomerByIDUseCase,
		createCustomerUseCase:  createCustomerUseCase,
		updateCustomerUseCase:  updateCustomerUseCase,
		deleteCustomerUseCase:  deleteCustomerUseCase,
	}
}

// GetCustomers godoc
// @Summary      List all customers
// @Description  Returns a list of all customers
// @Tags         Customers
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array}  dto.CustomerResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /customers [get]
func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	response, err := h.getCustomersUseCase.Execute(c.Request.Context())
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetCustomerByID godoc
// @Summary      Get a customer by ID
// @Description  Returns a single customer by its ID
// @Tags         Customers
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Success      200 {object} dto.CustomerResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /customers/{id} [get]
func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.getCustomerByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateCustomer godoc
// @Summary      Create a new customer
// @Description  Creates a new customer linked to an existing user with the client role
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateCustomerRequest true "Customer data"
// @Success      201 {object} dto.CustomerResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Failure      409 {object} dto.ErrorResponse
// @Router       /customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.createCustomerUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateCustomer godoc
// @Summary      Update a customer
// @Description  Updates an existing customer's name and birthdate
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Param        request body dto.UpdateCustomerRequest true "Customer data"
// @Success      200 {object} dto.CustomerResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      404 {object} dto.ErrorResponse
// @Router       /customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	var req dto.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": "VALIDATION_ERROR"})
		return
	}

	response, err := h.updateCustomerUseCase.Execute(c.Request.Context(), id, &req)
	if err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteCustomer godoc
// @Summary      Delete a customer
// @Description  Deletes a customer record (does not delete the associated user)
// @Tags         Customers
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "Customer ID"
// @Success      204
// @Failure      404 {object} dto.ErrorResponse
// @Router       /customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id", "code": "VALIDATION_ERROR"})
		return
	}

	if err := h.deleteCustomerUseCase.Execute(c.Request.Context(), id); err != nil {
		middleware.HandleError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
