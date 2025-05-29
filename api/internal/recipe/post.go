package recipe

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"net/http"
)

type Post struct {
	*ActionBase
}

func (p *Post) Process(body *PostRecipeRequest, r *http.Request) utils.Response {
	p.logger.Info("here is am!!!")
	return utils.Created(body)
}

func NewPost(action *ActionBase) *utils.Controller[PostRecipeRequest] {
	return utils.NewController(&Post{action})
}
