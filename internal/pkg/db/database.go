package db

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Database struct represents a connection to a PostgreSQL database pool
type Database struct {
	cluster *pgxpool.Pool
}

// newDatabase creates a new Database instance with the given pgxpool.Pool
func newDatabase(cluster *pgxpool.Pool) *Database {
	return &Database{cluster: cluster}
}

// GetPool returns the underlying pgxpool.Pool instance
func (db Database) GetPool(_ context.Context) *pgxpool.Pool {
	return db.cluster
}

// Get executes a query and scans the result into the destination struct
func (db Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.cluster, dest, query, args...)
}

// Select executes a query and scans all the rows into the destination slice
func (db Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.cluster, dest, query, args...)
}

// Exec executes a query without returning any rows, returning a command tag
func (db Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.cluster.Exec(ctx, query, args...)
}

// ExecQueryRow executes a query that is expected to return at most one row
func (db Database) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.cluster.QueryRow(ctx, query, args...)
}
