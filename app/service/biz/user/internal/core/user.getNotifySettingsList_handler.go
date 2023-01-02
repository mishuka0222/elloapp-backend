package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
)

// UserGetNotifySettingsList
// user.getNotifySettingsList user_id:long peers:Vector<PeerUtil> = Vector<PeerNotifySettings>;
func (c *UserCore) UserGetNotifySettingsList(in *user.TLUserGetNotifySettingsList) (*user.Vector_PeerPeerNotifySettings, error) {
	var (
		settings = &user.Vector_PeerPeerNotifySettings{
			Datas: []*user.PeerPeerNotifySettings{},
		}
	)

	if _, err := c.svcCtx.Dao.UserNotifySettingsDAO.SelectListWithCB(c.ctx,
		in.UserId,
		in.Peers,
		func(i int, v *dataobject.UserNotifySettingsDO) {
			settings.Datas = append(settings.Datas, user.MakeTLPeerPeerNotifySettings(&user.PeerPeerNotifySettings{
				PeerType: v.PeerType,
				PeerId:   v.PeerId,
				Settings: makePeerNotifySettingsByDO(v),
			}).To_PeerPeerNotifySettings())
		}); err != nil {

		c.Logger.Errorf("user.getNotifySettingsList - error: %v", err)
		return nil, err
	}

	return settings, nil
}
