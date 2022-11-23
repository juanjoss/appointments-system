package models

type PaymentMethod struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
