package notification

import "github.com/google/wire"

var Set = wire.NewSet(InitializeService)
