package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pedrolucaspalma/go-template/internal/domain"
)

type userRepository struct {
	tx pgx.Tx
}

func NewUserRepository(tx pgx.Tx) userRepository {
	return userRepository{
		tx: tx,
	}
}

func (r userRepository) FindByID(ctx context.Context, ID string) (domain.User, error) {
	// TODO finish this method, for now it is a mock
	return domain.User{}, nil
}
