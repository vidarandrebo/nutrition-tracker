package utils

import (
	"encoding/json"
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
func Unauthorized() Response {
	return Response{nil, http.StatusUnauthorized}
}
func BadRequest() Response {
	return Response{nil, http.StatusBadRequest}
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
	if response.data != nil {
		enc := json.NewEncoder(w)
		enc.Encode(response.data)
	}
}
func NewController[T any](action Action[T]) *Controller[T] {
	return &Controller[T]{action}
}
