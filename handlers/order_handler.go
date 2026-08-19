package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"mini-order-api/models"
	"mini-order-api/services"

	"github.com/gin-gonic/gin"
)

// OrderHandler bertugas menangani request dan response HTTP terkait Order
type OrderHandler struct {
	service services.OrderService
}

// NewOrderHandler menginisialisasi OrderHandler dengan dependency injection OrderService
func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// CreateOrder menangani HTTP POST /api/v1/orders
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var input models.CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validasi gagal: format request tidak sesuai",
			"error":   err.Error(),
		})
		return
	}

	order, err := h.service.CreateOrder(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal membuat pesanan",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Pesanan berhasil dibuat",
		"data":    order,
	})
}

// GetAllOrders menangani HTTP GET /api/v1/orders
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data pesanan",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Daftar pesanan berhasil diambil",
		"total":   len(orders),
		"data":    orders,
	})
}

// GetOrderByID menangani HTTP GET /api/v1/orders/:id
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID pesanan harus berupa angka bulat positif",
		})
		return
	}

	order, err := h.service.GetOrderByID(uint(id))
	if err != nil {
		if errors.Is(err, services.ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil detail pesanan",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Detail pesanan berhasil ditemukan",
		"data":    order,
	})
}

// PaymentWebhook menangani HTTP POST /api/v1/webhooks/payment
func (h *OrderHandler) PaymentWebhook(c *gin.Context) {
	var payload models.PaymentWebhookPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Payload webhook tidak valid. Pastikan payment_status bernilai 'paid', 'failed', atau 'cancelled'",
			"error":   err.Error(),
		})
		return
	}

	order, err := h.service.ProcessPaymentWebhook(payload)
	if err != nil {
		if errors.Is(err, services.ErrOrderNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		if errors.Is(err, services.ErrOrderAlreadyPaid) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": err.Error(),
				"data":    order,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal memproses webhook pembayaran",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"message":        "Status pesanan berhasil diperbarui melalui Webhook",
		"transaction_id": payload.TransactionID,
		"data":           order,
	})
}
