package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pedrolucaspalma/go-template/internal/application/interfaces"
	"github.com/pedrolucaspalma/go-template/internal/infra/database/repositories"
)

type pgUoW struct {
	tx pgx.Tx
}

func (uow *pgUoW) Commit(ctx context.Context) error {
	return uow.tx.Commit(ctx)
}

func (uow *pgUoW) Rollback(ctx context.Context) error {
	return uow.tx.Rollback(ctx)
}

// Factory repository methods
func (uow *pgUoW) MakeUserRepository() interfaces.UserRepository {
	return repositories.NewUserRepository(uow.tx)
}

type transactionManager struct {
	conn *pgxpool.Conn
	opts pgx.TxOptions
}

func NewTransactionManager(conn *pgxpool.Conn, opts pgx.TxOptions) *transactionManager {
	return &transactionManager{conn: conn, opts: opts}
}

func (tm *transactionManager) WithTx(ctx context.Context, f func(uow interfaces.UnitOfWork) error) error {
	tx, err := tm.conn.BeginTx(ctx, tm.opts)

	pgUoW := pgUoW{tx: tx}

	err = f(&pgUoW)
	if err != nil {
		err := pgUoW.Rollback(ctx)
		return err
	}
	err = pgUoW.Commit(ctx)
	return err
}
