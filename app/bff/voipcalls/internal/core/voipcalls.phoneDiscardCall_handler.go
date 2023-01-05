package core

import (
	"github.com/gogo/protobuf/types"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *VoipcallsCore) PhoneDiscardCall(in *mtproto.TLPhoneDiscardCall) (*mtproto.Updates, error) {
	var (
		peer = in.GetPeer().To_InputPhoneCall()
	)
	callSession, err := c.svcCtx.Dao.PhonecallClient.MakePhoneCallSessionByLoad(c.ctx, &phonecall.TLMakePhoneCallSessionByLoad{SessionId: peer.GetId()})
	if err != nil {
		c.Logger.Errorf("invalid peer: {%v}, err: %v", peer, err)
		return nil, err
	}

	phoneCallDiscarded := &mtproto.TLPhoneCallDiscarded{Data2: &mtproto.PhoneCall{
		Id:        callSession.Id,
		NeedDebug: true,
		Reason:    in.GetReason(),
		Duration:  &types.Int32Value{Value: in.GetDuration()},
	}}
	c.Logger.Errorf("phoneCallDiscarded, err: %v", phoneCallDiscarded)
	//adminUpdatesData := (&mtproto.TLUpdatePhoneCall{
	//	Data2: &mtproto.Update{
	//		UserId: callSession.AdminId,
	//		PhoneCall: phoneCallDiscarded.To_PhoneCall(),
	//	},
	//}).To_Update()
	//participantUpdatesData := (&mtproto.TLUpdatePhoneCall{
	//	Data2: &mtproto.Update{
	//		UserId: callSession.ParticipantId,
	//		PhoneCall: phoneCallDiscarded.To_PhoneCall(),
	//	},
	//}).To_Update()
	//
	//action := &mtproto.TLMessageActionPhoneCall{Data2: &mtproto.MessageAction{
	//	CallId:   callSession.Id,
	//	Reason:   in.GetReason(),
	//	Duration: &types.Int32Value{Value: in.GetDuration()},
	//}}
	//peer2 := &mtproto.PeerUtil{
	//	PeerType: mtproto.PEER_USER,
	//	PeerId:   callSession.ParticipantId,
	//}

	return &mtproto.Updates{}, nil
}
