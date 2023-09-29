package model

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

const (
	DefaultPort      = "8000"
	UserStatusActive = "Active"

	ErrUnexpectedSigningMethod = "Unexpected signing method"
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
	ErrFetchingUsers           = "Error fetch user"
)

var ErrIncorrectPassword = errors.New("error incorrect password")
var ErrInvalidToken = errors.New("error invalid token")
var ErrInvalidClaims = errors.New("error invalid claims")
var ErrADMIN_PASSNotSet = errors.New("error ADMIN_PASS environment variable not set")

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
