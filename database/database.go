package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ThePratikSah/silent-flag/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	_port := config.Config("PORT")
	port, err := strconv.ParseInt(_port, 10, 32)
	if err != nil {
		port = 5432 // Default port if parsing fails
	}
	host := config.Config("HOSTNAME")
	user := config.Config("POSTGRES_USER")
	password := config.Config("POSTGRES_PASSWORD")
	dbname := config.Config("POSTGRES_DB")
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

	DB = Dbinstance{
		Db: db,
	}
}
