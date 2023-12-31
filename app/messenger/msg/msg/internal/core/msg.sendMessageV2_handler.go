package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MsgSendMessageV2
// msg.sendMessageV2 user_id:long auth_key_id:long peer_type:int peer_id:long message:Vector<OutboxMessage> = UpdateList;
func (c *MsgCore) MsgSendMessageV2(in *msg.TLMsgSendMessageV2) (*mtproto.UpdateList, error) {
	// TODO: check request valid

	if len(in.Message) == 0 {
		err := mtproto.ErrMessageIdsEmpty
		c.Logger.Errorf("msg.sendMessageV2 - error: %v", err)
		return nil, err
	}

	switch in.PeerType {
	case mtproto.PEER_SELF,
		mtproto.PEER_USER,
		mtproto.PEER_CHAT:
		return c.sendOutgoingMessageV2(in)
	case mtproto.PEER_CHANNEL:
		// TODO
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.sendMessageV2 - error: %v", err)
		return nil, err
	}

	return nil, mtproto.ErrMethodNotImpl
}

func (c *MsgCore) sendOutgoingMessageV2(in *msg.TLMsgSendMessageV2) (*mtproto.UpdateList, error) {
	var (
		rUpdates = &mtproto.UpdateList{Updates: nil}
	)

	// c.svcCtx.Dao.SendUserMessage()
	return rUpdates, nil
}
