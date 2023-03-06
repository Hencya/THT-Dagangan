package response

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Page struct {
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalData int64 `json:"total_data"`
}
