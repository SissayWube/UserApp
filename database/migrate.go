package db

import (
	"UserApp/model"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Migrate() error {

	// Create users table
	if !db.Migrator().HasTable(&model.User{}) {

		// Create the users table
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			return err
		}

		// Create the credentials tables
		if !db.Migrator().HasTable(&model.UserCredential{}) {
			err := db.AutoMigrate(&model.UserCredential{})
			if err != nil {
				return err
			}
		}

		// Create the default admin user if not exists
		var count int64
		db.Model(&model.User{}).Where("username = ?", "admin").Count(&count)
		if count == 0 {
			// Create default admin user credential
			adminPass, found := os.LookupEnv("ADMIN_PASS")
			if !found {
				return model.ErrADMIN_PASSNotSet
			}

			// Hash password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPass), 14)
			if err != nil {
				return err
			}

			admin := model.UserCreate{
				Username: "admin",
				Password: string(hashedPassword),
			}
			_, err = CreateUser(&admin)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
