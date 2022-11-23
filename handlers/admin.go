package handlers

import (
	"appointments-system/db"
	"appointments-system/models"
	"appointments-system/ports"
	"appointments-system/utils"
	"fmt"
	"text/template"

	"github.com/dongri/phonenumber"
	"github.com/gofiber/fiber/v2"
)

/*
Handler for the administrator's appointments page.
*/
func AppointmentsPage(c *fiber.Ctx) error {
	return c.Render("admin/appointments", nil)
}

/*
Handler for the administrator's profile page.
*/
func AdminProfilePage(c *fiber.Ctx) error {
	// validate JWT token
	_, _, err := GetJWT(c)
	if err != nil {
		return err
	}

	// get the administrator
	admin, err := db.Get().GetAdmin()
	if err != nil {
		return c.RedirectToRoute(
			"login",
			ports.ErrorParam(ports.GetAdminError),
		)
	}

	// get attention days
	days, err := db.Get().GetAttentionDays()
	if err != nil {
		return c.RedirectToRoute(
			"admin_profile",
			ports.ErrorParam(ports.GetAttentionDaysError),
		)
	}

	return c.Render("admin/profile", fiber.Map{
		"Admin":         admin,
		"AttentionDays": days,
		"error":         c.Query("error"),
	})
}

/*
Handler to update administrator's information.
*/
func UpdateAdminProfile(c *fiber.Ctx) error {
	// check if at leats one appointment type exists
	appointmentTypesCount, err := db.Get().GetAppointmentTypesCount()
	if err != nil {
		return c.RedirectToRoute(
			"admin_profile",
			ports.ErrorParam(ports.GetAppointmentTypesCountError),
		)
	}

	if appointmentTypesCount == 0 {
		return c.RedirectToRoute(
			"admin_profile",
			ports.ErrorParam(ports.NoAppointmentTypesFoundError),
		)
	}

	var request models.Administrator

	// get administrator's data
	err = c.BodyParser(&request)
	if err != nil {
		return c.RedirectToRoute(
			"admin_profile",
			ports.ErrorParam(ports.BodyParsingError),
		)
	}

	// validate phone number
	phoneNumber := phonenumber.ParseWithLandLine(request.PhoneNumber, "AR")

	if phoneNumber == "" {
		return c.RedirectToRoute(
			"admin_profile",
			ports.ErrorParam(ports.PhoneNumberError),
		)
	}

	// update admin profile
	err = db.Get().UpdateAdminProfile(request)
	if err != nil {
		return c.RedirectToRoute(
			"admin_profile",
			ports.ErrorParam(ports.UpdateAdminProfileError),
		)
	}

	// get attention days
	days, err := db.Get().GetAttentionDays()
	if err != nil {
		return c.RedirectToRoute(
			"admin_profile",
			ports.ErrorParam(ports.GetAttentionDaysError),
		)
	}

	// get the checkbox values from the form
	for _, day := range days {
		if value := c.FormValue(day.Day, "off"); value == "on" {
			day.Active = true
		} else {
			day.Active = false
		}

		// update attention day
		err := db.Get().UpdateAttentionDay(day)
		if err != nil {
			return c.RedirectToRoute(
				"admin_profile",
				ports.FmtErrorParam(ports.UpdateAttentionDayError, day.Day),
			)
		}
	}

	return c.RedirectToRoute("admin_profile", nil)
}

/*
Handler for the administrator's bookings page.
*/
func AdminBookingsPage(c *fiber.Ctx) error {
	// get bookings
	bookings, err := db.Get().GetBookings()
	if err != nil {
		return c.RedirectToRoute(
			"admin_bookings",
			ports.ErrorParam(ports.GetBookingsError),
		)
	}

	var bookingsView []ports.BookingView

	// create the bookingView
	for _, booking := range bookings {
		// find client
		client, err := db.Get().GetClient(booking.ClientId)
		if err != nil {
			return c.RedirectToRoute(
				"admin_bookings",
				ports.FmtErrorParam(ports.GetClientError, booking.ClientId),
			)
		}

		// find appointment type
		appointmentType, err := db.Get().GetAppointmentType(booking.AppointmentTypeId)
		if err != nil {
			return c.RedirectToRoute(
				"admin_bookings",
				ports.FmtErrorParam(ports.GetAppointmentTypeError, booking.AppointmentTypeId),
			)
		}

		// find payment method
		paymentMethod, err := db.Get().GetPaymentMethod(booking.PaymentMethodId)
		if err != nil {
			return c.RedirectToRoute(
				"admin_bookings",
				ports.FmtErrorParam(ports.GetPaymentMethodError, booking.PaymentMethodId),
			)
		}

		// find appointment
		appointment, err := db.Get().GetAppointment(booking.AppointmentId)
		if err != nil {
			return c.RedirectToRoute(
				"admin_bookings",
				ports.FmtErrorParam(ports.GetAppointmentError, booking.AppointmentId),
			)
		}

		bookingsView = append(bookingsView, ports.BookingView{
			Client:          client,
			AppointmentType: appointmentType,
			PaymentMethod:   paymentMethod,
			Appointment:     appointment,
		})
	}

	// functions used in the template
	funcMap := template.FuncMap{
		"toDDMMYYYY": utils.ToDDMMYYYY,
		"bookingId":  utils.BuildBookingViewId,
	}

	return c.Render("admin/bookings", fiber.Map{
		"Bookings": bookingsView,
		"FuncMap":  funcMap,
		"error":    c.Query("error"),
	})
}

