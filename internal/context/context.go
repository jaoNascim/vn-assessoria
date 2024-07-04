package context

import "vn-assessoria/internal/context/db"

type Realty interface {
	Create(realty db.Realtys)
	Update(id int, realty db.Realtys)
	Delete(id int)
}
