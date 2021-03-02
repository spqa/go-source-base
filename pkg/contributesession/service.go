package contributesession

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"time"
)

type Service struct {
	cfg        *config.Config
	repository *repository
	enforcer   *casbin.Enforcer
}

func InitializeService(
	cfg *config.Config,
	repository *repository,
	enforcer *casbin.Enforcer,
) *Service {
	return &Service{
		cfg:        cfg,
		repository: repository,
		enforcer:   enforcer,
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

func (s Service) FindById(ctx context.Context, id int64) (*SessionRes, error) {
	entity, err := s.repository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(apperror.ErrNotFound, "contribution session not found", err)
		}
		return nil, err
	}
	return mapEntityToRes(entity), nil
}

func (s Service) Create(ctx context.Context, body *SessionCreateReq) (*SessionRes, error) {
	if err := body.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s Service) Update(ctx context.Context, id int, body *SessionUpdateReq) (*SessionRes, error) {
	return nil, nil
}

func (s Service) Delete(ctx context.Context, id int) error {
	return nil
}

func (s Service) GetCurrentSession(ctx context.Context) (*SessionRes, error) {
	entity, err := s.repository.FindSessionFromTime(ctx, time.Now())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.New(
			apperror.ErrNotFound,
			"there is not contribution session at current time", err)
	}
	if err != nil {
		return nil, err
	}
	return mapEntityToRes(entity), nil
}

func mapEntityToRes(entity *Entity) *SessionRes {
	return &SessionRes{
		Id:               entity.Id,
		OpenTime:         entity.OpenTime,
		ClosureTime:      entity.ClosureTime,
		FinalClosureTime: entity.FinalClosureTime,
		ExportedAssets:   entity.ExportedAssets,
		TrackTime: common.TrackTime{
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
	}
}

func mapEntitiesToRes(e []*Entity) []*SessionRes {
	var result []*SessionRes
	for i := range e {
		result = append(result, mapEntityToRes(e[i]))
	}
	return result
}
