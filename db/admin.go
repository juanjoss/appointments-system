package db

import "appointments-system/models"

func (pr *postgresRepo) GetAdmin() (models.Administrator, error) {
	response := models.Administrator{}

	err := pr.DB.Get(
		&response,
		`SELECT * FROM administrators`,
	)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (pr *postgresRepo) UpdateAdminProfile(request models.Administrator) error {
	_, err := pr.DB.NamedExec(
		`UPDATE administrators 
		SET 
			full_name=:full_name, 
			phone_number=:phone_number, 
			email=:email
		WHERE email = :email`,
		request,
	)
	if err != nil {
		return err
	}

	return nil
}
