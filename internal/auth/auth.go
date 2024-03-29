package auth

import (
	"context"

	"github.com/go-crypto/internal/model"
	"github.com/go-crypto/internal/store"
)

type UserService interface {
	RegisterUser(ctx context.Context, user model.Register) error
	LoginUser(ctx context.Context, user model.Login) error
	ResetPassword(ctx context.Context, reset model.ResetPassword) error
}

type userService struct {
	userRepository store.UserService
}

func (u userService) ResetPassword(ctx context.Context, reset model.ResetPassword) error {
	return u.userRepository.ResetPassword(ctx, reset)
}

func (u userService) LoginUser(ctx context.Context, user model.Login) error {
	return u.userRepository.LoginUser(ctx, user)
}

func (u userService) RegisterUser(ctx context.Context, user model.Register) error {
	return u.userRepository.RegisterUser(ctx, user)
}

func NewUserService(store store.UserService) UserService {
	return &userService{userRepository: store}
}
