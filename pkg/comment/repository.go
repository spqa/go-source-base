package comment

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

func (r repository) Find(ctx context.Context, query *IndexQuery) ([]*Entity, error) {
	var results []*Entity
	r.db.WithContext(ctx)
	r.db.Limit(query.GetLimit())
	r.db.Offset(query.GetOffSet())
	db := r.db.Find(results)
	return results, db.Error
}

func (r repository) Create(ctx context.Context, entity *Entity) (*Entity, error) {
	db := r.db.WithContext(ctx).Create(entity)
	return entity, db.Error
}

func (r repository) Update(ctx context.Context, entity *Entity) (*Entity, error) {
	db := r.db.WithContext(ctx).Save(entity)
	return entity, db.Error
}

func (r repository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(id).Error
}
