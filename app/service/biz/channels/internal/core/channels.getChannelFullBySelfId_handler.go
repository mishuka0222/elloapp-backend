package core

import (
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/media"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"time"
)

func (c *ChannelsCore) GetChannelFullBySelfId(in *channels.GetChannelFullBySelfIdReq) (res *channels.GetChannelFullBySelfIdResp, err error) {

	res = &channels.GetChannelFullBySelfIdResp{}

	var (
		photo          *mtproto.Photo
		isAdmin        *channels.CheckUserIsAdministratorResp
		notifySettings *mtproto.PeerNotifySettings
		sizes          *media.PhotoSizeList
	)
	sizes, err = c.svcCtx.MediaClient.MediaGetPhotoSizeList(c.ctx, &media.TLMediaGetPhotoSizeList{
		Constructor: media.CRC32_media_getPhotoSizeList,
		SizeId:      in.ChannelData.Channel.PhotoId,
	})
	// photo2 := photo2.MakeUserProfilePhoto(photoId, sizes)

	if in.ChannelData.Channel.GetPhotoId() == 0 {
		photoEmpty := (&mtproto.Photo{
			Id: 0,
		}).To_PhotoEmpty()
		photo = photoEmpty.To_Photo()
	} else {
		channelPhoto := (&mtproto.Photo{
			Id:          in.ChannelData.Channel.PhotoId,
			HasStickers: false,
			AccessHash:  in.ChannelData.Channel.PhotoId, // photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
			Date:        int32(time.Now().Unix()),
			Sizes:       sizes.Sizes,
		}).To_Photo()
		photo = channelPhoto.To_Photo()
	}

	peer := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_CHANNEL,
		PeerId:   in.ChannelData.Channel.GetId(),
	}
	notifySettings, err = c.svcCtx.UserService.UserGetNotifySettings(c.ctx, &userpb.TLUserGetNotifySettings{
		UserId:   c.MD.UserId,
		PeerType: peer.PeerType,
		PeerId:   peer.PeerId,
	})

	res.Channel = (&mtproto.ChatFull{
		Id:                in.ChannelData.Channel.Id,
		About:             in.ChannelData.Channel.Title,
		ParticipantsCount: &types.Int32Value{Value: in.ChannelData.Channel.ParticipantCount},
		AdminsCount:       &types.Int32Value{Value: 1}, // TODO: write logic: calc admins count
		ChatPhoto:         photo,
		NotifySettings:    notifySettings,
		ExportedInvite:    mtproto.MakeTLChatInviteExported(nil).To_ExportedChatInvite(), // TODO: write logic
		BotInfo:           []*mtproto.BotInfo{},
	}).To_ChannelFull()

	isAdmin, err = c.CheckUserIsAdministrator(&channels.CheckUserIsAdministratorReq{Channel: in.ChannelData, UserId: in.SelfUserId})
	if isAdmin.Status {
		res.Channel.SetCanViewParticipants(true)
		res.Channel.SetCanSetUsername(true)
	}

	exportedInvite := &mtproto.TLChatInviteExported{Data2: &mtproto.ExportedChatInvite{
		Link: in.ChannelData.Channel.Link,
	}}

	res.Channel.SetExportedInvite(exportedInvite.To_ExportedChatInvite())

	return
}
