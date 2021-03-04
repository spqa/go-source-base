package common

type PaginateQuery struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}

func (q PaginateQuery) GetOffSet() int {
	return q.GetLimit() * q.Page
}

func (q PaginateQuery) GetLimit() int {
	if q.Limit > 50 {
		return 50
	}
	if q.Limit <= 0 {
		return 20
	}
	return q.Limit
}

type PaginateResponse struct {
	Total       int64       `json:"total"`
	CurrentPage int         `json:"currentPage"`
	LastPage    int         `json:"lastPage"`
	PerPage     int         `json:"perPage"`
	Data        interface{} `json:"data"`
}

func NewPaginateResponse(data interface{}, total int64, page int, limit int) *PaginateResponse {
	return &PaginateResponse{
		Total:       total,
		CurrentPage: page,
		LastPage:    calculateLastPage(total, limit),
		PerPage:     limit,
		Data:        data,
	}
}

func calculateLastPage(total int64, limit int) int {
	if total%int64(limit) > 0 {
		return int(total / int64(limit))
	}
	return int(total/int64(limit)) - 1
}
