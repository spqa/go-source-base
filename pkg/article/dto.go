package article

import (
	"mcm-api/pkg/common"
	"time"
)

type ArticleRes struct {
	Id          int64                `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Versions    []*ArticleVersionRes `json:"versions"`
	common.TrackTime
}

type ArticleVersionRes struct {
	Id           int64     `json:"id"`
	LinkOriginal string    `json:"linkOriginal"`
	LinkPdf      string    `json:"linkPdf"`
	CreatedAt    time.Time `json:"createdAt"`
}
