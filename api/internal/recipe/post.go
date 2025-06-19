package recipe

import (
	"net/http"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/controller"
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
	entries := make([]Entry, 0, len(body.Entries))
	for _, e := range body.Entries {
		entries = append(entries, Entry{
			Amount:     e.Amount,
			FoodItemID: e.FoodItemID,
		})
	}
	recipe, err := p.store.Add(Recipe{Name: body.Name, Entries: entries, OwnerID: userID})
	if err != nil {
		return controller.BadRequest()
	}
	return controller.Created(recipe.ToResponse())
}
