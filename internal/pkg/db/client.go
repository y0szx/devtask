package db

import (
	"context"
	"devtask/internal/config"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"
)

func NewDb(ctx context.Context) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, generateDsn())
	if err != nil {
		return nil, err
	}
	return newDatabase(pool), nil
}

func generateDsn() string {
	_, db, err := config.Read()
	if err != nil {
		return ""
	}

	port, _ := strconv.ParseInt(db.Port, 10, 64)
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.Host, port, db.User, db.Password, db.Name)
}
