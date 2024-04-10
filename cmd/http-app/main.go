package main

import (
	"context"
	"devtask/internal/app"
	"devtask/internal/config"
	"devtask/internal/pkg/db"
	"devtask/internal/service/pvz"
	"devtask/internal/storage/postgres"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDb(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer database.GetPool(ctx).Close()

	pvzsRepo := postgres.NewPVZs(database)
	pvzService := pvz.NewService(pvzsRepo)

	auth, _, err := config.Read()

	app.RunHTTP(ctx, pvzService, auth)
}
