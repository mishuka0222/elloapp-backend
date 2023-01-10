package models

import (
	"database/sql"
	"time"
)

type Users struct {
	ID                               uint
	UserType                         int
	AccessHash                       sql.NullInt64
	SecretKeyId                      int
	FirstName                        string
	LastName                         string
	Username                         string
	Phone                            sql.NullString
	CountryCode                      sql.NullString
	Verified                         bool
	Support                          bool
	Scam                             bool
	Fake                             bool
	Premium                          bool
	About                            string
	State                            bool
	IsBot                            bool
	AccountDaysTTL                   int `gorm:"default:180"`
	PhotoId                          int
	Restricted                       bool
	RestrictionReason                string
	ArchiveAndMuteNewNoncontactPeers bool
	EmojiStatusDocumentId            int
	EmojiStatusUntil                 int
	Deleted                          bool
	DeleteReason                     string
	CreatedAt                        time.Time
	UpdatedAt                        time.Time
}
