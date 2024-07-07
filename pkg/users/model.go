package users

import "time"

type User struct {
	ID uint64

	Email string

	FirstName string
	LastName  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
