package dataobject

type UsernameDO struct {
	Id       int64  `db:"id"`
	Username string `db:"username"`
	PeerType int32  `db:"peer_type"`
	PeerId   int64  `db:"peer_id"`
	Deleted  bool   `db:"deleted"`
}
