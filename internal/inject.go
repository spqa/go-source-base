//+build wireinject

package internal

import (
	"github.com/google/wire"
	"mcm-api/pkg/user"
)

func InjectUserRouter() *UserRouter {
	panic(wire.Build(user.CreateUserService, NewUserRouter))
}

func InjectDocumentRouter() {

}
