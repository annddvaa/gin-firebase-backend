package models

import "gorm.io/gorm"

// Order mewakili satu transaksi pesanan oleh user
type Order struct {
	gorm.Model
	UserID        uint        `gorm:"not null" json:"user_id"`
	TotalPrice    float64     `gorm:"not null" json:"total_price"`
	Status        string      `gorm:"type:enum('menunggu', 'dikerjakan', 'selesai', 'diambil');default:'menunggu'" json:"status"`
	PaymentStatus string      `gorm:"type:enum('unpaid', 'paid');default:'unpaid'" json:"payment_status"`
	OrderItems    []OrderItem `json:"order_items"`
}

// OrderItem mewakili detail produk yang dipesan dalam satu Order
type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"` // Harga satuan saat pesanan dibuat

	Product Product `gorm:"foreignKey:ProductID" json:"product"` // Relasi ke tabel Product
}
