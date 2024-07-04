package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Realtys struct {
	District     string
	Description  string
	Price        int
	Mode         string
	Type         string
	LandArea     int
	BuildingArea int
	CreateDate   time.Time
	UpdateDate   time.Time
}

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

func (r *Realtys) Create(realtys Realtys) {
	db := connect()
	db.Create(&realtys)
}

func (r *Realtys) Update(id int, realtys Realtys) {
	db := connect()
	db.Model(&Realtys{}).Where("id = ?", id).Updates(realtys)
}

func (r *Realtys) Delete(id int) {
	db := connect()
	db.Where("id = ?", id).Delete(&Realtys{})
}
