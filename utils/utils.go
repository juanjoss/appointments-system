package utils

import (
	"appointments-system/models"
	"appointments-system/ports"
	"strconv"
	"time"
)

func ToDDMMYYYY(m time.Time) string {
	return strconv.Itoa(m.Day()) + "/" +
		strconv.Itoa(int(m.Month())) + "/" +
		strconv.Itoa(m.Year())
}

func ToYYYYMMDD(m time.Time) string {
	year := strconv.Itoa(m.Year())
	month := strconv.Itoa(int(m.Month()))
	day := strconv.Itoa(m.Day())

	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}

	return year + "-" + month + "-" + day
}

func BuildBookingViewId(bv ports.BookingView) string {
	return strconv.Itoa(bv.Client.Id) +
		strconv.Itoa(bv.Appointment.Id) +
		strconv.Itoa(bv.AppointmentType.Id) +
		strconv.Itoa(bv.PaymentMethod.Id)
}

func BuildBookingId(b models.Booking) string {
	return strconv.Itoa(b.ClientId) +
		strconv.Itoa(b.AppointmentId) +
		strconv.Itoa(b.AppointmentTypeId) +
		strconv.Itoa(b.PaymentMethodId)
}
