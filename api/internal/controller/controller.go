package controller

import (
	"encoding/json"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"net/http"
)

type Controller struct {
	Action
}
type ControllerWithBody[T any] struct {
	ActionWithBody[T]
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := c.Process(r)
	w.WriteHeader(response.statusCode)
	if response.data != nil {
		enc := json.NewEncoder(w)
		enc.Encode(response.data)
	}
}
func NewController(action Action) *Controller {
	return &Controller{action}
}

func (c *ControllerWithBody[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request, err := utils.ParseJson[T](r.Body)
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
func NewControllerWithBody[T any](action ActionWithBody[T]) *ControllerWithBody[T] {
	return &ControllerWithBody[T]{action}
}
