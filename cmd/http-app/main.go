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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDb(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer database.GetPool(ctx).Close()

	infoRepo := postgres.NewInfo(database)
	infoService := info.NewService(infoRepo)

	auth, _, err := config.Read()

	app.RunHTTP(ctx, infoService, auth)
}
