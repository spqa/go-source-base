//+build wireinject

package internal

import (
	"github.com/google/wire"
	"mcm-api/internal/core"
	"mcm-api/internal/router"
	"mcm-api/pkg/document"
	"mcm-api/pkg/user"
)

func InitializeServer() *Server {
	panic(wire.Build(
		core.Set,
		user.Set,
		document.Set,
		router.Set,
		newServer,
	))
}
