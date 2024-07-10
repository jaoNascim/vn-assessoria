package db

import (
	"time"

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

func (r *Realtys) CreateRealty(realtys Realtys) {
	db := connect()
	db.Create(&realtys)
}

func (r *Realtys) UpdateRealty(id int, realtys Realtys) {
	db := connect()
	db.Model(&Realtys{}).Where("id = ?", id).Updates(realtys)
}

func (r *Realtys) DeleteRealty(id int) {
	db := connect()
	db.Where("id = ?", id).Delete(&Realtys{})
}
