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

	updates := map[string]string{
		"Layar":    "assets/images/services/layar.png",
		"Baterai":  "assets/images/services/baterai.png",
		"Kamera":   "assets/images/services/kamera.png",
		"Charging": "assets/images/services/charging.png",
		"Software": "assets/images/services/software.png",
		"Lainnya":  "assets/images/services/lainnya.png",
	}

	for category, url := range updates {
		res := config.DB.Model(&models.Product{}).Where("category = ?", category).Update("image_url", url)
		fmt.Printf("Updated %d products in category %s\n", res.RowsAffected, category)
	}
}
