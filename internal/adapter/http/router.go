package http

import (
	"kochappi/internal/adapter/http/handler"
	"kochappi/internal/adapter/http/middleware"
	"kochappi/internal/application/port"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	authHandler *handler.AuthHandler,
	exerciseHandler *handler.ExerciseHandler,
	customerHandler *handler.CustomerHandler,
	templateHandler *handler.TemplateHandler,
	tokenProvider port.TokenProvider,
) *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", authHandler.Register)
			authGroup.POST("/login", authHandler.Login)
			authGroup.POST("/forgot-password", authHandler.ForgotPassword)
			authGroup.POST("/reset-password", authHandler.ResetPassword)
			authGroup.POST("/refresh", authHandler.RefreshToken)
		}

		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(tokenProvider))
		{
			protected.GET("/auth/me", authHandler.Me)

			customersGroup := protected.Group("/customers")
			{
				customersGroup.GET("", customerHandler.GetCustomers)
				customersGroup.GET("/:id", customerHandler.GetCustomerByID)
				customersGroup.POST("", customerHandler.CreateCustomer)
				customersGroup.PUT("/:id", customerHandler.UpdateCustomer)
				customersGroup.DELETE("/:id", customerHandler.DeleteCustomer)
			}

			templatesGroup := protected.Group("/templates")
			{
				templatesGroup.GET("", templateHandler.GetTemplates)
				templatesGroup.GET("/:id", templateHandler.GetTemplateByID)
				templatesGroup.POST("", templateHandler.CreateTemplate)
				templatesGroup.PUT("/:id", templateHandler.UpdateTemplate)
				templatesGroup.DELETE("/:id", templateHandler.DeleteTemplate)
				templatesGroup.POST("/:id/details", templateHandler.AddTemplateDetail)
				templatesGroup.DELETE("/:id/details/:detailId", templateHandler.RemoveTemplateDetail)
			}

			exercisesGroup := protected.Group("/exercises")
			{
				exercisesGroup.GET("", exerciseHandler.GetExercises)
				exercisesGroup.GET("/:id", exerciseHandler.GetExerciseByID)
				exercisesGroup.POST("", exerciseHandler.CreateExercise)
				exercisesGroup.PUT("/:id", exerciseHandler.UpdateExercise)
				exercisesGroup.DELETE("/:id", exerciseHandler.DeleteExercise)
			}
		}
	}

	return router
}
