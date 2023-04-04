package domain

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	CompanyId string
	CreatedAt time.Time
	UpdatedAt time.Time
}
