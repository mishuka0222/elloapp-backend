package user

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

func FindPeerPeerNotifySettings(settingsList []*PeerPeerNotifySettings, peer *mtproto.PeerUtil) *mtproto.PeerNotifySettings {
	for _, s := range settingsList {
		if s.PeerType == peer.PeerType && s.PeerId == peer.PeerId {
			return s.Settings
		}
	}

	return mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
}
