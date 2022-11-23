package models

type Booking struct {
	ClientId          int `form:"client_id" db:"client_id"`
	AppointmentTypeId int `form:"appointment_type_id" db:"appointment_type_id"`
	PaymentMethodId   int `form:"payment_method_id" db:"payment_method_id"`
	AppointmentId     int `form:"appointment_id" db:"appointment_id"`
}
