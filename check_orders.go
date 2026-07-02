//go:build ignore

package main

import (
	"fmt"
	"github.com/annddvaa/gin-firebase-backend/config"
	"github.com/annddvaa/gin-firebase-backend/models"
	"github.com/joho/godotenv"
	"encoding/json"
)

func main() {
	godotenv.Load()
	config.InitDatabase()

	var orders []models.Order
	config.DB.Preload("OrderItems.Product").Find(&orders)

	b, _ := json.MarshalIndent(orders, "", "  ")
	fmt.Println(string(b))
}
