package db

import (
	"e-commerce-listing/database/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var primaryDB *gorm.DB
var replicaDB *gorm.DB

type DBConn struct {
	Primary *gorm.DB
	Replica *gorm.DB
}

func New() *DBConn {
	return &DBConn{
		Primary: primaryDB,
		Replica: replicaDB,
	}
}

func InitializeDB() {

	primaryConnectionString := "host=host user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	replicaConnectionString := "host=host user=postgres password=password dbname=postgres port=5432 sslmode=disable"

	primaryDBInstance, err := gorm.Open(postgres.Open(primaryConnectionString), &gorm.Config{})
	if err != nil {
		log.Println("Unable to connect to primary database:", err)
		return
	}

	primaryDB = primaryDBInstance

	log.Println("primary database connected successfully!")

	replicaDBInstance, err := gorm.Open(postgres.Open(replicaConnectionString), &gorm.Config{})
	if err != nil {
		log.Println("Unable to connect to replica database:", err)
		return
	}

	replicaDB = replicaDBInstance

	log.Println("replica database connected successfully!")

	AutoMigrateTables()
}

func AutoMigrateTables() {

	if primaryDB == nil {
		log.Println("Database connection is nil, skipping migration")
		return
	}

	primaryDB.AutoMigrate(&models.Product{})
}
