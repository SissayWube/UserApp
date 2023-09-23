package db

import "UserApp/model"

func Migrate() error {
	db := GetDBCon()

	dbc, err := db.DB()
	if err != nil {
		return err
	}

	defer dbc.Close()	

	// Create users table
	if !db.Migrator().HasTable(&model.User{}) {
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			return err
		}
	}

	if !db.Migrator().HasTable(&model.UserCredential{}) {
		err := db.AutoMigrate(&model.UserCredential{})
		if err != nil {
			return err
		}
	}
	return nil
}
