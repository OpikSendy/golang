package repositories

import (
	"mini-order-api/models"

	"gorm.io/gorm"
)

// OrderRepository mendefinisikan kontrak interface akses database untuk entitas Order
type OrderRepository interface {
	Create(order *models.Order) error
	FindAll() ([]models.Order, error)
	FindByID(id uint) (*models.Order, error)
	UpdateStatus(order *models.Order, status string) error
}

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository menginisialisasi implementasi OrderRepository dengan dependency GORM
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// Create menyimpan data order baru ke PostgreSQL
func (r *orderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

// FindAll mengambil seluruh data order diurutkan dari ID terbaru
func (r *orderRepository) FindAll() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Order("id desc").Find(&orders).Error
	return orders, err
}

// FindByID mencari satu data order berdasarkan ID
func (r *orderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateStatus mengubah kolom status pesanan di database
func (r *orderRepository) UpdateStatus(order *models.Order, status string) error {
	return r.db.Model(order).Update("status", status).Error
}
