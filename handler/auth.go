package handler

import (
	db "UserApp/database"
	"UserApp/model"
	"UserApp/utility"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LogIn(c echo.Context) error {
	// Parse the request body into a Login struct
	login := new(model.Login)
	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusBadRequest, &model.ErrorResponse{Message: err.Error()})
	}

	// Validate the login data
	if err := model.Validate.Struct(login); err != nil {
		return c.JSON(http.StatusBadRequest, &model.ErrorResponse{Message: err.Error()})
	}

	// Attempt to log in and retrieve tokens
	tokens, err := utility.Login(login)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &model.ErrorResponse{Message: "Invalid Credentials"})
	}

	// Return a successful response with tokens
	return c.JSON(http.StatusOK, &model.Response{Data: tokens})
}

func LogOut(c echo.Context) error {
	id, _ := strconv.Atoi(c.Request().Header.Get("id"))
	// revoke the refresh token
	db.RevokeRefreshToken(uint(id))

	return c.JSON(http.StatusOK, nil)
}

func RefreshAccessToken(c echo.Context) error {
	// Parse the request to get the refresh token
	cookie, err := c.Cookie("accesstoken")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.ErrorResponse{Message: err.Error()})
	}

	if cookie.Value == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	// Verify the refresh token and get the associated user
	user, err := utility.ValidateRefreshToken(cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &model.ErrorResponse{Message: model.ErrInvalidRefreshToken})
	}

	// Generate a new access token for the user
	accessToken, err := utility.GenerateAccessToken(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.ErrorResponse{Message: model.ErrGeneratingAccessToken})
	}

	// Return the new access token
	response := model.Response{Data: map[string]string{"access_token": accessToken}}
	return c.JSON(http.StatusOK, response)
}
