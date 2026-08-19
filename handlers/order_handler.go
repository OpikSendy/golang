package handlers

import (
	"net/http"
	"time"

	"mini-order-api/models"

	"github.com/gin-gonic/gin"
)

// GetOrdersExample adalah handler contoh data pesanan di Modul 1
func GetOrdersExample(c *gin.Context) {
	orders := []models.Order{
		{
			ID:           1,
			CustomerName: "Budi Santoso",
			ItemName:     "Kopi Susu Gula Aren",
			Amount:       18000,
			Status:       "pending",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           2,
			CustomerName: "Siti Aminah",
			ItemName:     "Roti Bakar Cokelat",
			Amount:       25000,
			Status:       "paid",
			CreatedAt:    time.Now().Add(-1 * time.Hour),
			UpdatedAt:    time.Now().Add(-30 * time.Minute),
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data order contoh berhasil diambil (Modul 1)",
		"data":    orders,
	})
}
