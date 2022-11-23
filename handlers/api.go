package handlers

import (
	"appointments-system/db"
	"appointments-system/ports"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

/*
Returns the settings for the UI calendar.
*/
func GetCalendar(c *fiber.Ctx) error {
	var settings ports.CalendarSettings

	// excluded week days
	var excludedDays string

	inactiveDays, err := db.Get().GetInactiveAttentionDays()
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(ports.ErrorParam(ports.GetInactiveAttentionDaysError))
	}

	excludedDays += strconv.Itoa(ports.WeekDays["Sábado"])
	excludedDays += strconv.Itoa(ports.WeekDays["Domingo"])

	for _, inactiveDay := range inactiveDays {
		excludedDays += strconv.Itoa(ports.WeekDays[inactiveDay.Day])
	}

	// end date (now + administrator's availability period)
	admin, err := db.Get().GetAdmin()
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(ports.ErrorParam(ports.GetAdminError))
	}
	endDate := time.Now().AddDate(0, admin.AvailabilityPeriod, 0)

	// calendar settings
	settings = ports.CalendarSettings{
		DateFormat:       ports.CalendarDateFormat,
		Language:         ports.CalendarDefaultLanguage,
		ExcludedWeekDays: excludedDays,
		EndDate:          fmt.Sprintf("%d-%d-%d", endDate.Year(), endDate.Month(), endDate.Day()),
		MaxViewMode:      ports.CalendarMaxViewMode,
	}

	return c.Status(fiber.StatusOK).JSON(settings)
}

/*
Returns the attention hours from an appointment's requested date.
*/
func GetAttentionHours(c *fiber.Ctx) error {
	var request ports.AttentionHoursRequest

	// get date
	err := c.BodyParser(&request)
	if err != nil {
		return c.JSON(ports.ErrorParam(ports.BodyParsingError))
	}

	// find attention hours for requested date
	attentionHours, err := db.Get().GetAttentionHours()
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(ports.ErrorParam(ports.GetAttentionHoursError))
	}

	var response []ports.AttentionHourResponse

	now := time.Now()
	year, month, day := now.Date()

	for _, hour := range attentionHours {
		// check if the appointment time has already passed
		if request.Year == year && time.Month(request.Month) == month && request.Day == day {
			if now.Hour() >= hour.StartHour.Hour() {
				continue
			}
		}

		// get the appointment
		appointment, err := db.Get().GetAppointmentFor(
			time.Date(request.Year, time.Month(request.Month), request.Day, 0, 0, 0, 0, time.UTC),
			hour,
		)
		if err != nil {
			// appointment not in the DB, status = 'disabled'
			response = append(response, ports.AttentionHourResponse{
				StartHour: hour.StartHour.Hour(),
				EndHour:   hour.EndHour.Hour(),
			})
			continue
		}

		response = append(response, ports.AttentionHourResponse{
			StartHour: hour.StartHour.Hour(),
			EndHour:   hour.EndHour.Hour(),
			Status:    appointment.Status,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

/*
Updates or inserts an appointment.
*/
func UpdateAppointment(c *fiber.Ctx) error {
	var request ports.UpsertAppointmentRequest

	// get appointment data
	err := c.BodyParser(&request)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(ports.ErrorParam(ports.BodyParsingError))
	}

	// find attention day
	var day string
	for weekday, dayIndex := range ports.WeekDays {
		if dayIndex == request.WeekDay {
			day = weekday
			break
		}
	}

	// get status
	var status string
	if request.Status {
		status = "available"
	} else {
		status = "disabled"
	}

	// get attention hour id
	attentionHourId, err := db.Get().GetAttentionHourId(request.StartHour, request.EndHour)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(ports.ErrorParam(ports.GetAttentionHourIdError))
	}

	// update appointment
	err = db.Get().UpsertAppointment(request.Date, day, attentionHourId, status)
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(ports.ErrorParam(ports.UpsertAppointmentError))
	}

	return c.SendStatus(fiber.StatusOK)
}
