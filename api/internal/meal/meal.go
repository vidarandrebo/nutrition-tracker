package meal

import (
	"time"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
)

type Meal struct {
	ID             int64
	SequenceNumber int64
	Timestamp      time.Time
	Entries        []Entry
	OwnerID        int64
}

func (m Meal) ToResponse() api.MealResponse {
	entries := make([]api.MealEntryResponse, 0, len(m.Entries))

	for _, e := range m.Entries {
		entries = append(entries, e.ToResponse())
	}
	return api.MealResponse{
		Id:             m.ID,
		SequenceNumber: m.SequenceNumber,
		Timestamp:      m.Timestamp,
		Entries:        entries,
	}
}
