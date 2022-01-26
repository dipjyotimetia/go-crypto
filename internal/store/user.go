package store

import (
	"context"
	"time"

	"github.com/go-crypto/internal/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (s Store) RegisterUser(ctx context.Context, user model.Register) {
	_, err := s.Collection("users").Doc(user.Email).Set(ctx, model.Register{
		Name:      user.Name,
		Password:  createHashedPassword(user.Password),
		Email:     user.Email,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
		UpdatedAt: time.Now().UTC().Format(time.RFC3339),
	})
	if err != nil {
		log.Fatal("error while adding price history", err.Error())
	}
}

func (s Store) LoginUser(ctx context.Context, user model.Login) error {
	var userData model.Register
	data, err := s.Collection("users").Doc(user.Email).Get(ctx)
	if err != nil {
		log.Fatal("error while getting user info", err.Error())
		return err
	}
	if err = data.DataTo(&userData); err != nil {
		log.Fatal("error while updating price info", err.Error())
		return err
	}
	err = CheckPasswordHash(user.Password, userData.Password)
	if err != nil {
		log.Fatal("error wrong password", err.Error())
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
