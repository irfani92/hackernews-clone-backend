package domain

import "time"

type User struct {
	ID        string  `json:"id"`
	Created   int64   `json:"created"`
	Karma     int     `json:"karma"`
	About     *string `json:"about"`
	Submitted []int   `json:"submitted"`
	UpdatedAt time.Time
}
