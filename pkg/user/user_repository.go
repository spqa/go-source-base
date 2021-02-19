package user

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func InitializeRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r repository) FindById(id int) (*Entity, error) {
	var result Entity
	db := r.db.First(&result, id)
	return &result, db.Error
}
