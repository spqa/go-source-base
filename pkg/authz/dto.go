package authz

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"mcm-api/pkg/user"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	*user.UserResponse
}
