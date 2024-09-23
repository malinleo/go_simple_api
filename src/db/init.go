package db

import (
	"fmt"
	"simple_api/src/db/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Initialize() error {
	dsn := "host=localhost port=5432 user=go_simple_api password=go_simple_api dbname=go_simple_api sslmode=disable"
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connection establishing failed: %w", err)
	}
	migrationErr := dbConn.AutoMigrate(&models.User{})
	if migrationErr != nil {
		return fmt.Errorf("migrations have failed: %w", migrationErr)
	}
	sqlDB, sqlErr := dbConn.DB()
	if sqlErr != nil {
		return fmt.Errorf("connection establishing failed: %w", sqlErr)
	}
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return fmt.Errorf("some generic error: %s", "substituted value")
}