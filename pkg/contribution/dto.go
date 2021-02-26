package contribution

import (
	"mcm-api/pkg/common"
	"mcm-api/pkg/user"
)

type IndexQuery struct {
	common.PaginateQuery
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

type ContributionCreateReq struct {
	Article ArticleReq `json:"article"`
	Images  []string   `json:"images"`
}

type ContributionUpdateReq struct {
	Article ArticleReq `json:"article"`
	Images  []string   `json:"images"`
}
