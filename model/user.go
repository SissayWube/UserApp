package model

import (
	"database/sql"
	"time"
)

// User represents the user entity.
type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Status   string `json:"status"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

// UserCredential represents the users credentials.
type UserCredential struct {
	ID                 uint   `json:"id"`
	UserID             uint   `json:"user_id"`
	HashedPassword     string `json:"hashed_password"`
	RefreshTokenStatus string `json:"refresh_token_status"`
}

// UserCreate Request represents the request body for creating a new user.
type UserCreate struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
