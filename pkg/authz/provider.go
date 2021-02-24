package authz

import "github.com/google/wire"

var Set = wire.NewSet(InitializeAuthService)
