package models

import "time"

// Order merepresentasikan entitas pesanan
type Order struct {
	ID           uint      `json:"id"`
	CustomerName string    `json:"customer_name"`
	ItemName     string    `json:"item_name"`
	Amount       float64   `json:"amount"`
	Status       string    `json:"status"` // "pending", "paid", "cancelled"
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
