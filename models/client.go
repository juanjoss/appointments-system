package models

import "time"

type Client struct {
	Id          int       `db:"id"`
	FullName    string    `form:"full_name" db:"full_name"`
	Birthdate   time.Time `form:"birthdate" db:"birthdate"`
	Address     string    `form:"address" db:"address"`
	Locality    string    `form:"locality" db:"locality"`
	Country     string    `form:"country" db:"country"`
	CountryCode string    `form:"country_code" db:"country_code"`
	Email       string    `form:"email" db:"email"`
	PhoneNumber string    `form:"phone_number" db:"phone_number"`
}
