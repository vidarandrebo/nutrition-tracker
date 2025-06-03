package recipe

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"net/http"
)

type Post struct {
	*ActionBase
}

func NewPost(action *ActionBase) *utils.Controller[PostRecipeRequest] {
	return utils.NewController(&Post{action})
}
func (p *Post) Process(body *PostRecipeRequest, r *http.Request) utils.Response {
	userID, err := auth.UserIDFromCtx(r.Context())
	if err != nil {
		return utils.Unauthorized()
	}
	recipe, err := p.store.Add(Recipe{OwnerID: userID})
	if err != nil {
		return utils.BadRequest()
	}
	return utils.Created(recipe.ToResponse())
}
