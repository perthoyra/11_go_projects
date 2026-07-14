package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func CreateDBConnection() error {
	dbUsername := "root"
	dbPassword := "root"
	dbName := "TodoItems"
	dbHost := "localhost"
	dbPort := "3306"

	// dbUsername := getEnv("DB_USERNAME", "root")
	// dbPassword := getEnv("DB_PASSWORD", "")
	// dbName := getEnv("DB_NAME", "TodoItems")
	// dbHost := getEnv("DB_HOST", "localhost")
	// dbPort := getEnv("DB_PORT", "3306")

	// Example of mysql dsn
	// "root:root@tcp(127.0.0.1:3306)/TodoItems?parseTime=true&loc=Local"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
	})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
		return err
	}

	db = database
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func init() {
	fmt.Println("Configuring connections...")
}
