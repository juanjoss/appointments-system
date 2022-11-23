package db

import "appointments-system/models"

func (pr *postgresRepo) GetAttentionHours() ([]models.AttentionHour, error) {
	response := []models.AttentionHour{}

	err := pr.DB.Select(
		&response,
		`SELECT start_hour, end_hour 
		FROM attention_hours`,
	)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (pr *postgresRepo) GetAttentionHourId(start, end int) (int, error) {
	var id int

	err := pr.DB.Get(
		&id,
		`SELECT id 
		FROM attention_hours
		WHERE 
			EXTRACT(HOUR FROM start_hour) = $1 AND 
			EXTRACT(HOUR FROM end_hour) = $2`,
		start, end,
	)
	if err != nil {
		return id, err
	}

	return id, nil
}
