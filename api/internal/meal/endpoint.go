package meal

import "github.com/vidarandrebo/nutrition-tracker/api/internal/api"

type Endpoint struct {
}

func (e Endpoint) GetApiMeals(ctx context.Context, request api.GetApiMealsRequestObject) (api.GetApiMealsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (e Endpoint) PostApiMeals(ctx context.Context, request api.PostApiMealsRequestObject) (api.PostApiMealsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (e Endpoint) GetApiMealsId(ctx context.Context, request api.GetApiMealsIdRequestObject) (api.GetApiMealsIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (e Endpoint) PostApiMealsIdEntries(ctx context.Context, request api.PostApiMealsIdEntriesRequestObject) (api.PostApiMealsIdEntriesResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
