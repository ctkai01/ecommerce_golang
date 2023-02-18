package common

import "time"

type SQLModel struct {
	ID	int	`json:"id"`
	CreatedAt *time.Time `json:"create_at"`
	Updated_at *time.Time `json:"updated_at"`
}