package handlers

import (
	"appointments-system/db"
	"appointments-system/models"
	"appointments-system/ports"
	"appointments-system/utils"
	"fmt"
	"log"

	"github.com/dongri/phonenumber"
	"github.com/gofiber/fiber/v2"
)

func BookingsPage(c *fiber.Ctx) error {
	appointmentTypes, err := db.Get().GetAllAppointmentTypes()
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.GetAllAppointmentTypesError),
		)
	}

	paymentMethods, err := db.Get().GetPaymentMethods()
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.GetPaymentMethodsError),
		)
	}

	return c.Render("client/booking", fiber.Map{
		"AppointmentTypes": appointmentTypes,
		"PaymentMethods":   paymentMethods,
		"error":            c.Query("error"),
	})
}

func AppointmentBookedPage(c *fiber.Ctx) error {
	if c.Query("email") == "" {
		return BookingsPage(c)
	}

	return c.Render("client/booked", fiber.Map{
		"error": c.Query("error"),
		"Email": c.Query("email"),
	})
}

// func ProcessPayment(c *fiber.Ctx) error {
// 	return c.Render("client/payment", fiber.Map{
// 		"error": c.Query("error"),

// 	})
// }

func BookAppointment(c *fiber.Ctx) error {
	var request ports.BookAppointmentRequest

	// read the form
	err := c.BodyParser(&request)
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.BodyParsingError),
		)
	}

	// validate phone number
	phoneNumber := phonenumber.ParseWithLandLine(request.PhoneNumber, request.CountryCode)

	if phoneNumber == "" {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.PhoneNumberError),
		)
	}

	// find attention hour id
	attentionHourId, err := db.Get().GetAttentionHourId(request.AttentionHourStart, request.AttentionHourEnd)
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.GetAttentionHourIdError),
		)
	}

	// find attention day id
	var attentionDayId string
	for day, dayIndex := range ports.WeekDays {
		if dayIndex == request.AttentionDay {
			attentionDayId = day
			break
		}
	}

	// find the appointment
	appointmentId, err := db.Get().GetAppointmentId(request.AppointmentDate, attentionDayId, attentionHourId)
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.AppointmentNotAvailableError),
		)
	}

	// process payment
	log.Println("payment processed")
	// ProcessPayment(c)

	// create client
	clientId, err := db.Get().CreateClient(request.Client)
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.CreateClientError),
		)
	}

	// create booking
	booking := models.Booking{
		ClientId:          clientId,
		AppointmentTypeId: request.AppointmentTypeId,
		PaymentMethodId:   request.PaymentMethodId,
		AppointmentId:     appointmentId,
	}

	err = db.Get().CreateBooking(booking)
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.CreateBookingError),
		)
	}

	// send email
	admin, err := db.Get().GetAdmin()
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.GetAdminError),
		)
	}

	appointmentType, err := db.Get().GetAppointmentType(booking.AppointmentTypeId)
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.GetAppointmentTypeError),
		)
	}

	paymentMethod, err := db.Get().GetPaymentMethod(booking.PaymentMethodId)
	if err != nil {
		return c.RedirectToRoute(
			"client_booking",
			ports.ErrorParam(ports.GetPaymentMethodError),
		)
	}

	err = utils.SendEmail(
		admin.Email,
		request.Client.Email,
		fmt.Sprintf("Reserva %s", utils.BuildBookingId(booking)),
		fmt.Sprintf(
			utils.BookingSuccessTemplate(),
			attentionDayId,
			utils.ToDDMMYYYY(request.AppointmentDate),
			request.AttentionHourStart,
			request.AttentionHourEnd,
			appointmentType.Name,
			paymentMethod.Name,
			appointmentType.Price,
			admin.PhoneNumber,
			admin.Email,
		),
	)
	if err != nil {
		panic(err)
	}

	return c.RedirectToRoute("appointment_booked", fiber.Map{
		"queries": map[string]string{
			"email": request.Client.Email,
		},
	})
}
