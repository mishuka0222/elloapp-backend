package server

const (
	STATE_ERROR              = 0x0000
	STATE_CONNECTED2         = 0x0100
	STATE_HANDSHAKE          = 0x0200
	STATE_pq                 = 0x0201
	STATE_pq_res             = 0x0202
	STATE_pq_ack             = 0x0203
	STATE_DH_params          = 0x0204
	STATE_DH_params_res      = 0x0205
	STATE_DH_params_res_fail = 0x0206
	STATE_DH_params_ack      = 0x0207
	STATE_dh_gen             = 0x0208
	STATE_dh_gen_res         = 0x0209
	STATE_dh_gen_res_retry   = 0x020a
	STATE_dh_gen_res_fail    = 0x020b
	STATE_dh_gen_ack         = 0x020c
	STATE_AUTH_KEY           = 0x0300
)

const (
	RES_STATE_NONE  = 0x00
	RES_STATE_OK    = 0x01
	RES_STATE_ERROR = 0x02
)
