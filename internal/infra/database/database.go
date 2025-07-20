package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnection(ctx context.Context) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	return dbpool, err
}
