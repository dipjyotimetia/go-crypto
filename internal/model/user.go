package model

import (
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Credentials Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Register struct
type Register struct {
	Name      string `json:"name" firestore:"name" validate:"required"`
	Password  string `json:"password" firestore:"password" validate:"required"`
	Email     string `json:"email" firestore:"email" validate:"required"`
	CreatedAt string `json:"created_at" firestore:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at" firestore:"updated_at,omitempty"`
}

type ResetPassword struct {
	ID              int    `json:"id" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// Login struct
type Login struct {
	Password string `json:"password" firestore:"password" validate:"required"`
	Email    string `json:"email" firestore:"email" validate:"required,email"`
}

type CreateReset struct {
	Email string `json:"email" validate:"required,email"`
}

// HashPassword hashes user password
func HashPassword(user *Register) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
}

// CheckPasswordHash compares hash with password
func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
