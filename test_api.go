//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Buat token JWT untuk user ID 4
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": float64(4),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("12345678901234567890123456789012"))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/v1/orders", nil)
	req.Header.Add("Authorization", "Bearer "+tokenString)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))
}
