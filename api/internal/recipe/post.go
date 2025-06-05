package recipe

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/controller"
	"net/http"
)

type Post struct {
	*ActionBase
}

func NewPost(action *ActionBase) *controller.ControllerWithBody[PostRecipeRequest] {
	return controller.NewControllerWithBody(&Post{action})
}
func (p *Post) Process(body PostRecipeRequest, r *http.Request) controller.Response {
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
