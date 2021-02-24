package user

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
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
	err = s.repository.CreateUser(ctx, entity)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s.cfg.AdminPassword), 10)
	if err != nil {
		return err
	}
	return s.repository.CreateUser(ctx, &Entity{
		Name:     "Administrators",
		Email:    s.cfg.AdminEmail,
		Password: string(hashedPassword),
		Role:     common.Administrator,
	})
}

func mapEntityToResponse(entity *Entity) *UserResponse {
	return &UserResponse{
		Id:    entity.Id,
		Name:  entity.Name,
		Role:  entity.Role,
		Email: entity.Email,
		TrackTime: common.TrackTime{
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		},
	}
}
