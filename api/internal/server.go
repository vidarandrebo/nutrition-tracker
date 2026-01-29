package internal

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/fooditem"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/meal"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/recipe"
)

type (
	recipeEndpoint   = recipe.Endpoint
	mealEndpoint     = meal.Endpoint
	foodItemEndpoint = fooditem.Endpoint
	authEndpoint     = auth.Endpoint
)

// StrictServerInterface
type Server struct {
	*recipeEndpoint
	*mealEndpoint
	*foodItemEndpoint
	*authEndpoint
}

func NewServer(recipeEndpoint *recipeEndpoint, mealEndpoint *mealEndpoint, foodItemEndpoint *foodItemEndpoint, authEndpoint *authEndpoint) *Server {
	return &Server{recipeEndpoint: recipeEndpoint, mealEndpoint: mealEndpoint, foodItemEndpoint: foodItemEndpoint, authEndpoint: authEndpoint}
}
