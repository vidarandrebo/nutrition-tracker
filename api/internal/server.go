package internal

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/meal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/recipe"
)

type recipeEndpoint = recipe.Endpoint
type mealEndpoint = meal.Endpoint
type foodItemEndpoint = fooditem.Endpoint
type authEndpoint = auth.Endpoint

type Server struct {
	recipeEndpoint
	mealEndpoint
	foodItemEndpoint
	authEndpoint
}

func NewServer(recipeEndpoint recipeEndpoint, mealEndpoint mealEndpoint, foodItemEndpoint foodItemEndpoint, authEndpoint authEndpoint) *Server {
	return &Server{recipeEndpoint: recipeEndpoint, mealEndpoint: mealEndpoint, foodItemEndpoint: foodItemEndpoint, authEndpoint: authEndpoint}
}
