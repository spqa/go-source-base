package response

import "net/http"

type ApiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Err        error  `json:"-"`
}

func (a *ApiError) Error() string {
	return a.Message
}

func newApiError(status int, message string, err error) *ApiError {
	return &ApiError{
		Message:    message,
		StatusCode: status,
		Err:        err,
	}
}

func NewApiInternalError(err error) *ApiError {
	return newApiError(http.StatusInternalServerError, "Internal Server Error", err)
}

func NewApiBadRequestError(message string, err error) *ApiError {
	return newApiError(http.StatusBadRequest, message, err)
}
