package db

import _ "github.com/lib/pq"

type ContactLogs struct {
	Name    string
	Email   string
	Phone   string
	Subject string
	Message string
}

func (r *ContactLogs) CreateContactLog(contactLogs ContactLogs) {
	db := connect()
	db.Create(&contactLogs)
}

func (r *ContactLogs) UpdateContactLog(id int, contactLogs ContactLogs) {
	db := connect()
	db.Model(&ContactLogs{}).Where("id = ?", id).Updates(contactLogs)
}

func (r *ContactLogs) DeleteContactLog(id int) {
	db := connect()
	db.Where("id = ?", id).Delete(&ContactLogs{})
}
