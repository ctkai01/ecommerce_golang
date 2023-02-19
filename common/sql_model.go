package common

import "time"

type SQLModel struct {
	ID	int	`json:"id"`
	CreatedAt *time.Time `json:"create_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}