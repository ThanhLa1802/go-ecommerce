package main

import (
	"go-ecommerce-backend-api/internal/initialize"
)

func main() {
	// r := routers.NewRouter()

	// if err := r.Run(); err != nil {
	// 	log.Fatalf("failed to run server: %v", err)
	// }
	initialize.Run()
}

// Handler function phải nằm ngoài main
