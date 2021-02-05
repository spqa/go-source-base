package user

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func initializeRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}
