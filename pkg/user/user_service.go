package user

type Service struct {
	repository *repository
}

func initializeService(repository *repository) *Service {
	return &Service{
		repository: repository,
	}
}
