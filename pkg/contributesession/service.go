package contributesession

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"mcm-api/pkg/log"
	"time"
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

func (s Service) FindById(ctx context.Context, id int) (*SessionRes, error) {
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
	lastSession, err := s.repository.GetLastSession(ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if body.OpenTime.Before(lastSession.FinalClosureTime) || body.OpenTime.Equal(lastSession.FinalClosureTime) {
			return nil, apperror.New(apperror.ErrConflict, "conflict with last contribute session", nil)
		}
	}
	entity, err := s.repository.Create(ctx, &Entity{
		OpenTime:         body.OpenTime,
		ClosureTime:      body.ClosureTime,
		FinalClosureTime: body.FinalClosureTime,
	})
	if err != nil {
		return nil, err
	}
	return mapEntityToRes(entity), nil
}

func (s Service) Update(ctx context.Context, id int, body *SessionUpdateReq) (*SessionRes, error) {
	entity, err := s.repository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(apperror.ErrNotFound, "contribute session not found", err)
		}
		return nil, err
	}
	lastSession, err := s.repository.GetLastSession(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == nil {
		if body.OpenTime.Before(lastSession.FinalClosureTime) || body.OpenTime.Equal(lastSession.FinalClosureTime) {
			return nil, apperror.New(apperror.ErrConflict, "conflict with last contribute session", nil)
		}
	}
	entity.OpenTime = body.OpenTime
	entity.ClosureTime = body.ClosureTime
	entity.FinalClosureTime = body.FinalClosureTime
	entity, err = s.repository.Update(ctx, entity)
	if err != nil {
		return nil, err
	}
	return mapEntityToRes(entity), nil
}

func (s Service) Delete(ctx context.Context, id int) error {
	hasContribution, err := s.repository.HasContribution(ctx, id)
	if err != nil {
		return err
	}
	if hasContribution {
		return apperror.New(apperror.ErrInvalid, "cant delete session that has contribution", nil)
	}
	err = s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}
	log.Logger.Debug("deleted contribution session", zap.Int("id", id))
	return nil
}

func (s Service) GetCurrentSession(ctx context.Context) (*SessionRes, error) {
	entity, err := s.repository.FindSessionFromTime(ctx, time.Now())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.New(
			apperror.ErrNotFound,
			"there is no contribution session at current time", err)
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
