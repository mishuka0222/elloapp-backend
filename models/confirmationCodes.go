package models

import "time"

type ConfirmationCodes struct {
	ID        uint
	UserID    uint
	Code      string
	ExpiredAt *time.Time
}
