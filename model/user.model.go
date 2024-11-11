package model

type User struct {
	Email string
}

type Item struct {
	ID       int     `gorm:"primaryKey" json:"id"`     // Item ID
	Name     string  `gorm:"not null" json:"name"`     // Item Name
	Quantity int     `gorm:"not null" json:"quantity"` // Quantity
	Price    float64 `gorm:"not null" json:"price"`    // Price
	Actions  string  `json:"actions"`                  // Actions (optional)
}
