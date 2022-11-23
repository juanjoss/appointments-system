package ports

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const BodyParsingError = "No se pudo obtener los datos del formulario."
const ParamParsingError = "No se pudo obtener el parámetro."
const PhoneNumberError = "El número de teléfono ingresado es invalido."
const InvalidEmailError = "El email ingresado no es válido."
const InvalidPasswordError = "La contraseña ingresada no es válida."
const SignedStringError = "No se pudo firmar el token."
const EmailNotificationError = "Hubo un error al enviar el email."

// Administrator errors
const (
	GetAdminError           = "Error de autenticación."
	UpdateAdminProfileError = "No se pudo actualizar los datos."
)

// Client errors
const (
	GetClientError    = "No se pudo obtener al cliente (%d)."
	GetClientsError   = "No se pudo obtener a los clientes."
	CreateClientError = "No se pudo crear al cliente."
	UpdateClientError = "No se pudo actualizar al cliente."
)

// Booking errors
const (
	GetBookingsError   = "No se pudo obtener las reservas."
	CreateBookingError = "No se pudo realizar la reserva."
	DeleteBookingError = "No se pudo eliminar la reserva."
)

// AttentionDay errors
const (
	GetAttentionDaysError         = "No se pudo obtener los dias de atención."
	GetInactiveAttentionDaysError = "No se pudo obtener los días de atención inactivos."
	UpdateAttentionDayError       = "No se pudo actualizar el día de atención %s."
)

// AttentionHour errors
const (
	GetAttentionHoursError  = "No se pudo obtener los horarios de atención."
	GetAttentionHourIdError = "No se pudo obtener el horario de atención."
)

// Appointment errors
const (
	GetAppointmentError          = "No se pudo obtener el turno (%d)."
	UpsertAppointmentError       = "No se pudo crear o actualizar el turno."
	AppointmentNotAvailableError = "El turno elegido no se encuentra disponible."
)

// AppointmentType errors
const (
	GetAllAppointmentTypesError   = "No se pudo obtener los tipos de consulta."
	GetAppointmentTypeError       = "No se pudo obtener el tipo de consulta (%d)."
	GetAppointmentTypesCountError = "No se pudo obtener la cantidad de tipos de consulta."
	NoAppointmentTypesFoundError  = "Debe crear al menos un tipo de consulta antes de habilitar días de atención."
	CreateAppointmentTypeError    = "No se pudo crear el tipo de consulta."
	UpdateAppointmentTypeError    = "No se pudo actualizar el tipo de consulta."
	DeleteAppointmentTypeError    = "No se pudo eliminar el tipo de consulta. Verifique que no haya reservas activas para el tipo de consulta"
)

// PaymentMethod errors
const (
	GetPaymentMethodsError = "No se pudo obtener los tipos de pago."
	GetPaymentMethodError  = "No se pudo obtener el tipo de pago (%d)."
)

func ErrorParam(msg string) fiber.Map {
	return fiber.Map{
		"queries": map[string]string{
			"error": msg,
		},
	}
}

func FmtErrorParam(msg string, params any) fiber.Map {
	return fiber.Map{
		"queries": map[string]string{
			"error": fmt.Sprintf(
				msg,
				params,
			),
		},
	}
}
