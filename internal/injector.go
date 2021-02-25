//+build wireinject

package internal

import (
	"github.com/google/wire"
	"mcm-api/internal/core"
	"mcm-api/pkg/authz"
	"mcm-api/pkg/document"
	"mcm-api/pkg/faculty"
	"mcm-api/pkg/startup"
	"mcm-api/pkg/user"
)

func InitializeServer() *Server {
	panic(wire.Build(
		core.InfraSet,
		user.Set,
		document.Set,
		authz.Set,
		startup.Set,
		faculty.Set,
		core.HandlerSet,
		newServer,
	))
}
