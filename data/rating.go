package data

import "database/sql"

type Rating struct {
	ID          uint
	Description string
	Sentiment   sql.NullBool
	DishID      uint
}
