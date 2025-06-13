package controller

import "net/http"

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
