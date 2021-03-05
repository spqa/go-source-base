package user

import "mcm-api/pkg/common"

type UserResponse struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	FacultyId *int        `json:"facultyId"`
	Role      common.Role `json:"role"`
	common.TrackTime
}
