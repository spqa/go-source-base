package document

type Service struct {
	repository *repository
}

func NewService(repository *repository) *Service {
	return &Service{
		repository: repository,
	}
}
