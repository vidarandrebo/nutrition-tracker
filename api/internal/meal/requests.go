package meal

import "time"

type PostMealRequest struct {
	Timestamp time.Time `json:"timestamp"`
}
