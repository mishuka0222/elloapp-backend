package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"time"
)

// ChatExportChatInvite
// chat.exportChatInvite flags:# chat_id:long admin_id:long legacy_revoke_permanent:flags.2?true request_needed:flags.3?true expire_date:flags.0?int usage_limit:flags.1?int title:flags.4?string = ExportedChatInvite;
func (c *ChatCore) ChatExportChatInvite(in *chat.TLChatExportChatInvite) (*mtproto.ExportedChatInvite, error) {
	chatInviteDO := &dataobject.ChatInvitesDO{
		ChatId:        in.ChatId,
		AdminId:       in.AdminId,
		Link:          chat.GenChatInviteHash(),
		Permanent:     false,
		Revoked:       false,
		RequestNeeded: in.RequestNeeded,
		StartDate:     0,
		ExpireDate:    int64(in.GetExpireDate().GetValue()),
		UsageLimit:    in.GetUsageLimit().GetValue(),
		Usage2:        0,
		Requested:     0,
		Title:         in.GetTitle().GetValue(),
		Date2:         time.Now().Unix(),
	}

	_, _, err := c.svcCtx.Dao.ChatInvitesDAO.Insert(c.ctx, chatInviteDO)
	if err != nil {
		c.Logger.Errorf("chat.exportChatInvite - error: %v", err)
		return nil, err
	}

	return c.svcCtx.Dao.MakeChatInviteExported(c.ctx, chatInviteDO), nil
}
