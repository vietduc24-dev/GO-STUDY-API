package main

import (
	"go-ecommerce-backend-api/internal/initialize"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	initialize.Run()
}
