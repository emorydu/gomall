package model

import "time"

type Base struct {
	ID        int `grom:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
