package services

import (
	"errors"
	"github.com/annddvaa/gin-firebase-backend/models"
	"github.com/annddvaa/gin-firebase-backend/repositories"
	"gorm.io/gorm"
)

type CartService struct {
	repo *repositories.CartRepository
}

func NewCartService() *CartService {
	return &CartService{
		repo: repositories.NewCartRepository(),
	}
}

func (s *CartService) GetCart(userID uint) ([]models.CartItem, error) {
	return s.repo.FindByUser(userID)
}

func (s *CartService) Add(userID, productID uint, qty int) error {
	if qty < 1 {
		return errors.New("qty minimal 1")
	}

	item, err := s.repo.FindItem(userID, productID)

	if err == nil {
		return s.repo.UpdateQty(item.ID, item.Quantity+qty)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	newItem := &models.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  qty,
	}

	return s.repo.Create(newItem)
}

func (s *CartService) UpdateQty(id uint, qty int) error {
	if qty <= 0 {
		return s.repo.Delete(id)
	}

	return s.repo.UpdateQty(id, qty)
}

func (s *CartService) Remove(id uint) error {
	return s.repo.Delete(id)
}

func (s *CartService) Clear(userID uint) error {
	return s.repo.Clear(userID)
}