package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UsersGetUsers
// users.getUsers#d91a548 id:Vector<InputUser> = Vector<User>;
func (c *UsersCore) UsersGetUsers(in *mtproto.TLUsersGetUsers) (*mtproto.Vector_User, error) {
	idList := make([]int64, 0, len(in.Id))
	for _, inputUser := range in.Id {
		peer := mtproto.FromInputUser(c.MD.UserId, inputUser)
		switch peer.PeerType {
		case mtproto.PEER_SELF, mtproto.PEER_USER:
			idList = append(idList, peer.PeerId)
		default:
			c.Logger.Errorf("invalid userId")
		}
	}

	mUsers, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})

	return &mtproto.Vector_User{
		Datas: mUsers.GetUserListByIdList(c.MD.UserId, idList...),
	}, nil
}
