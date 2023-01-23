package dataobject

import "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"

type ChannelDO struct {
	Id               int64  `db:"id"`
	CreatorUserId    int64  `db:"creator_user_id"`
	AccessHash       int64  `db:"access_hash"`
	RandomId         int64  `db:"random_id"`
	ParticipantCount int32  `db:"participant_count"`
	Title            string `db:"title"`
	About            string `db:"about"`
	PhotoId          int64  `db:"photo_id"`
	Link             string `db:"link"`
	AdminsEnabled    int8   `db:"admins_enabled"`
	Deactivated      int8   `db:"deactivated"`
	Version          int32  `db:"version"`
	Date             int32  `db:"date"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}

func (ch *ChannelDO) ToChannel() *channels.Channel {
	return &channels.Channel{
		Id:               ch.Id,
		CreatorUserId:    ch.CreatorUserId,
		AccessHash:       ch.AccessHash,
		RandomId:         ch.RandomId,
		ParticipantCount: ch.ParticipantCount,
		Title:            ch.Title,
		About:            ch.About,
		PhotoId:          ch.PhotoId,
		Link:             ch.Link,
		AdminsEnabled:    int32(ch.AdminsEnabled),
		Deactivated:      int32(ch.Deactivated),
		Version:          ch.Version,
		Date:             ch.Date,
		CreatedAt:        ch.CreatedAt,
		UpdatedAt:        ch.UpdatedAt,
	}
}
