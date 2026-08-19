package services

import (
	"errors"
	"time"

	"mini-order-api/models"
	"mini-order-api/repositories"

	"gorm.io/gorm"
)

// Definisi sentinel errors untuk domain business logic
var (
	ErrOrderNotFound    = errors.New("pesanan tidak ditemukan")
	ErrOrderAlreadyPaid = errors.New("pesanan sudah berstatus paid sebelumnya")
)

// OrderService mendefinisikan kontrak interface untuk business logic Order
type OrderService interface {
	CreateOrder(input models.CreateOrderInput) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	ProcessPaymentWebhook(payload models.PaymentWebhookPayload) (*models.Order, error)
}

type orderService struct {
	repo repositories.OrderRepository
}

// NewOrderService menginisialisasi OrderService dengan dependency injection OrderRepository
func NewOrderService(repo repositories.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

// CreateOrder memvalidasi dan memproses pembuatan order baru
func (s *orderService) CreateOrder(input models.CreateOrderInput) (*models.Order, error) {
	order := &models.Order{
		CustomerName: input.CustomerName,
		ItemName:     input.ItemName,
		Amount:       input.Amount,
		Status:       "pending",
	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

// GetAllOrders mengambil seluruh daftar pesanan
func (s *orderService) GetAllOrders() ([]models.Order, error) {
	return s.repo.FindAll()
}

// GetOrderByID mencari detail pesanan berdasarkan ID
func (s *orderService) GetOrderByID(id uint) (*models.Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	return order, nil
}

// ProcessPaymentWebhook menangani alur logika webhook pembayaran
func (s *orderService) ProcessPaymentWebhook(payload models.PaymentWebhookPayload) (*models.Order, error) {
	order, err := s.repo.FindByID(payload.OrderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	// Idempotency check: jika status sudah paid
	if order.Status == "paid" {
		return order, ErrOrderAlreadyPaid
	}

	// Update status pesanan
	if err := s.repo.UpdateStatus(order, payload.PaymentStatus); err != nil {
		return nil, err
	}

	order.Status = payload.PaymentStatus
	order.UpdatedAt = time.Now()

	return order, nil
}
