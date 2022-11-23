package ports

import (
	"time"
)

var WeekDays = map[string]int{
	"Domingo":   0,
	"Lunes":     1,
	"Martes":    2,
	"Miércoles": 3,
	"Jueves":    4,
	"Viernes":   5,
	"Sábado":    6,
}

const CalendarDateFormat = "yyyy-mm-dd"
const CalendarDefaultLanguage = "es"
const CalendarMaxViewMode = 1

/*
	/api/calendar response model
*/
type CalendarSettings struct {
	DateFormat       string `json:"date_format"`
	Language         string `json:"language"`
	ExcludedWeekDays string `json:"excluded_week_days"`
	EndDate          string `json:"end_date"`
	MaxViewMode      int    `json:"max_view_mode"`
}

/*
	/api/attentionHours request model
*/
type AttentionHoursRequest struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

/*
	/api/attentionHours response model
*/
type AttentionHourResponse struct {
	StartHour int    `json:"start_hour"`
	EndHour   int    `json:"end_hour"`
	Status    string `json:"status"`
}

/*
	/api/appointments request model
*/
type UpsertAppointmentRequest struct {
	Date      time.Time `json:"date"`
	WeekDay   int       `json:"weekday"`
	StartHour int       `json:"start_hour"`
	EndHour   int       `json:"end_hour"`
	Status    bool      `json:"status"`
}
