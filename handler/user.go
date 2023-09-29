package handler

import (
	db "UserApp/database"
	"UserApp/model"
	"UserApp/utility"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateUser handles POST requests to /users and creates a new user record
func CreateUser(c echo.Context) error {
	// Get user from request body
	user := new(model.UserCreate)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: model.ErrInvalidRequest})
	}

	// Validate user data
	if err := model.Validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}

	// Hash password
	hashedPassword, err := utility.HashPassword(user.Password)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)

	}
	user.Password = hashedPassword

	newUser := new(model.User)
	// Save user to database
	if newUser, err = db.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: model.ErrCreatingUser})
	}

	// Return success response
	return c.JSON(http.StatusCreated, newUser)
}

// GetUser handles GET requests to /users/:id
func GetUser(c echo.Context) error {
	// Get id from path parameter
	id := c.Param("id")
	// convert id to uint
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: model.ErrInvalidUserID})
	}
	// Get user from database
	user, err := db.GetUserByID(uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: model.ErrUserNotFound})
		}
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: model.ErrFetchingUser})
	}

	// Return response
	return c.JSON(http.StatusOK, user)
}

// UpdateUser handles PUT requests to /users/:id
func UpdateUser(c echo.Context) error {
	// Get id from path parameter
	id := c.Param("id")
	// convert id to uint
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: model.ErrInvalidUserID})
	}

	// Get user from the database
	_, err = db.GetUserByID(uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: model.ErrUserNotFound})
		}
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: model.ErrFetchingUser})
	}

	// Get updated user from request body
	updatedUser := new(model.User)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: model.ErrInvalidRequest})
	}

	// Validate updated user data
	if err := model.Validate.Struct(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
	}

	// Update user in database
	if err := db.UpdateUser(updatedUser); err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: model.ErrUpdatingUser})
	}

	// Return response
	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles DELETE requests to /users/:id
func DeleteUser(c echo.Context) error {
	// Get id from path parameter
	id := c.Param("id")
	// convert id to uint
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: model.ErrInvalidUserID})
	}
	// Delete user from database
	if err := db.DeleteUser(uint(userID)); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, model.ErrorResponse{Message: model.ErrUserNotFound})
		}
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: model.ErrDeletingUser})
	}

	// Return response
	return c.NoContent(http.StatusNoContent)
}

// Get users fetches all users from the database
func GetUsers(c echo.Context) error {
	// Get all users
	users, err := db.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: model.ErrFetchingUsers})
	}

	// Return response
	return c.JSON(http.StatusOK, users)

}
