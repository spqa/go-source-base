package faculty

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"mcm-api/pkg/common"
)

type IndexQuery struct {
	common.PaginateQuery
}

type FacultyResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	common.TrackTime
}

type FacultyCreateReq struct {
	Name string `json:"name"`
}

func (f *FacultyCreateReq) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.Name, validation.Required, validation.Length(6, 100)))
}

type FacultyUpdateReq struct {
	Name string `json:"name"`
}

func (f *FacultyUpdateReq) Validate() error {
	return validation.ValidateStruct(f,
		validation.Field(&f.Name, validation.Required, validation.Length(6, 100)))
}
