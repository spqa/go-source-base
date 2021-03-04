package common

const (
	limitDefault = 20
	limitMax     = 100
)

type PaginateQuery struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}

func (q PaginateQuery) GetOffSet() int {
	return q.GetLimit() * q.Page
}

func (q PaginateQuery) GetLimit() int {
	if q.Limit > limitMax {
		return limitMax
	}
	if q.Limit <= 0 {
		return limitDefault
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

func NewEmptyPaginateResponse() *PaginateResponse {
	return &PaginateResponse{
		Total:       0,
		CurrentPage: 0,
		LastPage:    0,
		PerPage:     limitDefault,
		Data:        nil,
	}
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
	if total == 0 {
		return 0
	}
	if total%int64(limit) > 0 {
		return int(total / int64(limit))
	}
	return int(total/int64(limit)) - 1
}
