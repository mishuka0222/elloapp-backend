// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"github.com/teamgram/proto/mtproto"
	"github.com/zeromicro/go-zero/core/jsonx"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/dal/dataobject"
)

func (d *Dao) MakeImmutableChatByDO(chatsDO *dataobject.ChatsDO) (chat *chatpb.ImmutableChat) {
	chat = &chatpb.ImmutableChat{
		Id:                     chatsDO.Id,
		Creator:                chatsDO.CreatorUserId,
		Title:                  chatsDO.Title,
		Photo:                  mtproto.MakeTLPhotoEmpty(nil).To_Photo(),
		Deactivated:            chatsDO.Deactivated,
		CallActive:             false,
		CallNotEmpty:           false,
		Noforwards:             chatsDO.Noforwards,
		ParticipantsCount:      chatsDO.ParticipantCount,
		Date:                   chatsDO.Date,
		Version:                chatsDO.Version,
		MigratedTo:             nil,
		DefaultBannedRights:    mtproto.BannedRights(chatsDO.DefaultBannedRights).ToChatBannedRights(),
		CanSetUsername:         false,
		About:                  chatsDO.About,
		ExportedInvite:         nil,
		BotInfo:                nil,
		Call:                   nil,
		AvailableReactionsType: chatsDO.AvailableReactionsType,
		AvailableReactions:     nil,
		TtlPeriod:              chatsDO.TtlPeriod,
	}

	if chatsDO.MigratedToId != 0 && chatsDO.MigratedToAccessHash != 0 {
		chat.MigratedTo = mtproto.MakeTLInputChannel(&mtproto.InputChannel{
			ChannelId:  chatsDO.MigratedToId,
			AccessHash: chatsDO.MigratedToAccessHash,
		}).To_InputChannel()
	}
	//
	////// chat_photo && photo
	//if chatsDO.PhotoId != 0 {
	//	chat.Photo, _ = d.MediaClient.MediaGetPhoto(ctx, &media.TLMediaGetPhoto{
	//		PhotoId: chatsDO.PhotoId,
	//	})
	//}
	//if chat.Photo == nil {
	//	chat.Photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	//}

	chat.ExportedInvite = nil // model.ExportedChatInviteEmpty

	if chatsDO.AvailableReactionsType == mtproto.ChatReactionsTypeSome {
		jsonx.UnmarshalFromString(chatsDO.AvailableReactions, &chat.AvailableReactions)
	}

	return
}

func (d *Dao) MakeImmutableChatParticipant(chatParticipantsDO *dataobject.ChatParticipantsDO) (participant *chatpb.ImmutableChatParticipant) {
	participant = &chatpb.ImmutableChatParticipant{
		Id:              chatParticipantsDO.Id,
		ChatId:          chatParticipantsDO.ChatId,
		UserId:          chatParticipantsDO.UserId,
		State:           chatParticipantsDO.State,
		ParticipantType: chatParticipantsDO.ParticipantType,
		Link:            chatParticipantsDO.Link,
		InviterUserId:   chatParticipantsDO.InviterUserId,
		InvitedAt:       chatParticipantsDO.InvitedAt,
		KickedAt:        chatParticipantsDO.KickedAt,
		LeftAt:          chatParticipantsDO.LeftAt,
		AdminRights:     nil,
		Date:            chatParticipantsDO.Date2,
	}

	if participant.ParticipantType == mtproto.ChatMemberAdmin {
		participant.AdminRights = mtproto.MakeDefaultChatAdminRights()
	}

	return
}
