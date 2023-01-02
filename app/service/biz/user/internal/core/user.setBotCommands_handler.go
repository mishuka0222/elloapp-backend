package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserSetBotCommands
// user.setBotCommands user_id:long bot_id:long commands:Vector<BotCommand> = Bool;
func (c *UserCore) UserSetBotCommands(in *user.TLUserSetBotCommands) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("user.setBotCommands - error: method UserSetBotCommands not impl")

	return nil, mtproto.ErrMethodNotImpl
}
