package recipe

import (
	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
)

type Recipe struct {
	ID              int64
	Name            string
	FoodItemEntries []*Entry
	OwnerID         int64
}

func NewRecipe() *Recipe {
	return &Recipe{}
}

func (r *Recipe) ToResponse() api.RecipeResponse {
	entries := make([]api.RecipeFoodItemEntryEntryResponse, 0, len(r.FoodItemEntries))
	for _, e := range r.FoodItemEntries {
		entries = append(entries, e.ToResponse())
	}
	return api.RecipeResponse{
		Id:              r.ID,
		Name:            r.Name,
		FoodItemEntries: entries,
	}
}

func (r *Recipe) ToTable() TableRecipe {
	return TableRecipe{
		ID:      r.ID,
		Name:    r.Name,
		OwnerID: r.OwnerID,
	}
}

func FromRecipeTable(tbl TableRecipe) *Recipe {
	return &Recipe{
		ID:              tbl.ID,
		Name:            tbl.Name,
		FoodItemEntries: make([]*Entry, 0),
		OwnerID:         tbl.OwnerID,
	}
}

func (r *Recipe) FromPost(request *api.RecipePostRequest) *Recipe {
	entries := make([]*Entry, 0, len(request.FoodItemEntries))
	for _, e := range request.FoodItemEntries {
		entries = append(entries, FromEntryPost(e))
	}
	r.Name = request.Name
	r.FoodItemEntries = entries
	return r
}
