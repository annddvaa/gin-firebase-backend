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
	config.DB.Find(&products)

	categories := make(map[string]int)
	for _, p := range products {
		categories[p.Category]++
	}

	fmt.Println("--- Categories ---")
	for cat, count := range categories {
		fmt.Printf("- %s: %d\n", cat, count)
	}

	fmt.Println("\n--- Example Products ---")
	for i, p := range products {
		if i < 10 {
			fmt.Printf("ID: %d | Name: %s | Cat: %s | Img: %s\n", p.ID, p.Name, p.Category, p.ImageURL)
		}
	}
}
