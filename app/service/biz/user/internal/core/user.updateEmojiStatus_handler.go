package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserUpdateEmojiStatus
// user.updateEmojiStatus user_id:long emoji_status_document_id:long emoji_status_until:int = Bool;
func (c *UserCore) UserUpdateEmojiStatus(in *user.TLUserUpdateEmojiStatus) (*mtproto.Bool, error) {
	rB := c.svcCtx.Dao.UpdateUserEmojiStatus(
		c.ctx,
		in.UserId,
		in.EmojiStatusDocumentId,
		in.EmojiStatusUntil)

	return mtproto.ToBool(rB), nil
}
