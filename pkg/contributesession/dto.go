package contributesession

import (
	"mcm-api/pkg/common"
	"time"
)

type SessionRes struct {
	Id               int64     `json:"id"`
	OpenTime         time.Time `json:"openTime"`
	ClosureTime      time.Time `json:"closureTime"`
	FinalClosureTime time.Time `json:"finalClosureTIme"`
	ExportedAssets   string    `json:"exportedAssets"`
	common.TrackTime
}

type SessionCreateReq struct {
	OpenTime         time.Time `json:"openTime"`
	ClosureTime      time.Time `json:"closureTime"`
	FinalClosureTime time.Time `json:"finalClosureTIme"`
}

type SessionUpdateReq struct {
	OpenTime         time.Time `json:"openTime"`
	ClosureTime      time.Time `json:"closureTime"`
	FinalClosureTime time.Time `json:"finalClosureTIme"`
}

type IndexQuery struct {
	common.PaginateQuery
}
