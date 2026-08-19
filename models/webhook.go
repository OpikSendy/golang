package models

// PaymentWebhookPayload adalah DTO untuk menerima payload webhook simulasi pembayaran
type PaymentWebhookPayload struct {
	OrderID       uint   `json:"order_id" binding:"required"`
	PaymentStatus string `json:"payment_status" binding:"required,oneof=paid failed cancelled"`
	TransactionID string `json:"transaction_id" binding:"required"`
}
