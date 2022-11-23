package handlers

import (
	"appointments-system/db"
	"appointments-system/models"
	"appointments-system/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

/*
Handler to get all appointment types.
*/
func GetAllAppointmentTypes(c *fiber.Ctx) error {
	// get appointment types
	appointmentTypes, err := db.Get().GetAllAppointmentTypes()
	if err != nil {
		return c.RedirectToRoute(
			"appointment_types",
			ports.ErrorParam(ports.GetAllAppointmentTypesError),
		)
	}

	return c.Render("admin/appointment_type", fiber.Map{
		"Types": appointmentTypes,
		"error": c.Query("error"),
	})
}

/*
Handler to create or update an appointment type.
*/
func CreateUpdateAppointmentType(c *fiber.Ctx) error {
	var request models.AppointmentType

	// get appointment type data
	err := c.BodyParser(&request)
	if err != nil {
		return c.RedirectToRoute(
			"appointment_types",
			ports.ErrorParam(ports.BodyParsingError),
		)
	}

	// check Id to see if record exists
	if request.Id == 0 {
		// create appointment type
		err = db.Get().CreateAppointmentType(request)
		if err != nil {
			return c.RedirectToRoute(
				"appointment_types",
				ports.ErrorParam(ports.CreateAppointmentTypeError),
			)
		}
	} else {
		// update appointment type
		err = db.Get().UpdateAppointmentType(request)
		if err != nil {
			return c.RedirectToRoute(
				"appointment_types",
				ports.ErrorParam(ports.UpdateAppointmentTypeError),
			)
		}
	}

	return c.RedirectToRoute("appointment_types", nil)
}

/*
Handler to delete an appointment type.
*/
func DeleteAppointmentType(c *fiber.Ctx) error {
	// get appointment type id
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.RedirectToRoute(
			"appointment_types",
			ports.ErrorParam(ports.ParamParsingError),
		)
	}

	// delete appointment type
	err = db.Get().DeleteAppointmentType(id)
	if err != nil {
		return c.RedirectToRoute(
			"appointment_types",
			ports.ErrorParam(ports.DeleteAppointmentTypeError),
		)
	}

	return c.RedirectToRoute("appointment_types", nil)
}
