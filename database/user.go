package db

import (
	"UserApp/model"

	"gorm.io/gorm"
)

// Create a new user
func CreateUser(user *model.UserCreate) (*model.User, error) {
	userToCreate := model.User{
		Name:     user.Name,
		Username: user.Username,
		Surname:  user.Surname,
		Phone:    user.Phone,
		Address:  user.Address,
		Status:   model.UserStatusActive,
	}

	db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Create(&userToCreate).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.UserCredential{
			UserID:         userToCreate.ID,
			HashedPassword: user.Password,
		}).Error; err != nil {
			return err
		}

		return nil
	})

	return &userToCreate, nil
}

// Retrieve a user by ID
func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update a user's information
func UpdateUser(user *model.User) error {
	return db.Save(user).Error
}

// Delete a user by ID
func DeleteUser(id uint) error {
	return db.Delete(&model.User{}, id).Error
}

func GetUserByUserName(username string) (*model.User, error) {

	user := &model.User{}
	if err := db.Where("username = ?", username).First(user).Error; err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func GetUserCredentials(id uint) (*model.UserCredential, error) {
	var userCredential model.UserCredential
	err := db.Where("user_id = ?", id).First(&userCredential).Error
	if err != nil {
		return nil, err
	}
	return &userCredential, nil
}

// Get Users retrieves all the users from the database
func GetUsers() (*[]model.User, error) {
	var users []model.User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}
