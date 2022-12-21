package core

import "github.com/teamgram/proto/mtproto"

func (c ConfigurationCore) BizInvokeBizDataRaw(in *mtproto.TLBizInvokeBizDataRaw) (*mtproto.BizDataRaw, error) {
	c.Logger.Debugf("BizInvokeBizDataRaw data: %+v", in.BizData.Data)

	return (&mtproto.BizDataRaw{Data: []byte("other messege")}).To_BizDataRaw().To_BizDataRaw(), nil
}
