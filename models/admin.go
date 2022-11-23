package models

import "time"

type Administrator struct {
	FullName           string    `form:"full_name" db:"full_name"`
	PhoneNumber        string    `form:"phone_number" db:"phone_number"`
	Email              string    `form:"email" db:"email"`
	Password           string    `form:"password" db:"password"`
	AvailabilityPeriod int       `form:"availability_period" db:"availability_period"`
	AttentionHourStart time.Time `form:"attention_hour_start" db:"attention_hour_start"`
	AttentionHourEnd   time.Time `form:"attention_hour_end" db:"attention_hour_end"`
}
