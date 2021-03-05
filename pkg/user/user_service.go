package user

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"mcm-api/pkg/log"
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

func (s *Service) FindById(ctx context.Context, id int) (*UserResponse, error) {
	entity, err := s.repository.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(apperror.ErrNotFound, "user not found", err)
		}
		return nil, err
	}
	return mapEntityToResponse(entity), nil
}

func (s *Service) FindByEmailAndPassword(ctx context.Context, email string, password string) (*UserResponse, error) {
	entity, err := s.repository.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(apperror.ErrNotFound, "user not found", err)
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(password))
	if err != nil {
		return nil, apperror.New(apperror.ErrInvalid, "wrong password", err)
	}
	return mapEntityToResponse(entity), nil
}

func (s *Service) CreateUser(ctx context.Context, loggedInUser *common.LoggedInUser, req *UserCreateReq) (*UserResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	entity, err := s.repository.FindByEmail(ctx, req.Email)
	if err == nil {
		return nil, apperror.New(apperror.ErrConflict, "duplicate email", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}
	entity = &Entity{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}
	err = s.repository.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return mapEntityToResponse(entity), nil
}

func (s *Service) CreateDefaultAdmin(ctx context.Context) error {
	_, err := s.repository.FindByEmail(ctx, s.cfg.AdminEmail)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Logger.Info("skip create default admin")
		return nil
	}
	log.Logger.Info("init admin account", zap.String("email", s.cfg.AdminEmail))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.cfg.AdminPassword), 10)
	if err != nil {
		return err
	}
	return s.repository.Create(ctx, &Entity{
		Name:     "Administrators",
		Email:    s.cfg.AdminEmail,
		Password: string(hashedPassword),
		Role:     common.Administrator,
	})
}

func (s *Service) Find(ctx context.Context, user *common.LoggedInUser, query *UserIndexQuery) (*common.PaginateResponse, error) {
	entities, count, err := s.repository.FindAndCount(ctx, query)
	if err != nil {
		return nil, err
	}
	dtos := mapEntitiesToResponse(entities)
	return common.NewPaginateResponse(dtos, count, query.Page, query.GetLimit()), nil
}

func mapEntitiesToResponse(entity []*Entity) []*UserResponse {
	var result []*UserResponse
	for i := range entity {
		result = append(result, mapEntityToResponse(entity[i]))
	}
	return result
}

func mapEntityToResponse(entity *Entity) *UserResponse {
	return &UserResponse{
		Id:        entity.Id,
		Name:      entity.Name,
		Role:      entity.Role,
		Email:     entity.Email,
		FacultyId: entity.FacultyId,
		TrackTime: common.TrackTime{
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
	}
}
