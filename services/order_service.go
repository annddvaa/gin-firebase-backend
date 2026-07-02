package services

import (
	"errors"

	"github.com/annddvaa/gin-firebase-backend/config"
	"github.com/annddvaa/gin-firebase-backend/models"
	"gorm.io/gorm"
)

type OrderService struct {
	cartService *CartService
}

func NewOrderService() *OrderService {
	return &OrderService{
		cartService: NewCartService(),
	}
}

// CreateOrderFromCart membuat Order berdasarkan isi cart user saat ini
func (s *OrderService) CreateOrderFromCart(userID uint) (*models.Order, error) {
	items, err := s.cartService.GetCart(userID)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, errors.New("keranjang kosong")
	}

	var totalPrice float64
	var orderItems []models.OrderItem

	for _, item := range items {
		price := float64(item.Product.Price)
		totalPrice += price * float64(item.Quantity)

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     price,
		})
	}

	order := models.Order{
		UserID:        userID,
		TotalPrice:    totalPrice,
		Status:        "menunggu",
		PaymentStatus: "paid", // Asumsi jika sampai kesini berarti checkout + emoney sukses/dianggap lunas untuk simple flow
		OrderItems:    orderItems,
	}

	// Gunakan transaction untuk memastikan data aman
	err = config.DB.Transaction(func(tx *gorm.DB) error {
		// Simpan order beserta order_items
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// Kosongkan keranjang
		if err := tx.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order, nil
}

// GetUserOrders mengambil riwayat pesanan milik user tertentu
func (s *OrderService) GetUserOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := config.DB.Preload("OrderItems.Product").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&orders).Error
	return orders, err
}

// CompleteLatestOrder mengubah status pesanan terbaru user menjadi selesai
func (s *OrderService) CompleteLatestOrder(userID uint) error {
	var latestOrder models.Order
	err := config.DB.Where("user_id = ?", userID).Order("created_at desc").First(&latestOrder).Error
	if err != nil {
		return err
	}

	latestOrder.Status = "selesai"
	latestOrder.PaymentStatus = "paid"
	return config.DB.Save(&latestOrder).Error
}
