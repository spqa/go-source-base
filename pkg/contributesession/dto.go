package contributesession

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"mcm-api/pkg/common"
	"time"
)

type SessionRes struct {
	Id               int       `json:"id"`
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

func (s SessionCreateReq) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.OpenTime, validation.Required),
		validation.Field(&s.ClosureTime,
			validation.Required,
			validation.Min(s.OpenTime),
		),
		validation.Field(&s.FinalClosureTime,
			validation.Required,
			validation.Min(s.ClosureTime),
		),
	)
}

type SessionUpdateReq struct {
	OpenTime         time.Time `json:"openTime"`
	ClosureTime      time.Time `json:"closureTime"`
	FinalClosureTime time.Time `json:"finalClosureTIme"`
}

func (s SessionUpdateReq) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.OpenTime, validation.Required),
		validation.Field(&s.ClosureTime,
			validation.Required,
			validation.Min(s.OpenTime),
		),
		validation.Field(&s.FinalClosureTime,
			validation.Required,
			validation.Min(s.ClosureTime),
		),
	)
}

type IndexQuery struct {
	common.PaginateQuery
}
