package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-crypto/internal/auth"
	"github.com/go-crypto/internal/model"
	"github.com/go-crypto/internal/store"
	"github.com/go-crypto/pkg/errorz"
	"github.com/go-crypto/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Register Create the Register handler
func Register(ctx context.Context, validate *validator.Validate, conn store.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.NewString())
		var register model.Register
		// Get the JSON body and decode into credentials
		err := json.NewDecoder(r.Body).Decode(&register)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = validate.Struct(register)
		if err != nil {
			utils.ValidateRequest(err, w)
		}
		c := auth.NewUserService(conn)
		err = c.RegisterUser(ctx, register)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorz.NewError("registration error", "check request body", err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user created successfully"))
	}
}

func ResetUser(ctx context.Context, validate *validator.Validate, conn store.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("x-request-id", uuid.NewString())

		var reset model.ResetPassword

		err := json.NewDecoder(r.Body).Decode(&reset)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = validate.Struct(reset)
		if err != nil {
			utils.ValidateRequest(err, w)
		}
		if reset.Password != reset.ConfirmPassword {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorz.NewError("password reset error", "password not matched", "password and confirm password is not same"))
			return
		}
		c := auth.NewUserService(conn)
		err = c.ResetPassword(ctx, reset)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorz.NewError("password reset error", "check request body", err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("user password reset successful"))
	}
}
