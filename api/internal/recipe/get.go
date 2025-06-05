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
	recipe, err := p.store.Add(Recipe{OwnerID: userID})
	if err != nil {
		return controller.BadRequest()
	}
	return controller.Created(recipe.ToResponse())
}
