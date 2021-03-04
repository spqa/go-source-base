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
	result := new(Entity)
	db := r.db.WithContext(ctx).First(result, id)
	return result, db.Error
}

func (r repository) Find(ctx context.Context, query IndexQuery) ([]*Entity, error) {
	return nil, nil
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
	return r.db.WithContext(ctx).Delete(&Entity{Id: id}).Error
}

func (r repository) FindAndCount(ctx context.Context, query *IndexQuery) ([]*Entity, int64, error) {
	var entities []*Entity
	builder := r.db.WithContext(ctx).Model(&Entity{})
	var count int64
	result := builder.Count(&count)
	if result.Error != nil {
		return nil, 0, nil
	}
	builder.Offset(query.GetOffSet()).Limit(query.GetLimit())
	result = builder.Find(&entities)
	return entities, count, result.Error
}
