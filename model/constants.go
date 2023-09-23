package model

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

const (
	ErrUnexpectedSigningMethod = "unexpected signing method"
	ErrInvalidAccessToken      = "Invalid access token"
	ErrInvalidRefreshToken     = "Invalid refresh token"
	ErrInvalidCredentials      = "Invalid Credentials"
	ErrGeneratingAccessToken   = "Error generating access token"
	ErrFetchingUser            = "Error fetching User"
	ErrCreatingUser            = "Error creating User"
	ErrInvalidUserID           = "Error invalid user id"
	ErrInvalidRequest          = "Error invalid Request"
	ErrUserNotFound            = "Error use not found"
	ErrUpdatingUser            = "Error updating user"
	ErrDeletingUser            = "Error deleting user"
)

var ErrIncorrectPassword = errors.New("error incorrect password")
var ErrInvalidToken = errors.New("error invalid token")
var ErrInvalidClaims = errors.New("error invalid claims")

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// Response represents an error response
type Response struct {
	Data interface{} `json:"message"`
}
