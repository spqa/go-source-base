package user

import "gorm.io/gorm"

type service struct {
	repository *repository
}

func InitializeService(db *gorm.DB) *service {
	return &service{
		repository: InitializeRepository(db),
	}
}
