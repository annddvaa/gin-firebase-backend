//go:build ignore

package main

import (
	"fmt"
	"github.com/annddvaa/gin-firebase-backend/config"
	"github.com/annddvaa/gin-firebase-backend/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.InitDatabase()
	var products []models.Product
	config.DB.Where("name LIKE ?", "%iPhone 14%").Find(&products)
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s\n", p.ID, p.Name)
	}
}
