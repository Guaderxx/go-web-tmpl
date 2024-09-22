package service

type Response[T any] struct {
	Code int
	Data T
}

type ErrorResponse struct {
	Code  int
	Error string
}
