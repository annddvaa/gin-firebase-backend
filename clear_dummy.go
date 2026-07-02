//go:build ignore

package main

import (
	"fmt"
	"log"

	"github.com/annddvaa/gin-firebase-backend/config"
	"github.com/annddvaa/gin-firebase-backend/models"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	config.InitDatabase()

	// Find the dummy orders
	var dummyOrders []models.Order
	config.DB.Where("status IN ? AND total_price IN ?", []string{"menunggu", "dikerjakan", "selesai"}, []float64{850000, 350000, 150000}).Find(&dummyOrders)

	for _, o := range dummyOrders {
		fmt.Printf("Deleting dummy order ID %d...\n", o.ID)
		config.DB.Where("order_id = ?", o.ID).Delete(&models.OrderItem{})
		config.DB.Delete(&o)
	}

	fmt.Println("Dummy orders deleted successfully!")
}
