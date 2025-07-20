package interfaces

import (
	"context"

	"github.com/pedrolucaspalma/go-template/internal/domain"
)

type UserRepository interface {
	FindByID(ctx context.Context, ID string) (domain.User, error)
}
