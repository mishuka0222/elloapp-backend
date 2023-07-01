package core

import (
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func makePeerNotifySettingsByDO(do *dataobject.UserNotifySettingsDO) (settings *mtproto.PeerNotifySettings) {
	settings = mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
	if do.ShowPreviews != -1 {
		settings.ShowPreviews = mtproto.ToBool(do.ShowPreviews == 1)
	}
	if do.Silent != -1 {
		settings.Silent = mtproto.ToBool(do.Silent == 1)
	}
	if do.MuteUntil != -1 {
		settings.MuteUntil = &types.Int32Value{Value: do.MuteUntil}
	}
	if do.Sound != "-1" {
		settings.Sound = &types.StringValue{Value: do.Sound}
	}
	return
}

// UserGetAllNotifySettings
// user.getAllNotifySettings user_id:int = Vector<PeerNotifySettings>;
func (c *UserCore) UserGetAllNotifySettings(in *user.TLUserGetAllNotifySettings) (*user.Vector_PeerPeerNotifySettings, error) {
	var (
		settings = &user.Vector_PeerPeerNotifySettings{
			Datas: []*user.PeerPeerNotifySettings{},
		}
	)

	if _, err := c.svcCtx.Dao.UserNotifySettingsDAO.SelectAllWithCB(c.ctx,
		in.UserId,
		func(i int, v *dataobject.UserNotifySettingsDO) {
			settings.Datas = append(settings.Datas, user.MakeTLPeerPeerNotifySettings(&user.PeerPeerNotifySettings{
				PeerType: v.PeerType,
				PeerId:   v.PeerId,
				Settings: makePeerNotifySettingsByDO(v),
			}).To_PeerPeerNotifySettings())
		}); err != nil {

		c.Logger.Errorf("user.getAllNotifySettings - error: %v", err)
		return nil, err
	}

	return settings, nil
}
