package dataobject

type AuthOpLogsDO struct {
	Id        int64  `db:"id"`
	AuthKeyId int64  `db:"auth_key_id"`
	Ip        string `db:"ip"`
	OpType    int32  `db:"op_type"`
	LogText   string `db:"log_text"`
}
