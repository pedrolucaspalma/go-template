package interfaces

import "github.com/pedrolucaspalma/go-template/internal/domain"

type UserRepository interface {
	FindByID(ID string) domain.User
}
