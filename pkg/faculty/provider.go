package faculty

import "github.com/google/wire"

var Set = wire.NewSet(InitializeRepository, InitializeService)
