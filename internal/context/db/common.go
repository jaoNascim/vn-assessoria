package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	connStr := os.Getenv("connectionString")

	return connStr
}

func connect() *gorm.DB {
	db, err := sql.Open("postgres", getConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return gormDB
}
