package meal

import "time"

type Meal struct {
	ID             int64
	SequenceNumber int64
	Timestamp      time.Time
	Entries        []Entry
	OwnerID        int64
}
