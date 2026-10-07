package main

import "go-ecommerce-backend-api/internal/routers"

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := routers.NewRouter()

	r.Run(":8002")
}
