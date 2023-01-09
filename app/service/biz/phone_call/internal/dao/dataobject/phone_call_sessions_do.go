package dataobject

type PhoneCallSessionsDO struct {
	Id                    int64  `db:"id"`
	CallSessionId         int64  `db:"call_session_id"`
	AdminId               int64  `db:"admin_id"`
	AdminAccessHash       int64  `db:"admin_access_hash"`
	ParticipantId         int64  `db:"participant_id"`
	ParticipantAccessHash int64  `db:"participant_access_hash"`
	UdpP2P                bool   `db:"udp_p2p"`
	UdpReflector          bool   `db:"udp_reflector"`
	MinLayer              int32  `db:"min_layer"`
	MaxLayer              int32  `db:"max_layer"`
	GA                    string `db:"g_a"`
	GB                    string `db:"g_b"`
	State                 int32  `db:"state"`
	AdminDebugData        string `db:"admin_debug_data"`
	ParticipantDebugData  string `db:"participant_debug_data"`
	Date                  int32  `db:"date"`
	CreatedAt             string `db:"created_at"`
	UpdatedAt             string `db:"updated_at"`
}
