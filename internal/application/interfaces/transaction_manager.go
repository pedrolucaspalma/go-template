package interfaces

import "context"

type Transaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type UnitOfWork interface {
	Transaction

	MakeUserRepository() UserRepository
}

type TransactionManager interface {
	WithTx(ctx context.Context, f func(uow UnitOfWork) error) error
}
