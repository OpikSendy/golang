package models

import "time"

// Order merepresentasikan entitas pesanan dan tabel "orders" di database PostgreSQL
type Order struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerName string    `gorm:"type:varchar(100);not null" json:"customer_name"`
	ItemName     string    `gorm:"type:varchar(150);not null" json:"item_name"`
	Amount       float64   `gorm:"type:decimal(12,2);not null" json:"amount"`
	Status       string    `gorm:"type:varchar(50);not null;default:'pending'" json:"status"` // pending, paid, cancelled
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreateOrderInput adalah Data Transfer Object (DTO) untuk validasi request payload saat membuat pesanan baru
type CreateOrderInput struct {
	CustomerName string  `json:"customer_name" binding:"required,min=3,max=100"`
	ItemName     string  `json:"item_name" binding:"required,min=2,max=150"`
	Amount       float64 `json:"amount" binding:"required,gt=0"`
}
