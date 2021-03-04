//+build wireinject

package server

import (
	"github.com/google/wire"
	"mcm-api/internal/core"
	"mcm-api/pkg/authz"
	"mcm-api/pkg/contributesession"
	"mcm-api/pkg/contribution"
	"mcm-api/pkg/document"
	"mcm-api/pkg/faculty"
	"mcm-api/pkg/media"
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
		media.Set,
		contributesession.Set,
		contribution.Set,

		core.HandlerSet,
		newServer,
	))
}
