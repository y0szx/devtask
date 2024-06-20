package main

import (
	"context"
	"devtask/internal/app"
	"devtask/internal/config"
	"devtask/internal/pkg/db"
	"devtask/internal/service/info"
	"devtask/internal/storage/postgres"
	"log"
)

func main() {
	// Create a context with cancellation capability
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancellation of context when main function exits

	// Initialize the database connection
	database, err := db.NewDb(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer database.GetPool(ctx).Close() // Close the database connection pool when main function exits

	// Create a new repository instance using PostgreSQL storage
	infoRepo := postgres.NewInfo(database)

	// Create a new service instance using the repository
	infoService := info.NewService(infoRepo)

	// Read authentication credentials from configuration
	auth, _, err := config.Read()

	// Run the HTTP server with the initialized service and authentication credentials
	app.RunHTTP(ctx, infoService, auth)
}
