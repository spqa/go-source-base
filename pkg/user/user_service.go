package user

import "mcm-api/pkg/user/requests"

type Service struct {
	repository *repository
}

func InitializeService(repository *repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) index(query requests.Query)  {
	s.repository.
}
