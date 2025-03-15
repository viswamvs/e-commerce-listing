package db

import (
	"e-commerce-listing/database/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type DBConn struct {
	*gorm.DB
}

func New() *DBConn {
	return &DBConn{
		DB: DB,
	}
}

func InitializeDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	connectionString := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode

	dbInstance, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println("Unable to connect to database:", err)
		return
	}

	DB = dbInstance
	log.Println("Database connected successfully!")

	AutoMigrateTables()
}

func AutoMigrateTables() {
	if DB == nil {
		log.Println("Database connection is nil, skipping migration")
		return
	}
	DB.AutoMigrate(&models.Product{})
}
