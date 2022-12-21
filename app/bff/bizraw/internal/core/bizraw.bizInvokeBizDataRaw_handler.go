package core

import "github.com/teamgram/proto/mtproto"

func (c ConfigurationCore) BizInvokeBizDataRaw(in *mtproto.TLBizInvokeBizDataRaw) (*mtproto.BizDataRaw, error) {
	c.Logger.Debugf("BizInvokeBizDataRaw data: %+v", string(in.BizData.Data))

	return (&mtproto.BizDataRaw{
		Constructor: mtproto.CRC32_bizDataRaw,
		Data:        []byte("other messege")}).To_BizDataRaw().To_BizDataRaw(), nil
}
