package main

import (
	"fmt"
	"github.com/lsendoya/handleBook/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func ConnectDB() (*gorm.DB, error) {
	dsn, err := makeDSN()
	if err != nil {
		return nil, fmt.Errorf("makeDSN(), Error creating DSN, %w", err)
	}

	db, errConn := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open(), Failed to connect to database, %w", errConn)
	}

	log.Println("Connection opened to database")
	return db, nil
}

func makeDSN() (string, error) {
	portStr := Config("DB_PORT")
	port, err := strconv.ParseUint(portStr, 10, 32)
	if err != nil {
		return "", fmt.Errorf("strconv.ParseUint(), error converting DB_PORT to uint, %w", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		Config("DB_HOST"), port, Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_NAME"), Config("DB_SSL_MODE"))

	return dsn, nil
}

func MigrateDB(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	models := []interface{}{
		&model.Book{},
		&model.User{},
		&model.Loan{},
	}

	for _, mdl := range models {
		if err := db.AutoMigrate(mdl); err != nil {
			log.Printf("Error migrating model %T: %v", mdl, err)
			return fmt.Errorf("error migrating model %T: %w", mdl, err)
		}
	}

	log.Println("Database migration completed successfully")
	return nil
}
