package queue

import "github.com/google/wire"

var Set = wire.NewSet(InitializeRedisQueue)
