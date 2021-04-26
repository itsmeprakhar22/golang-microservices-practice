package utils

type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode uint64 `json:"status_code"`
	Code       string `json:"code"`
}
