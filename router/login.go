package router

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-crypto/internal/auth"
	"github.com/go-crypto/internal/model"
	"github.com/go-crypto/internal/store"
	"github.com/go-crypto/pkg/errorz"
	"github.com/go-crypto/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var jwtKey = []byte("my_secret_key")

// Signin Create the Signin handler
func Signin(ctx context.Context, validate *validator.Validate, conn store.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.NewString())
		var creds model.Login
		// Get the JSON body and decode into credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = validate.Struct(creds)
		if err != nil {
			utils.ValidateRequest(err, w)
		}
		// Get the expected password from our in memory map
		c := auth.NewUserService(conn)
		err = c.LoginUser(ctx, creds)

		// If a password exists for the given user
		// AND, if it is the same as the password we received, the we can move ahead
		// if NOT, then we return an "Unauthorized" status
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(errorz.NewError("login error", "login failed", err.Error()))
			return
		}

		// Declare the expiration time of the token
		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(5 * time.Minute)
		// Create the JWT claims, which includes the username and expiry time
		claims := &model.Claims{
			Email: creds.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "https://gocrypto.com",
				Subject:   "go-crypto",
				Audience:  nil,
				ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(5 * time.Minute)},
				NotBefore: &jwt.NumericDate{Time: time.Now()},
				IssuedAt:  &jwt.NumericDate{Time: time.Now()},
				ID:        uuid.NewString(),
			},
		}

		// Declare the token with the algorithm used for signing, and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Create the JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Finally, we set the client cookie for "token" as the JWT we just generated
		// we also set an expiry time which is the same as the token itself
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	}
}

func Logout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "token",
			MaxAge: -1,
		})
	}
}
