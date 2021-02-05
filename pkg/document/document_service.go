package document

type Service struct {
	repository *repository
}

func InitializeService(repository *repository) *Service {
	return &Service{
		repository: repository,
	}
}
