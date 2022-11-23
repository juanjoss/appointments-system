package ports

import (
	"appointments-system/models"
	"time"
)

/*
POST /bookings request model
*/
type BookAppointmentRequest struct {
	AppointmentDate    time.Time `form:"appointment_date"`
	AttentionHourStart int       `form:"attention_hour_start"`
	AttentionHourEnd   int       `form:"attention_hour_end"`
	AttentionDay       int       `form:"attention_day"`
	AppointmentTypeId  int       `form:"appointment_type"`
	PaymentMethodId    int       `form:"payment_method"`
	models.Client
}

/*
GET /bookings view model
*/
type BookingView struct {
	Client          models.Client
	AppointmentType models.AppointmentType
	PaymentMethod   models.PaymentMethod
	Appointment     models.Appointment
}
