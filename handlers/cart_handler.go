package handlers

import (
	"log"
	"net/http"
	"strconv"
	"github.com/annddvaa/gin-firebase-backend/services"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *services.CartService
}

func NewCartHandler() *CartHandler {
	return &CartHandler{
		service: services.NewCartService(),
	}
}

// SAFE USER ID 
func getUserID(c *gin.Context) (uint, bool) {
	val, exists := c.Get("user_id")
	if !exists {
		log.Println(" [CartHandler] user_id tidak ditemukan di context")
		return 0, false
	}

	if floatVal, ok := val.(float64); ok {
		return uint(floatVal), true
	}

	if uintVal, ok := val.(uint); ok {
		return uintVal, true
	}

	if strVal, ok := val.(string); ok {
		id, err := strconv.ParseUint(strVal, 10, 32)
		if err == nil {
			return uint(id), true
		}
	}

	log.Printf(" [CartHandler] Gagal konversi user_id. Tipe data asli: %T", val)
	return 0, false
}

// GET CART 
func (h *CartHandler) GetCart(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized: gagal identifikasi user"})
		return
	}

	items, err := h.service.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    items,
	})
}

// ADD TO CART 
func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "input tidak valid"})
		return
	}

	if req.Quantity <= 0 {
		req.Quantity = 1
	}

	if err := h.service.Add(userID, req.ProductID, req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "berhasil menambahkan ke keranjang",
	})
}

// UPDATE QTY 
func (h *CartHandler) UpdateQty(c *gin.Context) {
	_, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID item tidak valid"})
		return
	}

	var req struct {
		Quantity int `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "input tidak valid"})
		return
	}

	if err := h.service.UpdateQty(uint(id), req.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "kuantitas berhasil diperbarui"})
}

// REMOVE ITEM 
func (h *CartHandler) Remove(c *gin.Context) {
	_, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID item tidak valid"})
		return
	}

	if err := h.service.Remove(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item berhasil dihapus dari keranjang"})
}

// CLEAR CART 
func (h *CartHandler) Clear(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if err := h.service.Clear(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "keranjang berhasil dikosongkan"})
}

// CHECKOUT 
func (h *CartHandler) Checkout(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	orderService := services.NewOrderService()
	order, err := orderService.CreateOrderFromCart(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "checkout berhasil",
		"data":    order,
	})
}