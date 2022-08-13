package model

import (
	"github.com/golang-jwt/jwt/v4"
)

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Register struct
type Register struct {
	Name      string        `json:"name" firestore:"name" validate:"required"`
	Password  string        `json:"password" firestore:"password" validate:"required"`
	Email     string        `json:"email" firestore:"email" validate:"required"`
	CreatedAt string        `json:"created_at" firestore:"created_at,omitempty"`
	UpdatedAt string        `json:"updated_at" firestore:"updated_at,omitempty"`
	Status    AccountStatus `json:"status" firestore:"status,omitempty"`
}

type AccountStatus struct {
	Locked     bool    `json:"locked" firestore:"locked,omitempty"`
	LockedAt   *string `json:"lockedAt" firestore:"lockedAt,omitempty"`
	LoginCount int     `json:"loginCount" firestore:"loginCount,omitempty"`
}

type ResetPassword struct {
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// Login struct
type Login struct {
	Password string        `json:"password" firestore:"password" validate:"required"`
	Email    string        `json:"email" firestore:"email" validate:"required,email"`
	Status   AccountStatus `json:"status" firestore:"status,omitempty"`
}

type CreateReset struct {
	Email string `json:"email" validate:"required,email"`
}
