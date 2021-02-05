package document

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func InitializeRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}
