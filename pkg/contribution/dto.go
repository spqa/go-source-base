package contribution

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"mcm-api/pkg/common"
	"mcm-api/pkg/user"
)

type IndexQuery struct {
	common.PaginateQuery
	FacultyId             *int   `json:"facultyId"`
	StudentId             *int   `json:"studentId"`
	ContributionSessionId *int   `json:"contributionSessionId"`
	Status                Status `json:"status" enums:"accepted,rejected,reviewing"`
}

type ContributionRes struct {
	Id                  int64             `json:"id"`
	User                user.UserResponse `json:"user"`
	ContributeSessionId int64             `json:"contributeSessionId"`
	ArticleId           int64             `json:"articleId"`
	Images              []string          `json:"images"`
	Status              Status            `json:"status"`
	common.TrackTime
}

type ArticleReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

func (r *ArticleReq) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Title, validation.Required, validation.Length(10, 255)),
		validation.Field(&r.Description, validation.Length(15, 512)),
		validation.Field(&r.Link, validation.Required, validation.Length(10, 255)),
	)
}

type ContributionCreateReq struct {
	Article ArticleReq `json:"article"`
	Images  []string   `json:"images"`
}

func (r *ContributionCreateReq) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Article, validation.Required),
	)
}

type ContributionUpdateReq struct {
	Article ArticleReq `json:"article"`
	Images  []string   `json:"images"`
}

func (r *ContributionUpdateReq) Validate() error {
	return validation.ValidateStruct(r)
}
