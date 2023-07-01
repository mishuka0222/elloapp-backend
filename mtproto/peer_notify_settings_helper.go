package mtproto

import (
	"github.com/gogo/protobuf/types"
)

var (
	DefaultPeerNotifySettings = MakeDefaultPeerNotifySettings()
)

func MakePeerNotifySettings(settings *InputPeerNotifySettings) (*PeerNotifySettings, error) {
	notifySettings := MakeTLPeerNotifySettings(&PeerNotifySettings{
		ShowPreviews: settings.GetShowPreviews(),
		Silent:       settings.GetSilent(),
		MuteUntil:    settings.GetMuteUntil(),
		Sound:        settings.GetSound_FLAGSTRING(),
	}).To_PeerNotifySettings()

	return notifySettings, nil
}

// MakeDefaultPeerNotifySettings
/*
   { updateNotifySettings
     peer: { notifyChats },
     notify_settings: { peerNotifySettings
       flags: 15 [INT],
       show_previews: { boolTrue },
       silent: { boolFalse },
       mute_until: 0 [INT],
       sound: "100" [STRING],
     },
   },
*/
func MakeDefaultPeerNotifySettings() *PeerNotifySettings {
	settings := MakeTLPeerNotifySettings(&PeerNotifySettings{
		ShowPreviews: ToBool(true),
		Silent:       ToBool(true),
		MuteUntil:    &types.Int32Value{Value: 0},
		Sound:        &types.StringValue{Value: "default"},
	}).To_PeerNotifySettings()

	return settings
}
