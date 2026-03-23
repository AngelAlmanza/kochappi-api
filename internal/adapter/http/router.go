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
	routineHandler *handler.RoutineHandler,
	progressHandler *handler.ProgressHandler,
	workoutSessionHandler *handler.WorkoutSessionHandler,
	tokenProvider port.TokenProvider,
) *gin.Engine {
	router := gin.Default()

	router.Static("/uploads", "./uploads")

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

				progressGroup := customersGroup.Group("/:id/log_customer_progress")
				{
					progressGroup.GET("", progressHandler.GetProgressLogs)
					progressGroup.POST("", progressHandler.CreateProgressLog)
					progressGroup.GET("/:logId", progressHandler.GetProgressLogByID)
					progressGroup.DELETE("/:logId", progressHandler.DeleteProgressLog)
					progressGroup.POST("/:logId/photos", progressHandler.UploadProgressPhoto)
					progressGroup.DELETE("/:logId/photos/:photoId", progressHandler.DeleteProgressPhoto)
				}
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

			routinesGroup := protected.Group("/routines")
			{
				routinesGroup.GET("", routineHandler.GetRoutines)
				routinesGroup.GET("/:id", routineHandler.GetRoutineByID)
				routinesGroup.POST("", routineHandler.CreateRoutine)
				routinesGroup.PUT("/:id", routineHandler.UpdateRoutine)
				routinesGroup.POST("/:id/activate", routineHandler.ActivateRoutine)
				routinesGroup.POST("/:id/deactivate", routineHandler.DeactivateRoutine)
				routinesGroup.POST("/:id/details", routineHandler.AddRoutineDetail)
				routinesGroup.DELETE("/:id/details/:detailId", routineHandler.RemoveRoutineDetail)
				routinesGroup.GET("/:id/periods", routineHandler.GetRoutinePeriods)
			}

			exercisesGroup := protected.Group("/exercises")
			{
				exercisesGroup.GET("", exerciseHandler.GetExercises)
				exercisesGroup.GET("/:id", exerciseHandler.GetExerciseByID)
				exercisesGroup.POST("", exerciseHandler.CreateExercise)
				exercisesGroup.PUT("/:id", exerciseHandler.UpdateExercise)
				exercisesGroup.DELETE("/:id", exerciseHandler.DeleteExercise)
			}

			sessionsGroup := protected.Group("/workout-sessions")
			{
				sessionsGroup.GET("", workoutSessionHandler.GetWorkoutSessions)
				sessionsGroup.POST("/generate", workoutSessionHandler.GenerateDailySessions)
				sessionsGroup.GET("/:id", workoutSessionHandler.GetWorkoutSessionByID)
				sessionsGroup.PATCH("/:id/status", workoutSessionHandler.UpdateWorkoutSessionStatus)
				sessionsGroup.POST("/:id/logs", workoutSessionHandler.CreateExerciseLog)
				sessionsGroup.PUT("/:id/logs/:logId", workoutSessionHandler.UpdateExerciseLog)
				sessionsGroup.DELETE("/:id/logs/:logId", workoutSessionHandler.DeleteExerciseLog)
			}
		}
	}

	return router
}
