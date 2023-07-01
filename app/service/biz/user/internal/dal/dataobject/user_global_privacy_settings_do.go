package dataobject

type UserGlobalPrivacySettingsDO struct {
	Id                               int64 `db:"id"`
	UserId                           int64 `db:"user_id"`
	ArchiveAndMuteNewNoncontactPeers bool  `db:"archive_and_mute_new_noncontact_peers"`
}
