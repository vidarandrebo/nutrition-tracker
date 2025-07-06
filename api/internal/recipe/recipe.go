package recipe

import "github.com/vidarandrebo/nutrition-tracker/api/internal/api"

type Recipe struct {
	ID      int64
	Name    string
	Entries []Entry
	OwnerID int64
}

func (r Recipe) ToResponse() api.RecipeResponse {
	entries := make([]api.RecipeEntryResponse, 0, len(r.Entries))
	for _, e := range r.Entries {
		entries = append(entries, e.ToResponse())
	}
	return api.RecipeResponse{
		Id:      r.ID,
		Name:    r.Name,
		Entries: entries,
	}
}
