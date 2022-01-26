package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-crypto/internal/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Store) RegisterUser(ctx context.Context, user model.Register) error {
	_, err := s.Collection("users").Doc(user.Email).Get(ctx)
	if err != nil && status.Code(err) == codes.NotFound {
		_, err = s.Collection("users").Doc(user.Email).Set(ctx, model.Register{
			Name:      user.Name,
			Password:  createHashedPassword(user.Password),
			Email:     user.Email,
			CreatedAt: time.Now().UTC().Format(time.RFC3339),
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		})
		if err != nil {
			log.Error("error while registering user", err.Error())
			return fmt.Errorf("user already exists")
		}
	} else {
		log.Errorf("user already exists")
		return fmt.Errorf("user already exists")
	}
	return nil
}

func (s Store) LoginUser(ctx context.Context, user model.Login) error {
	var userData model.Register
	data, err := s.Collection("users").Doc(user.Email).Get(ctx)
	if err != nil {
		log.Error("error while getting user info", err.Error())
		return err
	}
	if err = data.DataTo(&userData); err != nil {
		log.Error("error while updating price info", err.Error())
		return err
	}
	err = CheckPasswordHash(user.Password, userData.Password)
	if err != nil {
		log.Error("error wrong password", err.Error())
		return err
	}
	return nil
}

// CreateHashedPassword created password hashed
func createHashedPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

// CheckPasswordHash compares hash with password
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
