package db

import "appointments-system/models"

func (pr *postgresRepo) GetAttentionDays() ([]models.AttentionDay, error) {
	response := []models.AttentionDay{}

	err := pr.DB.Select(
		&response,
		`SELECT day, active 
		FROM attention_days`,
	)
	if err != nil {
		return response, err
	}

	return response, err
}

func (pr *postgresRepo) GetInactiveAttentionDays() ([]models.AttentionDay, error) {
	response := []models.AttentionDay{}

	err := pr.DB.Select(
		&response,
		`SELECT day
		FROM attention_days
		WHERE active = FALSE`,
	)
	if err != nil {
		return response, err
	}

	return response, err
}

func (pr *postgresRepo) UpdateAttentionDay(day models.AttentionDay) error {
	_, err := pr.DB.NamedExec(
		`UPDATE attention_days 
		SET active = :active 
		WHERE day = :day`,
		day,
	)
	if err != nil {
		return err
	}

	/*
	* Disabling existing appointments for day.
	 */
	if !day.Active {
		_, err = pr.DB.Exec(
			`UPDATE appointments 
			SET status = 'disabled' 
			WHERE attention_day = $1`,
			day.Day,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
