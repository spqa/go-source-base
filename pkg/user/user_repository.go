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

func (r *repository) Create(ctx context.Context, entity *Entity) error {
	save := r.db.WithContext(ctx).Create(entity)
	return save.Error
}

func (r *repository) FindAndCount(ctx context.Context, query *UserIndexQuery) ([]*Entity, int64, error) {
	db := r.db.WithContext(ctx).Table("users")
	if query.Role != "" {
		db.Where("role = ?", query.Role)
	}
	var count int64
	db.Count(&count)
	db.Limit(query.GetLimit())
	db.Offset(query.GetOffSet())
	var entities []*Entity
	result := db.Find(&entities)
	return entities, count, result.Error
}
