package user

import (
	"errors"
	"gorm.io/gorm"
	"mcm-api/pkg/response"
	"mcm-api/pkg/user/responses"
)

type Service struct {
	repository *repository
}

func InitializeService(repository *repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) FindById(id int) (*responses.UserResponse, error) {
	entity, err := s.repository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewApiBadRequestError("User not found", nil)
		}
		return nil, err
	}
	return &responses.UserResponse{
		Id:   entity.Id,
		Name: entity.Name,
	}, nil
}
