package faculty

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
)

type Service struct {
	cfg        *config.Config
	repository *repository
}

func InitializeService(
	cfg *config.Config,
	repository *repository,
) *Service {
	return &Service{
		cfg:        cfg,
		repository: repository,
	}
}

func (s Service) Find(ctx context.Context, query *IndexQuery) (*common.PaginateResponse, error) {
	entities, count, err := s.repository.FindAndCount(ctx, query)
	if err != nil {
		return nil, err
	}
	res := mapEntitiesToRes(entities)
	return common.NewPaginateResponse(res, count, query.Page, query.GetLimit()), nil
}

func (s Service) FindById(ctx context.Context, id int) (*FacultyResponse, error) {
	entity, err := s.repository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(apperror.ErrNotFound, "faculty not found", err)
		}
		return nil, err
	}
	return mapEntityToRes(entity), nil
}

func (s Service) Create(ctx context.Context, body *FacultyCreateReq) (*FacultyResponse, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}
	entity, err := s.repository.Create(ctx, &Entity{
		Name: body.Name,
	})
	if err != nil {
		return nil, err
	}
	return mapEntityToRes(entity), nil
}

func (s Service) Update(ctx context.Context, id int, body *FacultyUpdateReq) (*FacultyResponse, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}
	entity, err := s.repository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(apperror.ErrNotFound, "faculty not found", err)
		}
		return nil, err
	}
	entity.Name = body.Name
	entity, err = s.repository.Update(ctx, entity)
	if err != nil {
		return nil, err
	}
	return mapEntityToRes(entity), nil
}

func (s Service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func mapEntityToRes(entity *Entity) *FacultyResponse {
	return &FacultyResponse{
		Id:   entity.Id,
		Name: entity.Name,
		TrackTime: common.TrackTime{
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
	}
}

func mapEntitiesToRes(entities []*Entity) []*FacultyResponse {
	var result []*FacultyResponse
	for i := range entities {
		result = append(result, mapEntityToRes(entities[i]))
	}
	return result
}
