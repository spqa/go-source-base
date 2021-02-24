package user

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

func (r *repository) FindById(ctx context.Context, id int) (*Entity, error) {
	var result Entity
	db := r.db.WithContext(ctx).First(&result, id)
	return &result, db.Error
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*Entity, error) {
	var result Entity
	db := r.db.WithContext(ctx).First(&result, "email = ?", email)
	return &result, db.Error
}

func (r *repository) CreateUser(ctx context.Context, entity *Entity) error {
	save := r.db.WithContext(ctx).Create(entity)
	return save.Error
}
