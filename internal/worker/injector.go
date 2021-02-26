//+build wireinject

package worker

import (
	"github.com/google/wire"
	"mcm-api/internal/core"
)

func InitializeWorker() *worker {
	panic(wire.Build(core.InfraSet, newWorker))
}
