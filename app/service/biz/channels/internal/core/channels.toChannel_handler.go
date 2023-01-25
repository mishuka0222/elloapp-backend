package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) ToChannel(in *channels.ToChannelReq) (res *channels.ToChannelResp, err error) {
	res = &channels.ToChannelResp{}
	var (
		forbidden = false
	)
	for i := range in.Channel.Participants {
		if in.Channel.Participants[i].UserId == in.SelfUserId && in.Channel.Participants[i].State == 1 {
			forbidden = true
			break
		}
	}

	if forbidden {
		var channel *mtproto.TLChannelForbidden
		channel = (&mtproto.Chat{
			Id:    in.Channel.Channel.Id,
			Title: in.Channel.Channel.Title,
		}).To_ChannelForbidden()
		res.Chat = channel.To_Chat()
		return
	}
	AdminRights := mtproto.MakeTLChatAdminRights(&mtproto.ChatAdminRights{
		ChangeInfo:     true,
		PostMessages:   true, // default false
		EditMessages:   false,
		DeleteMessages: true,
		BanUsers:       true,
		InviteUsers:    true,
		PinMessages:    true,
		AddAdmins:      in.Channel.Channel.AdminsEnabled == 1,
		Anonymous:      false,
		ManageCall:     true,
		Other:          true,
	}).To_ChatAdminRights()
	var channel *mtproto.TLChannel
	channel = (&mtproto.Chat{
		Creator:     in.Channel.Channel.CreatorUserId == in.SelfUserId,
		Id:          in.Channel.Channel.Id,
		Title:       in.Channel.Channel.Title,
		AdminRights: AdminRights,
		// Photo:             mtproto.NewTLChatPhotoEmpty().To_ChatPhoto(),
		ParticipantsCount_INT32: in.Channel.Channel.ParticipantCount,
		Date:                    in.Channel.Channel.Date,
		Version:                 in.Channel.Channel.Version,
	}).To_Channel()

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
	res.Chat = channel.To_Chat()

	return
}
