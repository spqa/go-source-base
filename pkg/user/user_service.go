package user

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/user/responses"
)

type Service struct {
	repository *repository
	enforcer   *casbin.Enforcer
}

func InitializeService(repository *repository, enforcer *casbin.Enforcer) *Service {
	return &Service{
		repository: repository,
		enforcer:   enforcer,
	}
}

func (s Service) FindById(ctx context.Context, id int) (*responses.UserResponse, error) {
	entity, err := s.repository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(apperror.ErrNotFound, "user not found", err)
		}
		return nil, err
	}
	return &responses.UserResponse{
		Id:   entity.Id,
		Name: entity.Name,
	}, nil
}
