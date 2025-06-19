package recipe

type Recipe struct {
	ID      int64
	Name    string
	Entries []Entry
	OwnerID int64
}

func (r Recipe) ToResponse() RecipeResponse {
	entries := make([]EntryResponse, 0, len(r.Entries))
	for _, e := range r.Entries {
		entries = append(entries, e.ToResponse())
	}
	return RecipeResponse{
		ID:      r.ID,
		Name:    r.Name,
		Entries: entries,
	}
}
