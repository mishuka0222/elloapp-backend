package dataobject

import (
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type ChannelDO struct {
	Id               int64  `db:"id"`
	CreatorUserId    int64  `db:"creator_user_id"`
	AccessHash       int64  `db:"access_hash"`
	RandomId         int64  `db:"random_id"`
	ParticipantCount int32  `db:"participant_count"`
	Title            string `db:"title"`
	About            string `db:"about"`
	Photo            string `db:"photo"`
	Link             string `db:"link"`
	Username         string `db:"username"`
	AdminsEnabled    bool   `db:"admins_enabled"`
	Deactivated      bool   `db:"deactivated"`
	Version          int32  `db:"version"`
	Date             int32  `db:"date"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}

func (ch *ChannelDO) ToChannel() *channels.Channel {
	var photo mtproto.Photo
	_ = jsonx.UnmarshalFromString(ch.Photo, &photo)

	return &channels.Channel{
		Id:               ch.Id,
		CreatorUserId:    ch.CreatorUserId,
		AccessHash:       ch.AccessHash,
		RandomId:         ch.RandomId,
		ParticipantCount: ch.ParticipantCount,
		Title:            ch.Title,
		About:            ch.About,
		Photo:            &photo,
		Link:             ch.Link,
		Username:         ch.Username,
		AdminsEnabled:    ch.AdminsEnabled,
		Deactivated:      ch.Deactivated,
		Version:          ch.Version,
		Date:             ch.Date,
		CreatedAt:        ch.CreatedAt,
		UpdatedAt:        ch.UpdatedAt,
	}
}
