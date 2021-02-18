package common

type PaginateQuery struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

type PaginateResponse struct {
	Total       int `json:"total"`
	CurrentPage int `json:"currentPage"`
	LastPage    int `json:"LastPage"`
	PerPage     int `json:"perPage"`
}
