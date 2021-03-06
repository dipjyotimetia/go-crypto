package utils

import (
	"errors"
	"net/http"

	"github.com/go-crypto/internal/model"
	"github.com/go-crypto/pkg/errorz"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// ValidateRequest
// this check is only needed when your code could produce
// an invalid value for validation such as interface with nil
// value most including myself do not usually have code like this.
func ValidateRequest(err error, w http.ResponseWriter) {
	var _t0 *validator.InvalidValidationError
	if ok := errors.Is(err, _t0); ok {
		log.Fatal(err.Error())
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(errorz.NewError("validation error", "check request body", err.Error()))
}

// ValidatePasswordReset validating password reset
func ValidatePasswordReset(resetPassword model.ResetPassword) (bool, string) {
	if len(resetPassword.Password) < 4 {
		return false, "Invalid password, password should be more than 4 characters"
	}
	if resetPassword.Password != resetPassword.ConfirmPassword {
		return false, "Password reset failed, passwords must match"
	}
	return true, "Password validated successfully"
}
