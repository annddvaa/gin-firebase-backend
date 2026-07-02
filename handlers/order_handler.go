package handlers

import (
	"net/http"

	"github.com/annddvaa/gin-firebase-backend/services"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		service: services.NewOrderService(),
	}
}

// GetOrders mengambil daftar pesanan user
func (h *OrderHandler) GetOrders(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	orders, err := h.service.GetUserOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal mengambil pesanan"})
		return
	}

	// Format response agar lebih rapi untuk mobile app (opsional)
	// Tapi kita bisa mereturn orders secara langsung
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    orders,
	})
}

// CompleteLatestOrder menyelesaikan pesanan terbaru user setelah pembayaran sukses
func (h *OrderHandler) CompleteLatestOrder(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	err := h.service.CompleteLatestOrder(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal update status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
