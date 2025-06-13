package recipe

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/controller"
	"net/http"
)

type Get struct {
	*ActionBase
}

func NewGet(action *ActionBase) *controller.Controller {
	return controller.NewController(&Get{action})
}
func (p *Get) Process(r *http.Request) controller.Response {
	userID, err := auth.UserIDFromCtx(r.Context())
	if err != nil {
		return controller.Unauthorized()
	}
	recipes, err := p.store.Get(userID)
	if err != nil {
		return controller.BadRequest()
	}
	responses := make([]RecipeResponse, 0, len(recipes))
	for _, recipe := range recipes {
		responses = append(responses, recipe.ToResponse())
	}
	return controller.Ok(responses)
}
