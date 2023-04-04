package domain

import "time"

type Company struct {
	ID        string
	Name      string
	Email     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
