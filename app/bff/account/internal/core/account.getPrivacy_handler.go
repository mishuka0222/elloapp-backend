package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountGetPrivacy
// account.getPrivacy#dadbc950 key:InputPrivacyKey = account.PrivacyRules;
func (c *AccountCore) AccountGetPrivacy(in *mtproto.TLAccountGetPrivacy) (*mtproto.Account_PrivacyRules, error) {
	var (
		key  = userpb.FromInputPrivacyKeyType(in.Key)
		rVal *mtproto.Account_PrivacyRules
	)

	// TODO(@benqi): Check request valid.
	if key == userpb.KEY_TYPE_INVALID {
		err := mtproto.ErrPrivacyKeyInvalid
		c.Logger.Errorf("account.getPrivacy - error: %v", err)
		return nil, err
	}

	ruleList, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
		UserId:  c.MD.UserId,
		KeyType: int32(key),
	})

	if len(ruleList.GetDatas()) == 0 {
		if key == userpb.PHONE_NUMBER {
			rVal = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
				Rules: []*mtproto.PrivacyRule{mtproto.MakeTLPrivacyValueDisallowAll(nil).To_PrivacyRule()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
			}).To_Account_PrivacyRules()
		} else {
			rVal = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
				Rules: []*mtproto.PrivacyRule{mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
			}).To_Account_PrivacyRules()
		}
	} else {
		rVal = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
			Rules: ruleList.GetDatas(),
			Users: []*mtproto.User{}, // TODO
			Chats: []*mtproto.Chat{}, // TODO
		}).To_Account_PrivacyRules()

		idHelper := mtproto.NewIDListHelper(c.MD.UserId)
		idHelper.PickByRules(ruleList.GetDatas())
		idHelper.Visit(
			func(userIdList []int64) {
				users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
					&userpb.TLUserGetMutableUsers{
						Id: userIdList,
					})
				rVal.Users = users.GetUserListByIdList(c.MD.UserId, userIdList...)
			},
			func(chatIdList []int64) {
				chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
					&chatpb.TLChatGetChatListByIdList{
						IdList: chatIdList,
					})
				rVal.Chats = chats.GetChatListByIdList(c.MD.UserId, chatIdList...)
			},

			func(channelIdList []int64) {
				// TODO
			})
	}

	return rVal, nil
}
