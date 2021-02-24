package startup

import (
	"context"
	"mcm-api/pkg/user"
	"time"
)

type Service struct {
	userService *user.Service
}

func InitializeStartUpService(service *user.Service) *Service {
	return &Service{
		userService: service,
	}
}

func (s Service) Run() error {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	err := s.userService.CreateDefaultAdmin(timeout)
	cancelFunc()
	return err
}
