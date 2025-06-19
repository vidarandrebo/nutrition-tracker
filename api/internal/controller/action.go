package controller

import "net/http"

type Action interface {
	Process(r *http.Request) Response
}
type ActionWithBody[T any] interface {
	Process(requestPayload T, r *http.Request) Response
}
