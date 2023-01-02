package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ChatGetMyChatList
// chat.getMyChatList user_id:long is_creator:Bool = Vector<MutableChat>;
func (c *ChatCore) ChatGetMyChatList(in *chat.TLChatGetMyChatList) (*chat.Vector_MutableChat, error) {
	var (
		chatList = make([]*chat.MutableChat, 0)
	)

	//
	if mtproto.FromBool(in.IsCreator) {
		c.svcCtx.Dao.ChatParticipantsDAO.SelectMyAdminListWithCB(
			c.ctx,
			in.UserId,
			func(i int, v int64) {
				chat, err := c.svcCtx.Dao.GetMutableChat(c.ctx, v, in.UserId)
				if err != nil {
					c.Logger.Errorf("chat.getMyChatList - error: %v", err)
				} else if chat != nil {
					chatList = append(chatList, chat)
				}
			})
	} else {
		// TODO(@benqi):
	}

	return &chat.Vector_MutableChat{
		Datas: chatList,
	}, nil
}
