package recipe

type Recipe struct {
	ID      int64
	Name    string
	Entries []Entry
	OwnerID int64
}

func (r Recipe) ToResponse() RecipeResponse {
	return RecipeResponse{}
}
