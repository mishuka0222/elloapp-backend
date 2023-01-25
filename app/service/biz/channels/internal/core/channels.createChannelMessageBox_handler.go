package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

// deprecated
// channel_message_box
func (c *ChannelsCore) CreateChannelMessageBox(in *channels.CreateChannelMessageBoxReq) (res *channels.ChannelMessageBox, err error) {
	var (
		now                 = int32(time.Now().Unix())
		boxId               int32
		messageDataDO       *dataobject.MessageDataDO
		channelMessageBoxDO *dataobject.ChannelMessageBoxDO
		messageData         []byte
	)
	//boxId = c.svcCtx.IDGenClient2.NextChannelMessageBoxId(c.ctx, in.ChannelId)
	boxId = c.svcCtx.IDGenClient2.NextMessageBoxId(c.ctx, in.ChannelId)
	messageDataDO = &dataobject.MessageDataDO{
		DialogId:     -in.ChannelId,
		MessageId:    c.svcCtx.IDGenClient2.NextId(c.ctx),
		SenderUserId: in.FromId,
		PeerType:     int8(mtproto.PEER_CHANNEL),
		PeerId:       in.ChannelId,
		RandomId:     in.ClientRandomId,
		Date:         now,
		Deleted:      0,
	}

	channelMessageBoxDO = &dataobject.ChannelMessageBoxDO{
		SenderUserId:        in.FromId,
		ChannelId:           in.ChannelId,
		ChannelMessageBoxId: boxId,
		MessageId:           messageDataDO.MessageId,
		Date:                now,
	}

	switch in.Message.GetConstructor() {
	case mtproto.CRC32_messageEmpty:
		messageDataDO.MessageType = MESSAGE_TYPE_MESSAGE_EMPTY
	case mtproto.CRC32_message:
		messageDataDO.MessageType = MESSAGE_TYPE_MESSAGE
		message := in.Message.To_Message()

		// mentioned = message.GetMentioned()
		message.SetId(channelMessageBoxDO.ChannelMessageBoxId)
	case mtproto.CRC32_messageService:
		messageDataDO.MessageType = MESSAGE_TYPE_MESSAGE_SERVICE
		message := in.Message.To_MessageService()

		// mentioned = message.GetMentioned()
		message.SetId(channelMessageBoxDO.ChannelMessageBoxId)
	}

	messageData, err = json.Marshal(in.Message)
	messageDataDO.MessageData = string(messageData)

	//// TODO(@benqi): pocess clientRandomId dup
	//_, err = c.svcCtx.Dao.MessageDatasDAO.Insert(c.ctx, messageDataDO)
	//if err != nil {
	//	return
	//}
	_, err = c.svcCtx.Dao.ChannelMessageBoxesDAO.Insert(c.ctx, channelMessageBoxDO)
	if err != nil {
		return
	}

	res = &channels.ChannelMessageBox{
		SenderUserId:        in.FromId,
		ChannelId:           in.ChannelId,
		ChannelMessageBoxId: boxId,
		MessageId:           channelMessageBoxDO.MessageId,
		RandomId:            in.ClientRandomId,
		Message:             in.Message,
	}

	return
}
