package handler

import (
	"UserApp/model"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.Use(AuthMiddleware)
	api := e.Group("/api")
	users := api.Group("/users")
	auth := api.Group("/auth")

	users.POST("", CreateUser)
	users.GET("", GetUsers)

	user := users.Group("/:id")

	user.GET("", GetUser)
	user.PATCH("", UpdateUser)
	user.DELETE("", DeleteUser)

	auth.POST("/refresh", RefreshAccessToken)
	auth.POST("/login", LogIn)
	auth.POST("/logout", LogOut)

	port, found := os.LookupEnv("PORT")
	if !found {
		port = model.DefaultPort
	}
	e.Logger.Fatal(e.Start(":" + port))

}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().URL.Path == "/api/auth/login" {
			// Skip authentication for this route
			return next(c)
		}
		// get access token from request
		cookie, err := c.Cookie("accesstoken")
		if err != nil {
			return c.JSON(http.StatusBadRequest, &model.ErrorResponse{Message: err.Error()})
		}

		if cookie.Value == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		accessToken := cookie.Value
		token, err := jwt.ParseWithClaims(accessToken, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%v: %v", model.ErrUnexpectedSigningMethod, token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			er := strings.Split(err.Error(), "by")
			return c.JSON(http.StatusBadRequest, &model.ErrorResponse{Message: model.ErrInvalidAccessToken + er[0]})
		}

		if !token.Valid {
			return c.JSON(http.StatusBadRequest, model.ErrInvalidAccessToken)
		}

		claims, isClaims := token.Claims.(*model.TokenClaims)
		if !isClaims {
			return c.JSON(http.StatusUnauthorized, &model.ErrorResponse{Message: model.ErrInvalidAccessToken})
		}

		if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
			// if the user's token has been expired respond with access token expired message so that the client can refresh the access token
			if err != nil {
				return c.JSON(http.StatusUnauthorized, &model.ErrorResponse{Message: model.ErrInvalidAccessToken})
			}
		}

		c.Request().Header.Add("id", strconv.Itoa(int(claims.UserID)))

		return next(c)
	}
}
