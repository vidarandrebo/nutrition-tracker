package meal

import "time"

type PostMealRequest struct {
	TimeStamp time.Time `json:"timeStamp"`
}
