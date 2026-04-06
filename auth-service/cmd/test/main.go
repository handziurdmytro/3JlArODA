package main

import (
	"fmt"
	"log"

	"github.com/handziurdmytro/3JlArODA/auth-service/internal/crypto"
)

func main() {
	client := crypto.NewClient("localhost:3030")
	defer client.Close()

	hash, err := client.HashPassword("testpassword")
	if err != nil {
		log.Fatalf("HashPassword failed: %v", err)
	}
	fmt.Println("Hash:", hash)

	valid, err := client.VerifyPassword("testpassword", hash)
	if err != nil {
		log.Fatalf("VerifyPassword failed: %v", err)
	}
	fmt.Println("Valid:", valid)
}
