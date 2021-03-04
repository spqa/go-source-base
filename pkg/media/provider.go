package media

import "github.com/google/wire"

var Set = wire.NewSet(NewStorageService, NewDarthsimImageProxyService)
