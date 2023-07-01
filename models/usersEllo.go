package models

import (
	"database/sql"
	"time"
)

type UsersEllo struct {
	ID             uint
	UserID         uint
	Username       string
	Password       string
	Email          string
	Gender         string
	Type           string `gorm:"default:personal"`
	Kind           string `gorm:"default:public"`
	Avatar         sql.NullString
	Link           sql.NullString
	DateOfBirth    sql.NullTime
	EmailConfirmed bool `gorm:"default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (UsersEllo) TableName() string {
	return "users_ello"
}
