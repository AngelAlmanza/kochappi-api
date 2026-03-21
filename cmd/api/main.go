package main

import (
	"fmt"
	"log"

	authAdapter "kochappi/internal/adapter/auth"
	"kochappi/internal/adapter/config"
	httpAdapter "kochappi/internal/adapter/http"
	"kochappi/internal/adapter/http/handler"
	"kochappi/internal/adapter/persistence/postgres"
	"kochappi/internal/adapter/storage"
	"kochappi/internal/application/service/auth"
	"kochappi/internal/application/service/customers"
	"kochappi/internal/application/service/exercises"
	"kochappi/internal/application/service/progress"
	"kochappi/internal/application/service/routines"
	"kochappi/internal/application/service/templates"
	"kochappi/internal/shared/logger"

	_ "kochappi/docs"
)

// @title           Kochappi API
// @version         1.0
// @description     API for the Kochappi fitness platform.

// @host            localhost:8081
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format: Bearer {token}

func main() {
	logger.Init()

	cfg := config.Load()

	db, err := postgres.NewConnection(cfg.DatabaseURL, cfg.LogLevel)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := postgres.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run auto-migration: %v", err)
	}

	logger.Info.Println("Database connected and migrated successfully")

	// Adapters
	passwordHasher := authAdapter.NewBcryptPasswordHasher()
	jwtProvider := authAdapter.NewJWTProvider(cfg.JWTSecret, cfg.JWTAccessExpiryMin, cfg.JWTRefreshExpiryDay)
	otpService := authAdapter.NewConsoleOTPService()
	fileStorage := storage.NewLocalFileStorage("./uploads", "/uploads")

	// Repositories
	userRepo := postgres.NewPostgresUserRepository(db)
	refreshTokenRepo := postgres.NewPostgresRefreshTokenRepository(db)
	exerciseRepo := postgres.NewPostgresExerciseRepository(db)
	customerRepo := postgres.NewCustomerRepository(db)
	templateRepo := postgres.NewPostgresTemplateRepository(db)
	templateDetailRepo := postgres.NewPostgresTemplateDetailRepository(db)
	routineRepo := postgres.NewPostgresRoutineRepository(db)
	routineDetailRepo := postgres.NewPostgresRoutineDetailRepository(db)
	routinePeriodRepo := postgres.NewPostgresRoutinePeriodRepository(db)
	logProgressRepo := postgres.NewLogCustomerProgressRepository(db)
	progressPhotoRepo := postgres.NewProgressPhotoRepository(db)

	// Use Cases
	registerUseCase := auth.NewRegisterUseCase(userRepo, passwordHasher, jwtProvider, refreshTokenRepo)
	loginUseCase := auth.NewLoginUseCase(userRepo, passwordHasher, jwtProvider, refreshTokenRepo)
	forgotPasswordUseCase := auth.NewForgotPasswordUseCase(userRepo, otpService, cfg.OTPExpiryMinutes)
	resetPasswordUseCase := auth.NewResetPasswordUseCase(userRepo, passwordHasher, refreshTokenRepo)
	refreshTokenUseCase := auth.NewRefreshTokenUseCase(userRepo, jwtProvider, refreshTokenRepo)

	getExercisesUseCase := exercises.NewGetExercisesUseCase(exerciseRepo)
	getExerciseByIDUseCase := exercises.NewGetExerciseByIDUseCase(exerciseRepo)
	createExerciseUseCase := exercises.NewCreateExerciseUseCase(exerciseRepo)
	updateExerciseUseCase := exercises.NewUpdateExerciseUseCase(exerciseRepo)
	deleteExerciseUseCase := exercises.NewDeleteExerciseUseCase(exerciseRepo)

	getTemplatesUseCase := templates.NewGetTemplatesUseCase(templateRepo)
	getTemplateByIDUseCase := templates.NewGetTemplateByIDUseCase(templateRepo, templateDetailRepo)
	createTemplateUseCase := templates.NewCreateTemplateUseCase(templateRepo, templateDetailRepo, exerciseRepo)
	updateTemplateUseCase := templates.NewUpdateTemplateUseCase(templateRepo)
	deleteTemplateUseCase := templates.NewDeleteTemplateUseCase(templateRepo)
	addTemplateDetailUseCase := templates.NewAddTemplateDetailUseCase(templateRepo, templateDetailRepo, exerciseRepo)
	removeTemplateDetailUseCase := templates.NewRemoveTemplateDetailUseCase(templateDetailRepo)

	getRoutinesUseCase := routines.NewGetRoutinesUseCase(routineRepo)
	getRoutineByIDUseCase := routines.NewGetRoutineByIDUseCase(routineRepo, routineDetailRepo)
	createRoutineUseCase := routines.NewCreateRoutineUseCase(routineRepo, routineDetailRepo, routinePeriodRepo, customerRepo, templateRepo, exerciseRepo)
	updateRoutineUseCase := routines.NewUpdateRoutineUseCase(routineRepo)
	activateRoutineUseCase := routines.NewActivateRoutineUseCase(routineRepo, routinePeriodRepo)
	deactivateRoutineUseCase := routines.NewDeactivateRoutineUseCase(routineRepo, routinePeriodRepo)
	addRoutineDetailUseCase := routines.NewAddRoutineDetailUseCase(routineRepo, routineDetailRepo, exerciseRepo)
	removeRoutineDetailUseCase := routines.NewRemoveRoutineDetailUseCase(routineDetailRepo)
	getRoutinePeriodsUseCase := routines.NewGetRoutinePeriodsUseCase(routineRepo, routinePeriodRepo)

	getCustomersUseCase := customers.NewGetCustomersUseCase(customerRepo)
	getCustomerByIDUseCase := customers.NewGetCustomerByIDUseCase(customerRepo)
	createCustomerUseCase := customers.NewCreateCustomerUseCase(customerRepo, userRepo)
	updateCustomerUseCase := customers.NewUpdateCustomerUseCase(customerRepo)
	deleteCustomerUseCase := customers.NewDeleteCustomerUseCase(customerRepo)

	getProgressLogsUseCase := progress.NewGetProgressLogsUseCase(customerRepo, logProgressRepo)
	getProgressLogByIDUseCase := progress.NewGetProgressLogByIDUseCase(customerRepo, logProgressRepo, progressPhotoRepo)
	createProgressLogUseCase := progress.NewCreateProgressLogUseCase(customerRepo, logProgressRepo)
	deleteProgressLogUseCase := progress.NewDeleteProgressLogUseCase(customerRepo, logProgressRepo, progressPhotoRepo, fileStorage)
	uploadProgressPhotoUseCase := progress.NewUploadProgressPhotoUseCase(customerRepo, logProgressRepo, progressPhotoRepo, fileStorage)
	deleteProgressPhotoUseCase := progress.NewDeleteProgressPhotoUseCase(customerRepo, logProgressRepo, progressPhotoRepo, fileStorage)

	// Handlers
	authHandler := handler.NewAuthHandler(
		registerUseCase,
		loginUseCase,
		forgotPasswordUseCase,
		resetPasswordUseCase,
		refreshTokenUseCase,
	)
	exerciseHandler := handler.NewExerciseHandler(
		getExercisesUseCase,
		getExerciseByIDUseCase,
		createExerciseUseCase,
		updateExerciseUseCase,
		deleteExerciseUseCase,
	)
	customerHandler := handler.NewCustomerHandler(
		getCustomersUseCase,
		getCustomerByIDUseCase,
		createCustomerUseCase,
		updateCustomerUseCase,
		deleteCustomerUseCase,
	)
	templateHandler := handler.NewTemplateHandler(
		getTemplatesUseCase,
		getTemplateByIDUseCase,
		createTemplateUseCase,
		updateTemplateUseCase,
		deleteTemplateUseCase,
		addTemplateDetailUseCase,
		removeTemplateDetailUseCase,
	)

	routineHandler := handler.NewRoutineHandler(
		getRoutinesUseCase,
		getRoutineByIDUseCase,
		createRoutineUseCase,
		updateRoutineUseCase,
		activateRoutineUseCase,
		deactivateRoutineUseCase,
		addRoutineDetailUseCase,
		removeRoutineDetailUseCase,
		getRoutinePeriodsUseCase,
	)

	progressHandler := handler.NewProgressHandler(
		getProgressLogsUseCase,
		getProgressLogByIDUseCase,
		createProgressLogUseCase,
		deleteProgressLogUseCase,
		uploadProgressPhotoUseCase,
		deleteProgressPhotoUseCase,
	)

	// Router
	router := httpAdapter.NewRouter(authHandler, exerciseHandler, customerHandler, templateHandler, routineHandler, progressHandler, jwtProvider)

	addr := fmt.Sprintf(":%s", cfg.Port)
	logger.Info.Printf("Server starting on %s (env: %s)", addr, cfg.Env)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
