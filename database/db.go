package db

import (
	"fmt"
	"os"
	"time"

	"github.com/eapache/go-resiliency/retrier"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {

	var err error
	dsn := fmt.Sprintf("host=%v user=%v	password=%v dbname=%v port=%v sslmode=disable  TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	r := retrier.New(retrier.ExponentialBackoff(10, 100*time.Millisecond), nil)

	err = r.Run(func() error {
		dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := dbc.DB()
			if err != nil {
				return err
			}

			// Set maximum idle and open connections for the underlying *sql.DB
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)

			db = dbc
		}
		return err
	})
	if err != nil {
		return err
	}

	return Migrate()
}

func GetDBCon() *gorm.DB {
	return db
}
