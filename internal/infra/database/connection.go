package database

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pedrolucaspalma/go-template/config"
)

func GetConnection(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)
	dbpool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("db connection: %w", err)
	}

	return dbpool, err
}

func GetTestConnection(ctx context.Context) (*pgxpool.Pool, error) {
	type testConnectionCfg struct {
		DB_USERNAME string
		DB_PASSWORD string
		DB_HOST     string
		DB_PORT     string
		DB_NAME     string
	}
	cfg := testConnectionCfg{
		DB_USERNAME: "postgres",
		DB_PASSWORD: "postgres",
		DB_HOST:     "localhost",
		DB_NAME:     fmt.Sprintf("test_database_%s", getPseudoUUID()),
		DB_PORT:     "3001",
	}
	url := fmt.Sprintf("postgres://test_user:%s@%s:%s/%s",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)
	dbpool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("creating pool: %w", err)
	}
	executeMigrations(ctx, dbpool)
	executeSeeds(ctx, dbpool)

	return dbpool, err
}

func executeMigrations(ctx context.Context, poll *pgxpool.Pool) error {
	// TODO
	return nil
}

func executeSeeds(ctx context.Context, poll *pgxpool.Pool) error {
	// TODO
	return nil
}

/*
This function doesn't generate actual UUIDs with spec-compliance, it is merely a random string that resembles an UUID

It should not be used on places that really need a UUID (such as database row identifiers)

It is only used to generate test database schema identifiers on the function database.NewTestMongoConnection()
*/
func getPseudoUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic("error when trying to generate identifier: \n" + err.Error())
	}

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}
