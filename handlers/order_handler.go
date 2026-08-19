package handlers

import (
	"errors"
	"net/http"

	"mini-order-api/config"
	"mini-order-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder membuat data order baru dan menyimpannya ke database PostgreSQL
func CreateOrder(c *gin.Context) {
	var input models.CreateOrderInput

	// 1. Binding & Validasi JSON Body menggunakan ShouldBindJSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validasi gagal: format request tidak sesuai",
			"error":   err.Error(),
		})
		return
	}

	// 2. Inisialisasi entity Model Order dengan status default 'pending'
	order := models.Order{
		CustomerName: input.CustomerName,
		ItemName:     input.ItemName,
		Amount:       input.Amount,
		Status:       "pending",
	}

	// 3. Simpan ke database menggunakan GORM Create()
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal menyimpan pesanan ke database",
			"error":   err.Error(),
		})
		return
	}

	// 4. Return response sukses HTTP 201 Created
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Pesanan berhasil dibuat",
		"data":    order,
	})
}

// GetAllOrders mengambil seluruh daftar pesanan dari PostgreSQL
func GetAllOrders(c *gin.Context) {
	var orders []models.Order

	// Mengambil semua data dengan query ORDER BY id DESC
	if err := config.DB.Order("id desc").Find(&orders).Error; err != nil {
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

// GetOrderByID mengambil detail pesanan spesifik berdasarkan parameter ID
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	// Mencari data order dengan ID yang sesuai
	if err := config.DB.First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Pesanan tidak ditemukan",
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
