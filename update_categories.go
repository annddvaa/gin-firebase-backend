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

	// Update "Layar" to "Layar LCD"
	res1 := config.DB.Model(&models.Product{}).Where("category = ?", "Layar").Update("category", "Layar LCD")
	fmt.Printf("Updated %d Layar products\n", res1.RowsAffected)

	// Update "Lainnya" to "Komponen HP Lainnya"
	res2 := config.DB.Model(&models.Product{}).Where("category = ?", "Lainnya").Update("category", "Komponen HP Lainnya")
	fmt.Printf("Updated %d Lainnya products\n", res2.RowsAffected)

	fmt.Println("Categories updated successfully")
}
