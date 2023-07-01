package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"time"

	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/updates"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UpdatesGetDifference
// updates.getDifference#25939651 flags:# pts:int pts_total_limit:flags.0?int date:int qts:int = updates.Difference;
func (c *UpdatesCore) UpdatesGetDifference(in *mtproto.TLUpdatesGetDifference) (*mtproto.Updates_Difference, error) {
	keyId, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionGetPermAuthKeyId(c.ctx, &authsession.TLAuthsessionGetPermAuthKeyId{
		AuthKeyId: c.MD.AuthId,
	})
	if err != nil {
		c.Logger.Errorf("updates.getDifference - error: %v", err)
		return nil, err
	}
	c.Logger.Infof("updates.getDifference - keyId: %v", keyId)

	updatesDiff, err := c.svcCtx.Dao.UpdatesClient.UpdatesGetDifferenceV2(c.ctx, &updates.TLUpdatesGetDifferenceV2{
		AuthKeyId:     keyId.GetV(),
		UserId:        c.MD.UserId,
		Pts:           in.Pts,
		PtsTotalLimit: in.PtsTotalLimit,
		Date:          int64(in.Date) - 1,
	})
	if err != nil {
		c.Logger.Errorf("updates.getDifference - error: %v", err)
		return nil, err
	}

	var (
		idHelper    = mtproto.NewIDListHelper(c.MD.UserId)
		rDifference *mtproto.Updates_Difference
	)

	switch updatesDiff.GetPredicateName() {
	case updates.Predicate_differenceEmpty:
		return mtproto.MakeTLUpdatesDifferenceEmpty(&mtproto.Updates_Difference{
			Date: updatesDiff.GetState().GetDate(),
			Seq:  updatesDiff.GetState().GetSeq(),
		}).To_Updates_Difference(), nil
	case updates.Predicate_difference:
		// TODO: fix date
		updatesDiff.State.Date = int32(time.Now().Unix())

		rDifference = mtproto.MakeTLUpdatesDifference(&mtproto.Updates_Difference{
			NewMessages:          updatesDiff.NewMessages,
			NewEncryptedMessages: []*mtproto.EncryptedMessage{},
			OtherUpdates:         updatesDiff.OtherUpdates,
			Chats:                nil,
			Users:                nil,
			State:                updatesDiff.State,
		}).To_Updates_Difference()
	case updates.Predicate_differenceSlice:
		rDifference = mtproto.MakeTLUpdatesDifferenceSlice(&mtproto.Updates_Difference{
			NewMessages:          updatesDiff.NewMessages,
			NewEncryptedMessages: []*mtproto.EncryptedMessage{},
			OtherUpdates:         updatesDiff.OtherUpdates,
			Chats:                nil,
			Users:                nil,
			IntermediateState:    updatesDiff.IntermediateState,
		}).To_Updates_Difference()

		// TODO: fix IntermediateState
	case updates.Predicate_differenceTooLong:
		// TODO: iOS
		rDifference = mtproto.MakeTLUpdatesDifferenceTooLong(&mtproto.Updates_Difference{
			Pts: updatesDiff.Pts,
		}).To_Updates_Difference()
	default:
	}

	idHelper.PickByMessages(rDifference.NewMessages...)
	idHelper.PickByUpdates(rDifference.OtherUpdates...)

	idHelper.Visit(
		func(userIdList []int64) {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: append(userIdList, c.MD.UserId),
				})
			rDifference.Users = users.GetUserListByIdList(c.MD.UserId, userIdList...)
		},
		func(chatIdList []int64) {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: chatIdList,
				})
			rDifference.Chats = chats.GetChatListByIdList(c.MD.UserId, chatIdList...)
		},
		func(channelIdList []int64) {
		})

	return rDifference, nil
}
