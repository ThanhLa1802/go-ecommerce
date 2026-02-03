package main

import (
	"go-ecommerce-backend-api/internal/routers"
	"log"
)

func main() {
	r := routers.NewRouter()

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

// Handler function phải nằm ngoài main
