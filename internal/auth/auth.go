package auth

import (
	"context"

	"github.com/go-crypto/internal/model"
	"github.com/go-crypto/internal/store"
)

type UserService interface {
	RegisterUser(ctx context.Context, user model.Register) error
	LoginUser(ctx context.Context, user model.Login) error
}

type userService struct {
	userRepository store.CryptoService
}

func (u userService) LoginUser(ctx context.Context, user model.Login) error {
	return u.userRepository.LoginUser(ctx, user)
}

func (u userService) RegisterUser(ctx context.Context, user model.Register) error {
	return u.userRepository.RegisterUser(ctx, user)
}

func NewUserService(store store.CryptoService) UserService {
	return &userService{userRepository: store}
}
