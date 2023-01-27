package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) ToChat(in *channels.ToChatReq) (res *channels.ToChatResp, err error) {
	var (
		forbidden = false
		admin     = false
		creator   = in.Channel.Channel.CreatorUserId == in.SelfUserId
		channel   *mtproto.TLChannel
	)
	for i := range in.Channel.Participants {
		if in.Channel.Participants[i].UserId == in.SelfUserId && in.Channel.Participants[i].State == 1 {
			forbidden = true
			break
		}
		if in.Channel.Participants[i].ParticipantType == channels.K_ChannelParticipantAdmin {
			admin = true
		}
	}

	if forbidden {
		var channel *mtproto.TLChannelForbidden
		channel = (&mtproto.Chat{
			Id:    in.Channel.Channel.Id,
			Title: in.Channel.Channel.Title,
		}).To_ChannelForbidden()
		res = &channels.ToChatResp{Chat: channel.To_Chat()}
		return
	}

	channel = (&mtproto.Chat{
		Creator:                 creator,
		Id:                      in.Channel.Channel.Id,
		Title:                   in.Channel.Channel.Title,
		ParticipantsCount_INT32: in.Channel.Channel.ParticipantCount,
		Date:                    in.Channel.Channel.Date,
		Version:                 in.Channel.Channel.Version,

		Left:        false,
		Deactivated: false,
		//CallActive:              m.CallActive(),
		//CallNotEmpty:            m.CallNotEmpty(),
		//Noforwards:              m.Noforwards(),

		MigratedTo:          nil,
		AdminRights:         nil,
		DefaultBannedRights: nil,
	}).To_Channel()

	if in.Channel.Channel.CreatorUserId == in.SelfUserId {
		AdminRights := mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
			ChangeInfo:     true,
			PostMessages:   true, // default false
			EditMessages:   true,
			DeleteMessages: true,
			BanUsers:       true,
			InviteUsers:    true,
			PinMessages:    true,
			AddAdmins:      true,
			Anonymous:      true,
			ManageCall:     true,
			Other:          true,
		}).To_ChatAdminRights()
		channel.SetAdminRights(AdminRights)
	} else if admin {
		AdminRights := mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
			ChangeInfo:     true,
			PostMessages:   false,
			EditMessages:   false,
			DeleteMessages: true,
			BanUsers:       true,
			InviteUsers:    true,
			PinMessages:    true,
			AddAdmins:      false,
			Anonymous:      false,
			ManageCall:     true,
			Other:          true,
		}).To_ChatAdminRights()
		channel.SetAdminRights(AdminRights)
		channel.SetBroadcast(true)
	} else {
		channel.SetDefaultBannedRights(mtproto.MakeDefaultBannedRights())
		channel.SetBroadcast(true)
	}

	if in.Channel.Channel.PhotoId == 0 {
		channel.SetPhoto(mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto())
	} else {
		var sizeList *media.PhotoSizeList

		sizeList, err = c.svcCtx.MediaClient.MediaGetPhotoSizeList(c.ctx, &media.TLMediaGetPhotoSizeList{
			Constructor: media.CRC32_media_getPhotoSizeList,
			SizeId:      in.Channel.Channel.PhotoId,
		})
		if err != nil {
			channel.SetPhoto(mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto())
		} else {
			channel.SetPhoto(mtproto.MakeChatPhotoByPhoto(&mtproto.Photo{
				PredicateName: mtproto.Predicate_photo,
				Id:            sizeList.SizeId,
				Sizes:         sizeList.Sizes,
			}))
		}
	}
	channel.SetPhoto(mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto())
	res = &channels.ToChatResp{Chat: channel.To_Chat()}

	return
}
