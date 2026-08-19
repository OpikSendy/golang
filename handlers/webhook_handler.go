package handlers

import (
	"errors"
	"net/http"

	"mini-order-api/config"
	"mini-order-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PaymentWebhook menangani notifikasi update status pembayaran dari Payment Gateway
func PaymentWebhook(c *gin.Context) {
	var payload models.PaymentWebhookPayload

	// 1. Parsing & validasi payload webhook dari request body
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Payload webhook tidak valid. Pastikan payment_status bernilai 'paid', 'failed', atau 'cancelled'",
			"error":   err.Error(),
		})
		return
	}

	// 2. Cari data order yang sesuai berdasarkan OrderID di database
	var order models.Order
	if err := config.DB.First(&order, payload.OrderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Pesanan dengan order_id tersebut tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Terjadi kesalahan saat mencari data pesanan",
			"error":   err.Error(),
		})
		return
	}

	// 3. Cek apakah status pesanan sudah 'paid' sebelumnya (Idempotency check)
	if order.Status == "paid" {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Pesanan ini sudah berstatus 'paid' sebelumnya",
			"data":    order,
		})
		return
	}

	// 4. Update status pesanan di database PostgreSQL
	if err := config.DB.Model(&order).Update("status", payload.PaymentStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal memperbarui status pesanan",
			"error":   err.Error(),
		})
		return
	}

	// 5. Response sukses webhook ke Payment Gateway
	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"message":        "Status pesanan berhasil diperbarui melalui Webhook",
		"transaction_id": payload.TransactionID,
		"data":           order,
	})
}
