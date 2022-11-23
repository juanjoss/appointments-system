package models

type AppointmentType struct {
	Id          int     `form:"id" db:"id"`
	Name        string  `form:"name" db:"name"`
	Description string  `form:"description" db:"description"`
	Price       float32 `form:"price" db:"price"`
}
