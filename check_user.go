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
	var u []models.User
	config.DB.Find(&u)
	fmt.Printf("Users: %+v\n", u)
}