/*
Handler to delete a booking.
*/
func DeleteBooking(c *fiber.Ctx) error {
	var request models.Booking

	// get booking data
	err := c.BodyParser(&request)
	if err != nil {
		return c.RedirectToRoute(
			"admin_bookings",
			ports.ErrorParam(ports.BodyParsingError),
		)
	}

	// delete booking
	err = db.Get().DeleteBooking(request)
	if err != nil {
		return c.RedirectToRoute(
			"admin_bookings",
			ports.ErrorParam(ports.DeleteBookingError),
		)
	}

	// send email
	admin, err := db.Get().GetAdmin()
	if err != nil {
		return c.RedirectToRoute(
			"admin_booking",
			ports.ErrorParam(ports.GetAdminError),
		)
	}

	client, err := db.Get().GetClient(request.ClientId)
	if err != nil {
		return c.RedirectToRoute(
			"admin_booking",
			ports.ErrorParam(ports.GetClientError),
		)
	}

	appointment, err := db.Get().GetAppointment(request.AppointmentId)
	if err != nil {
		return c.RedirectToRoute(
			"admin_booking",
			ports.ErrorParam(ports.GetAppointmentError),
		)
	}

	err = utils.SendEmail(
		admin.Email,
		client.Email,
		fmt.Sprintf("Reserva %s", utils.BuildBookingId(request)),
		fmt.Sprintf(
			utils.BookingCancelledTemplate(),
			appointment.Day,
			utils.ToDDMMYYYY(appointment.Date),
			appointment.StartHour.Hour(),
			appointment.EndHour.Hour(),
			admin.PhoneNumber,
			admin.Email,
		),
	)
	if err != nil {
		return c.RedirectToRoute(
			"admin_bookings",
			ports.ErrorParam(ports.EmailNotificationError),
		)
	}

	return c.RedirectToRoute("admin_bookings", nil)
}

/*
Handler to the admin's clients page.
*/
func AdminClientsPage(c *fiber.Ctx) error {
	// get clients
	clients, err := db.Get().GetClients()
	if err != nil {
		return c.RedirectToRoute(
			"admin_clients",
			ports.ErrorParam(ports.GetClientsError),
		)
	}

	// functions used in the template
	funcMap := template.FuncMap{
		"toDDMMYYYY": utils.ToDDMMYYYY,
		"toYYYYMMDD": utils.ToYYYYMMDD,
	}

	return c.Render("admin/clients", fiber.Map{
		"Clients": clients,
		"FuncMap": funcMap,
		"error":   c.Query("error"),
	})
}

/*
Handler to update clients.
*/
func UpdateClient(c *fiber.Ctx) error {
	var request models.Client

	// get client data
	err := c.BodyParser(&request)
	if err != nil {
		return c.RedirectToRoute(
			"admin_clients",
			ports.ErrorParam(ports.BodyParsingError),
		)
	}

	// validate phone number
	phoneNumber := phonenumber.ParseWithLandLine(request.PhoneNumber, request.CountryCode)

	if phoneNumber == "" {
		return c.RedirectToRoute(
			"admin_clients",
			ports.ErrorParam(ports.PhoneNumberError),
		)
	}

	// update client
	err = db.Get().UpdateClient(request)
	if err != nil {
		return c.RedirectToRoute(
			"admin_clients",
			ports.ErrorParam(ports.UpdateClientError),
		)
	}

	return c.RedirectToRoute("admin_clients", nil)
}
