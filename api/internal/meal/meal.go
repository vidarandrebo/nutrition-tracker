package meal

import "time"

type Meal struct {
	ID             int64
	SequenceNumber int64
	Timestamp      time.Time
	Entries        []Entry
	OwnerID        int64
}

func (m Meal) ToResponse() MealResponse {
	entries := make([]EntryResponse, 0, len(m.Entries))

	for _, e := range m.Entries {
		entries = append(entries, e.ToResponse())
	}
	return MealResponse{
		ID:             m.ID,
		SequenceNumber: m.SequenceNumber,
		Timestamp:      m.Timestamp,
		Entries:        entries,
	}
}
