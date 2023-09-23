package utility

import (
	db "UserApp/database"
	"UserApp/model"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) string {
	// Hash password using bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateAccessToken generates an access token from the given payload
func GenerateAccessToken(user model.User) (string, error) {

	// Create the token
	claims := model.TokenClaims{
		UserID:    user.ID,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			IssuedAt:  jwt.NewNumericDate(time.Now())},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

// GenerateRefreshToken generates an access token from the given payload
func GenerateRefreshToken(user model.User) (string, error) {

	// Create the token
	claims := model.TokenClaims{
		UserID:    user.ID,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 720)),
			IssuedAt:  jwt.NewNumericDate(time.Now())},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

// ValidateRefreshToken checks if a token is valid or not
func ValidateRefreshToken(tokenString string) (*model.User, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%v: %v", model.ErrUnexpectedSigningMethod, token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {

		claims, validClaims := token.Claims.(*model.TokenClaims)
		if !validClaims {
			return nil, model.ErrInvalidClaims
		}
		// Check if the token is a refresh token
		if claims.TokenType != "refresh" {
			return nil, model.ErrInvalidToken
		}
		return db.GetUserByID(claims.UserID)
	}

	return nil, model.ErrInvalidToken
}

// Login validates the login credentials of a user and returns an access toke
func Login(login *model.Login) (*model.TokenResponse, error) {
	// Get the user by its usernanme
	user, err := db.GetUserByUserName(login.Username)
	if err != nil {
		return nil, err
	}

	// Get the users credentials to help validate the password
	userCredentials, err := db.GetUserCredentials(user.ID)
	if err != nil {
		return nil, err
	}

	// Validate Password
	if valid := CheckPasswordHash(userCredentials.HashedPassword, login.Password); !valid {
		err = model.ErrIncorrectPassword
		return nil, err
	}

	// Issue access token
	accessToken, err := GenerateAccessToken(*user)
	if err != nil {
		return nil, err
	}

	// Issue refresh token
	refreshToken, err := GenerateRefreshToken(*user)
	if err != nil {
		return nil, err
	}

	return &model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}, nil

}
