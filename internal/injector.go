//+build wireinject

package internal

import (
	"github.com/google/wire"
	"mcm-api/internal/core"
	"mcm-api/pkg/document"
	"mcm-api/pkg/user"
)

func InitializeServer() *Server {
	panic(wire.Build(
		core.InfraSet,
		user.Set,
		document.Set,
		core.HandlerSet,
		newServer,
	))
}
