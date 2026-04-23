package repositories

import (
    "errors"
    "github.com/annddvaa/gin-firebase-backend/config"
    "github.com/annddvaa/gin-firebase-backend/models"
)

type CartRepository struct{}

func NewCartRepository() *CartRepository {
    return &CartRepository{}
}


func (r *CartRepository) FindByUser(userID uint) ([]models.CartItem, error) {
    var items []models.CartItem
    err := config.DB.Preload("Product").
        Where("user_id = ?", userID).
        Find(&items).Error
    return items, err
}


func (r *CartRepository) FindItem(userID, productID uint) (*models.CartItem, error) {
    var item models.CartItem
    err := config.DB.
        Where("user_id = ? AND product_id = ?", userID, productID).
        First(&item).Error
    return &item, err
}


func (r *CartRepository) Create(item *models.CartItem) error {
    return config.DB.Create(item).Error
}


func (r *CartRepository) UpdateQty(id uint, qty int) error {
    if qty < 1 {
        return errors.New("quantity minimal 1")
    }

    return config.DB.Model(&models.CartItem{}).
        Where("id = ?", id).
        Update("quantity", qty).Error
}

// Delete item
func (r *CartRepository) Delete(id uint) error {
    return config.DB.Delete(&models.CartItem{}, id).Error
}

// Clear cart
func (r *CartRepository) Clear(userID uint) error {
    return config.DB.Where("user_id = ?", userID).
        Delete(&models.CartItem{}).Error
}