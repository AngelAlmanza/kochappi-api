package routines

import (
	"context"
	"time"

	"kochappi/internal/application/dto"
	"kochappi/internal/application/port"
	"kochappi/internal/domain/entity"
	domainerror "kochappi/internal/domain/error"
)

type CreateRoutineUseCase struct {
	routineRepo       port.RoutineRepository
	routineDetailRepo port.RoutineDetailRepository
	routinePeriodRepo port.RoutinePeriodRepository
	customerRepo      port.CustomerRepository
	templateRepo      port.TemplateRepository
	exerciseRepo      port.ExerciseRepository
}

func NewCreateRoutineUseCase(
	routineRepo port.RoutineRepository,
	routineDetailRepo port.RoutineDetailRepository,
	routinePeriodRepo port.RoutinePeriodRepository,
	customerRepo port.CustomerRepository,
	templateRepo port.TemplateRepository,
	exerciseRepo port.ExerciseRepository,
) *CreateRoutineUseCase {
	return &CreateRoutineUseCase{
		routineRepo:       routineRepo,
		routineDetailRepo: routineDetailRepo,
		routinePeriodRepo: routinePeriodRepo,
		customerRepo:      customerRepo,
		templateRepo:      templateRepo,
		exerciseRepo:      exerciseRepo,
	}
}

func (uc *CreateRoutineUseCase) Execute(ctx context.Context, req *dto.CreateRoutineRequest) (*dto.RoutineWithDetailsResponse, error) {
	if _, err := uc.customerRepo.GetByID(ctx, req.CustomerID); err != nil {
		return nil, err
	}

	if req.TemplateID != nil {
		if _, err := uc.templateRepo.GetByID(ctx, *req.TemplateID); err != nil {
			return nil, err
		}
	}

	if req.IsActive {
		active, err := uc.routineRepo.GetActiveByCustomerID(ctx, req.CustomerID)
		if err != nil {
			return nil, err
		}
		if active != nil {
			return nil, &domainerror.ActiveRoutineExistsError{CustomerID: req.CustomerID}
		}
	}

	if len(req.Details) > 0 {
		seen := make(map[int]struct{}, len(req.Details))
		ids := make([]int, 0, len(req.Details))
		for _, d := range req.Details {
			if _, ok := seen[d.ExerciseID]; !ok {
				seen[d.ExerciseID] = struct{}{}
				ids = append(ids, d.ExerciseID)
			}
		}
		if _, err := uc.exerciseRepo.GetByIDs(ctx, ids); err != nil {
			return nil, err
		}
	}

	routine := entity.NewRoutine(req.CustomerID, req.TemplateID, req.Name)
	if req.IsActive {
		routine.Activate()
	}

	if err := uc.routineRepo.Create(ctx, routine); err != nil {
		return nil, err
	}

	detailEntities := make([]*entity.RoutineDetail, 0, len(req.Details))
	for _, d := range req.Details {
		detailEntities = append(detailEntities, entity.NewRoutineDetail(routine.ID, d.DayOfWeek, d.ExerciseID, d.Sets, d.Reps, d.DisplayOrder))
	}

	if len(detailEntities) > 0 {
		if err := uc.routineDetailRepo.CreateBulk(ctx, detailEntities); err != nil {
			return nil, err
		}
	}

	if req.IsActive {
		period := entity.NewRoutinePeriod(routine.ID, time.Now())
		if err := uc.routinePeriodRepo.Create(ctx, period); err != nil {
			return nil, err
		}
	}

	detailResponses := make([]dto.RoutineDetailResponse, 0, len(detailEntities))
	for _, d := range detailEntities {
		detailResponses = append(detailResponses, dto.RoutineDetailResponse{
			ID:           d.ID,
			DayOfWeek:    d.DayOfWeek,
			ExerciseID:   d.ExerciseID,
			Sets:         d.Sets,
			Reps:         d.Reps,
			DisplayOrder: d.DisplayOrder,
		})
	}

	return &dto.RoutineWithDetailsResponse{
		ID:         routine.ID,
		CustomerID: routine.CustomerID,
		TemplateID: routine.TemplateID,
		Name:       routine.Name,
		IsActive:   routine.IsActive,
		Details:    detailResponses,
	}, nil
}
