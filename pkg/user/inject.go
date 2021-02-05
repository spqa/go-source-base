//+build wireinject

package user

import (
	"github.com/google/wire"
	"mcm-api/internal/provider"
)

func CreateUserService() *Service {
	panic(wire.Build(provider.ProvideConfig, provider.ProvideDB, initializeRepository, initializeService))
}
