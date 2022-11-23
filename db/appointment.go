package db

import (
	"appointments-system/models"
	"time"
)

func (pr *postgresRepo) GetAppointment(id int) (models.Appointment, error) {
	var appointment models.Appointment

	err := pr.DB.Get(
		&appointment,
		`SELECT ap.id, ap.date, ad.day, ad.active, ah.start_hour, ah.end_hour, ap.status
		FROM 
			(attention_days AS ad JOIN appointments ap ON ad.day =  ap.attention_day)
			JOIN attention_hours AS ah ON ap.attention_hour = ah.id
		WHERE ap.id = $1`,
		id,
	)
	if err != nil {
		return appointment, err
	}

	return appointment, nil
}

func (pr *postgresRepo) GetAppointmentFor(date time.Time, hour models.AttentionHour) (models.Appointment, error) {
	var appointment models.Appointment

	err := pr.DB.Get(
		&appointment,
		`SELECT ap.id, ap.date, ad.day, ad.active, ah.start_hour, ah.end_hour, ap.status
		FROM 
			(attention_days AS ad JOIN appointments ap ON ad.day =  ap.attention_day)
			JOIN attention_hours AS ah ON ap.attention_hour = ah.id
		WHERE ap.date = $1 AND ah.start_hour = $2 AND ah.end_hour = $3`,
		date, hour.StartHour, hour.EndHour,
	)
	if err != nil {
		return appointment, err
	}

	return appointment, nil
}

func (pr *postgresRepo) GetAppointmentId(date time.Time, attentionDayId string, attentionHourId int) (int, error) {
	var id int

	err := pr.DB.Get(
		&id,
		`SELECT id
		FROM appointments
		WHERE 
			date = $1 AND 
			attention_day = $2 AND 
			attention_hour = $3 AND
			status = 'available'`,
		date, attentionDayId, attentionHourId,
	)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (pr *postgresRepo) UpsertAppointment(date time.Time, attentionDayId string, attentionHourId int, status string) error {
	_, err := pr.DB.Exec(
		`INSERT INTO appointments (date, attention_day, attention_hour, status)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (date, attention_day, attention_hour) 
			DO UPDATE SET 
				id = (
					SELECT id
					FROM appointments
					WHERE 
						date = $1 AND 
						attention_day = $2 AND 
						attention_hour = $3 AND
						status != 'reserved'
				), 
				status = $4`,
		date, attentionDayId, attentionHourId, status,
	)
	if err != nil {
		return err
	}

	return nil
}
