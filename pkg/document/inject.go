//+build wireinject

package document

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InjectService(db *gorm.DB) *Service {
	panic(wire.Build(NewRepository, NewService))
}
