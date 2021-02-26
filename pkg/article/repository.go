package article

import (
	"context"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func InitializeRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r repository) FindById(ctx context.Context, id int) (*Entity, error) {
	result := new(Entity)
	db := r.db.WithContext(ctx).First(result, id)
	return result, db.Error
}
