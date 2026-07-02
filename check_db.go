//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/annddvaa/gin-firebase-backend/config"
	"github.com/annddvaa/gin-firebase-backend/models"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config.InitDatabase()

	var orders []models.Order
	err := config.DB.Preload("OrderItems.Product").Find(&orders).Error
	if err != nil {
		fmt.Println("GORM Error:", err)
	}
	fmt.Printf("Total orders: %d\n", len(orders))
	b, _ := json.MarshalIndent(orders, "", "  ")
	fmt.Println(string(b))

	var products []models.Product
	config.DB.Find(&products)
	fmt.Printf("Total products: %d\n", len(products))
	var pIDs []uint
	for _, p := range products {
		pIDs = append(pIDs, p.ID)
	}
	fmt.Printf("Product IDs: %v\n", pIDs)
}
