package models

import (
	"database/sql"
	"time"
)

type UsersEllo struct {
	ID             uint
	UserID         uint
	Password       string
	Email          string
	Gender         string
	Avatar         sql.NullString
	Link           sql.NullString
	DateOfBirth    *time.Time
	EmailConfirmed bool `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (UsersEllo) TableName() string {
	return "users_ello"
}
