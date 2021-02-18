package responses

import "mcm-api/pkg/common"

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	common.TrackTime
}
