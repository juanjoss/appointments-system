package ports

import (
	"appointments-system/models"
	"time"
)

type Repository interface {
	/*
		Get the default admin.
	*/
	GetAdmin() (models.Administrator, error)

	/*
		Update an admin's data.
	*/
	UpdateAdminProfile(models.Administrator) error

	/*
		Get all appointment types.
	*/
	GetAllAppointmentTypes() ([]models.AppointmentType, error)

	/*
		Get the count of appointment types.
	*/
	GetAppointmentTypesCount() (int, error)

	/*
		Get an appointment types by its ID.
	*/
	GetAppointmentType(int) (models.AppointmentType, error)

	/*
		Create an appointment type.
	*/
	CreateAppointmentType(models.AppointmentType) error

	/*
		Update an appointment type.
	*/
	UpdateAppointmentType(models.AppointmentType) error

	/*
		Delete an appointment type.
	*/
	DeleteAppointmentType(int) error

	/*
		Get an appointment by its ID.
	*/
	GetAppointment(int) (models.Appointment, error)

	/*
		Get appointment for a given date.
	*/
	GetAppointmentFor(time.Time, models.AttentionHour) (models.Appointment, error)

	/*
		Get an appointment's ID.
	*/
	GetAppointmentId(time.Time, string, int) (int, error)

	/*
		Create or update an appointment if exists.
	*/
	UpsertAppointment(time.Time, string, int, string) error

	/*
		Get all attention days.
	*/
	GetAttentionDays() ([]models.AttentionDay, error)

	/*
		Get all inactive attention days.
	*/
	GetInactiveAttentionDays() ([]models.AttentionDay, error)

	/*
		Update the active state of an attention day.
	*/
	UpdateAttentionDay(models.AttentionDay) error

	/*
		Get all attention hours.
	*/
	GetAttentionHours() ([]models.AttentionHour, error)

	/*
		Get an attention hour id.
	*/
	GetAttentionHourId(int, int) (int, error)

	/*
		Get all payment types.
	*/
	GetPaymentMethods() ([]models.PaymentMethod, error)

	/*
		Get a payment types by its ID.
	*/
	GetPaymentMethod(int) (models.PaymentMethod, error)

	/*
		Get all clients.
	*/
	GetClients() ([]models.Client, error)

	/*
		Get a client by its ID.
	*/
	GetClient(int) (models.Client, error)

	/*
		Create a client.
	*/
	CreateClient(models.Client) (int, error)

	/*
		Update a client.
	*/
	UpdateClient(models.Client) error

	/*
		Get all bookings.
	*/
	GetBookings() ([]models.Booking, error)

	/*
		Create a booking.
	*/
	CreateBooking(models.Booking) error

	/*
		Delete a booking.
	*/
	DeleteBooking(models.Booking) error
}
