package usecases

import (
	"context"
	"fmt"

	"github.com/pedrolucaspalma/go-template/internal/application/interfaces"
)

func CreateUserUseCase(
	ctx context.Context,
	tm interfaces.TransactionManager,
	name string,
) error {

	err := tm.WithTx(ctx, func(uow interfaces.UnitOfWork) error {
		userRepo := uow.MakeUserRepository()
		_, err := userRepo.FindByID(ctx, "anything really")
		if err != nil {
			return fmt.Errorf("getting user by id: %w", err)
		}
		return nil
	})
	return fmt.Errorf("from transaction: %w", err)

}
