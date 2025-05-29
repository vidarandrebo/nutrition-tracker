package utils

import (
	"net/http"
)

type Response struct {
	data       any
	statusCode int
}

func Created(data any) Response {
	return Response{data: data, statusCode: http.StatusCreated}
}
func Ok(data any) Response {
	return Response{data: data, statusCode: http.StatusOK}
}

type Action[T any] interface {
	Process(requestPayload *T, r *http.Request) Response
}
type Controller[T any] struct {
	Action[T]
}

func (c *Controller[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request, err := ParseJson[T](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := c.Process(request, r)
	w.WriteHeader(response.statusCode)
}
func NewController[T any](action Action[T]) *Controller[T] {
	return &Controller[T]{action}
}
