package context

import "vn-assessoria/internal/context/db"

type Realty interface {
	CreateRealty(realty db.Realtys)
	UpdateRealty(id int, realty db.Realtys)
	DeleteRealty(id int)
}

type ContactLog interface {
	CreateContactLog(contactLogs db.ContactLogs)
	UpdateContactLog(id int, contactLogs db.ContactLogs)
	DeleteContactLog(id int)
}
