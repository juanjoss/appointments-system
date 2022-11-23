package db

import "appointments-system/models"

func (pr *postgresRepo) GetAllAppointmentTypes() ([]models.AppointmentType, error) {
	var appointmentTypes []models.AppointmentType

	err := pr.DB.Select(
		&appointmentTypes,
		`SELECT id, name, description, price 
		FROM appointment_types 
		ORDER BY id ASC`,
	)
	if err != nil {
		return appointmentTypes, err
	}

	return appointmentTypes, nil
}

func (pr *postgresRepo) GetAppointmentTypesCount() (int, error) {
	var count int

	err := pr.DB.Get(
		&count,
		`SELECT COUNT(id)
		FROM appointment_types`,
	)
	if err != nil {
		return count, err
	}

	return count, nil
}

func (pr *postgresRepo) GetAppointmentType(id int) (models.AppointmentType, error) {
	var appointmentType models.AppointmentType

	err := pr.DB.Get(
		&appointmentType,
		`SELECT id, name, description, price 
		FROM appointment_types 
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return appointmentType, err
	}

	return appointmentType, nil
}

func (pr *postgresRepo) CreateAppointmentType(request models.AppointmentType) error {
	_, err := pr.DB.NamedExec(
		`INSERT INTO appointment_types (name, description, price)
		VALUES (:name, :description, :price)
		ON CONFLICT (id) DO UPDATE SET (name, description, price) = (:name, :description, :price)`,
		request,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pr *postgresRepo) UpdateAppointmentType(request models.AppointmentType) error {
	_, err := pr.DB.NamedExec(
		`UPDATE appointment_types 
		SET name = :name, description = :description, price = :price
		WHERE id = :id`,
		request,
	)
	if err != nil {
		return err
	}

	return nil
}

func (pr *postgresRepo) DeleteAppointmentType(id int) error {
	_, err := pr.DB.Exec(`DELETE FROM appointment_types WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
