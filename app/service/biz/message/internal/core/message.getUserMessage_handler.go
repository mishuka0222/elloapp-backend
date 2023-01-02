package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessageGetUserMessage
// message.getUserMessage user_id:long id:int = MessageBox;
func (c *MessageCore) MessageGetUserMessage(in *message.TLMessageGetUserMessage) (*mtproto.MessageBox, error) {
	myDO, err := c.svcCtx.Dao.MessagesDAO.SelectByMessageId(c.ctx, in.UserId, in.Id)
	if err != nil {
		c.Logger.Errorf("message.getUserMessage - error: %v", err)
		return nil, err
	} else if myDO == nil {
		c.Logger.Errorf("message.getUserMessage - error: not found message(%s)", in.DebugString())
		return nil, mtproto.ErrMessageIdInvalid
	}
	return c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, myDO), nil
}
