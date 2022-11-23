package db

import "appointments-system/models"

func (pr *postgresRepo) GetBookings() ([]models.Booking, error) {
	var bookings []models.Booking

	err := pr.DB.Select(
		&bookings,
		`SELECT client_id, appointment_type_id, payment_method_id, appointment_id
		FROM bookings`,
	)
	if err != nil {
		return bookings, err
	}

	return bookings, nil
}

func (pr *postgresRepo) CreateBooking(booking models.Booking) error {
	_, err := pr.DB.NamedExec(
		`INSERT INTO bookings (
			client_id, 
			appointment_type_id, 
			payment_method_id, 
			appointment_id
		)
		VALUES (
			:client_id,
			:appointment_type_id,
			:payment_method_id,
			:appointment_id
		)`,
		booking,
	)
	if err != nil {
		return err
	}

	_, err = pr.DB.Exec(
		`UPDATE appointments 
		SET status = 'reserved' 
		WHERE id = $1`,
		booking.AppointmentId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pr *postgresRepo) DeleteBooking(booking models.Booking) error {
	_, err := pr.DB.NamedExec(
		`DELETE FROM bookings 
		WHERE 
			client_id = :client_id AND 
			appointment_type_id = :appointment_type_id AND
			payment_method_id = :payment_method_id AND
			appointment_id = :appointment_id`,
		booking,
	)
	if err != nil {
		return err
	}

	_, err = pr.DB.NamedExec(
		`UPDATE appointments
		SET status = 'available'
		WHERE id = :appointment_id`,
		booking,
	)
	if err != nil {
		return err
	}

	return nil
}
