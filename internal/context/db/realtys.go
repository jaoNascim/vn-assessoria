package db

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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

var connStr = "postgres://postgres:admin@localhost:5432/vn?sslmode=disable"

func connect() *gorm.DB {
	db, err := sql.Open("postgres", connStr)
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
