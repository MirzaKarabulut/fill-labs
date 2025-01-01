package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() error {
	var err error
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}
	fmt.Println("Connected to the database")
	return nil
}

func AutoMigrate() error {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("could not migrate the database: %w", err)
	}
	fmt.Println("Database migrated successfully")
	return nil
}
