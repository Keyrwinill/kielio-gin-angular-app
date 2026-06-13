package dto

type Response[T any] struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	Data      T      `json:"data"`
}
