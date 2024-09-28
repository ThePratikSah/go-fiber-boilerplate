package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ThePratikSah/silent-flag/config"
	"github.com/ThePratikSah/silent-flag/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// ConnectDb establishes a connection to the PostgreSQL database.
//
// It retrieves the database connection parameters from environment variables,
// constructs a DSN string, and uses it to open a GORM database connection.
// If the connection fails, it logs the error and exits the application.
// Otherwise, it sets up the database logger and assigns the connection to the DB variable.
//
// No parameters.
// No return values.
func ConnectDb() {
	_port := config.Config("PORT", "5432")
	port, err := strconv.ParseInt(_port, 10, 32)
	if err != nil {
		port = 5432 // Default port if parsing fails
	}
	host := config.Config("HOSTNAME", "localhost")
	user := config.Config("POSTGRES_USER", "postgres")
	password := config.Config("POSTGRES_PASSWORD", "")
	dbname := config.Config("POSTGRES_DB", "default_db")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	db.AutoMigrate(&model.User{})

	DB = Dbinstance{
		Db: db,
	}
}
