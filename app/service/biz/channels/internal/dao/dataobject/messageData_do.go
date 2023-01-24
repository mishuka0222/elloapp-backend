package dataobject

import "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"

type MessageDataDO struct {
	Id           int64  `db:"id"`
	DialogId     int64  `db:"dialog_id"`
	MessageId    int64  `db:"message_id"`
	SenderUserId int64  `db:"sender_user_id"`
	PeerType     int8   `db:"peer_type"`
	PeerId       int64  `db:"peer_id"`
	RandomId     int64  `db:"random_id"`
	MessageType  int8   `db:"message_type"`
	MessageData  string `db:"message_data"`
	Date         int32  `db:"date"`
	Deleted      int8   `db:"deleted"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}

func (m *MessageDataDO) ToMessageData() *channels.MessageData {
	return &channels.MessageData{
		Id:           m.Id,
		DialogId:     m.DialogId,
		SenderUserId: m.SenderUserId,
		PeerType:     int32(m.PeerType),
		PeerId:       m.PeerId,
		RandomId:     m.RandomId,
		MessageType:  int32(m.MessageType),
		MessageData:  m.MessageData,
		Date:         m.Date,
		Deleted:      int32(m.Deleted),
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}
