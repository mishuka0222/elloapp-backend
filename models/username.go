package models

import "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/time"

type Username struct {
	ID        uint
	Username  string
	PeerType  int64 `gorm:"default:2"`
	PeerId    uint
	Deleted   bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Username) TableName() string {
	return "username"
}
