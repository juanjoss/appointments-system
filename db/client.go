package db

import "appointments-system/models"

func (pr *postgresRepo) GetClients() ([]models.Client, error) {
	var clients []models.Client

	err := pr.DB.Select(
		&clients,
		`SELECT id, full_name, birthdate, address, locality, country, country_code, email, phone_number
		FROM clients`,
	)
	if err != nil {
		return clients, err
	}

	return clients, nil
}

func (pr *postgresRepo) GetClient(id int) (models.Client, error) {
	var client models.Client

	err := pr.DB.Get(
		&client,
		`SELECT id, full_name, birthdate, address, locality, country, country_code, email, phone_number
		FROM clients
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return client, err
	}

	return client, nil
}

func (pr *postgresRepo) CreateClient(client models.Client) (int, error) {
	var id int

	rows, err := pr.DB.NamedQuery(
		`INSERT INTO clients (full_name, birthdate, address, locality, country, country_code, email, phone_number)
		VALUES (:full_name, :birthdate, :address, :locality, :country, :country_code, :email, :phone_number)
		ON CONFLICT (email)
			DO UPDATE SET 
				(full_name, birthdate, address, locality, country, country_code, phone_number) 
				= (:full_name, :birthdate, :address, :locality, :country, :country_code, :phone_number)
		RETURNING id`,
		client,
	)
	if err != nil {
		return id, err
	}

	rows.Next()
	if err := rows.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (pr *postgresRepo) UpdateClient(client models.Client) error {
	_, err := pr.DB.NamedExec(
		`UPDATE clients
		SET 
			full_name = :full_name, 
			birthdate = :birthdate, 
			address = :address, 
			locality = :locality,
			country = :country, 
			country_code = :country_code,
			email = :email,
			phone_number = :phone_number
		WHERE id = :id`,
		client,
	)
	if err != nil {
		return err
	}

	return nil
}
