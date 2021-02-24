package startup

import "github.com/google/wire"

var Set = wire.NewSet(InitializeStartUpService)
