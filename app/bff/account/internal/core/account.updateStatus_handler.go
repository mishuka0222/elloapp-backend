package core

import (
	"github.com/zeromicro/go-zero/core/contextx"
	"github.com/zeromicro/go-zero/core/threading"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"time"
)

// AccountUpdateStatus
// account.updateStatus#6628562c offline:Bool = Bool;
func (c *AccountCore) AccountUpdateStatus(in *mtproto.TLAccountUpdateStatus) (*mtproto.Bool, error) {
	var (
		offline     = mtproto.FromBool(in.GetOffline())
		now         = time.Now().Unix()
		expires     = int32(0)
		pushUpdates = mtproto.MakeTLUpdateShort(&mtproto.Updates{
			Update: mtproto.MakeTLUpdateUserStatus(&mtproto.Update{
				UserId: c.MD.UserId,
				Status: nil,
			}).To_Update(),
			Date: int32(now) - 1,
		}).To_Updates()
	)

	if !offline {
		expires = 300
		pushUpdates.Update.Status = mtproto.MakeTLUserStatusOnline(&mtproto.UserStatus{
			Expires: int32(now) + expires,
		}).To_UserStatus()

		// online
		// s.UserFacade.UpdateUserStatus(ctx, md.UserId, now)
		// c.svcCtx.Dao.UserClient.UserUpdateLastSeen()
	} else {
		pushUpdates.Update.Status = mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
			WasOnline: int32(now),
		}).To_UserStatus()
	}

	// threading.GoSafe()
	threading.GoSafe(func() {
		c.svcCtx.Dao.UserClient.UserUpdateLastSeen(
			contextx.ValueOnlyFrom(c.ctx),
			&userpb.TLUserUpdateLastSeen{
				Id:         c.MD.UserId,
				LastSeenAt: now,
				Expires:    expires,
			})
	})

	// TODO: push
	//// log.Debugf("account.updateStatus - reply: {true}")
	//return model.WrapperGoFunc(mtproto.BoolTrue, func() {
	//	// log.Debugf("ready push to other contacts...")
	//	// push to other contacts.
	//	contactIdList := s.UserFacade.GetContactUserIdList(context.Background(), md.UserId)
	//	blockedIdList := s.UserFacade.CheckBlockUserList(context.Background(), md.UserId, contactIdList)
	//
	//	// TODO(@benqi): push updateUserStatus规则
	//	for _, id := range contactIdList {
	//		if md.UserId == id {
	//			// why??
	//			continue
	//		}
	//
	//		if blocked, _ := container2.Contains(id, blockedIdList); blocked {
	//			continue
	//		}
	//
	//		// log.Debugf("check blocked...")
	//		blocked := s.UserFacade.IsBlockedByUser(context.Background(), md.UserId, id)
	//		if blocked {
	//			continue
	//		}
	//
	//		sync_client.PushUpdates(context.Background(), id, pushUpdates)
	//	}
	//
	//}).(*mtproto.Bool), nil

	return mtproto.BoolTrue, nil
}
