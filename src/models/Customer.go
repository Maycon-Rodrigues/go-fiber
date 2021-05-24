package models

type Customer struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
	Phone     int    `json:"phone"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
