package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// AccessTokenClaims ...
type TokenClaims struct {
	TokenType string `json:"token_type"`
	UserID    uint   `json:"uid"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         User   `json:"user"`
}

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      // User ID associated with the refresh token
	Token     string    `gorm:"unique"`
	ExpiresAt time.Time `gorm:"index"`
	Revoked   bool      `gorm:"index"`
}
