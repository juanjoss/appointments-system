package models

import "time"

/*
	Appointment defines the day, month, year, schedule and
	status.

	Possible values for Status:

	1. disabled: attention day isn't active.

	2. available: attention day is active.

	3. reserved: appointment already booked.
*/
type Appointment struct {
	Id   int       `json:"id" db:"id"`
	Date time.Time `json:"date" db:"date"`
	AttentionDay
	AttentionHour
	Status string `json:"status" db:"status"`
}

/*
	Attention day defines a weekday with an active field
	to make it available or not for booking appointments.
*/
type AttentionDay struct {
	Day    string `json:"day" db:"day"`
	Active bool   `json:"active" db:"active"`
}

/*
	Attention time defines a time lapse.
*/
type AttentionHour struct {
	StartHour time.Time `json:"start_hour" db:"start_hour"`
	EndHour   time.Time `json:"end_hour" db:"end_hour"`
}
