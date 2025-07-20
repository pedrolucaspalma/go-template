package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnection(ctx context.Context) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("db connection: %w", err)
	}

	return dbpool, err
}
