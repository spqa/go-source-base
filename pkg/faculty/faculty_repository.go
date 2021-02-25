package faculty

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
	return nil, nil
}

func (r repository) Find(ctx context.Context, query IndexQuery) ([]*Entity, error) {
	return nil, nil
}

func (r repository) Create(ctx context.Context, entity *Entity) (*Entity, error) {
	return nil, nil
}

func (r repository) Update(ctx context.Context, entity *Entity) (*Entity, error) {
	return nil, nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	return nil
}
