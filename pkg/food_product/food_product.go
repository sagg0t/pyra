package foodproduct

import "time"

type FoodProduct struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
