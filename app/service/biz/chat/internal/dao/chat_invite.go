package dao

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg/env2"
)

func (d *Dao) MakeChatInviteExported(ctx context.Context, chatInviteDO *dataobject.ChatInvitesDO) *mtproto.ExportedChatInvite {
	rValue := mtproto.MakeTLChatInviteExported(&mtproto.ExportedChatInvite{
		Revoked:       chatInviteDO.Revoked,
		Permanent:     chatInviteDO.Permanent,
		RequestNeeded: chatInviteDO.RequestNeeded,
		Link:          "https://" + env2.TDotMe + "/+" + chatInviteDO.Link,
		AdminId:       chatInviteDO.AdminId,
		Date:          int32(chatInviteDO.Date2),
		StartDate:     mtproto.MakeFlagsInt32(int32(chatInviteDO.StartDate)),
		ExpireDate:    mtproto.MakeFlagsInt32(int32(chatInviteDO.ExpireDate)),
		UsageLimit:    mtproto.MakeFlagsInt32(chatInviteDO.UsageLimit),
		Usage:         mtproto.MakeFlagsInt32(chatInviteDO.Usage2),
		Requested:     mtproto.MakeFlagsInt32(chatInviteDO.Requested),
		Title:         mtproto.MakeFlagsString(chatInviteDO.Title),
	}).To_ExportedChatInvite()

	// TODO: calc
	sz := d.CommonDAO.CalcSize(ctx, "chat_invite_participants", map[string]interface{}{
		"link": chatInviteDO.Link,
	})
	rValue.Usage = mtproto.MakeFlagsInt32(int32(sz))

	return rValue
}
