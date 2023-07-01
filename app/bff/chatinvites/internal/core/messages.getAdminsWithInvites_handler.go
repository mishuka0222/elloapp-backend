package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetAdminsWithInvites
// messages.getAdminsWithInvites#3920e6ef peer:InputPeer = messages.ChatAdminsWithInvites;
func (c *ChatInvitesCore) MessagesGetAdminsWithInvites(in *mtproto.TLMessagesGetAdminsWithInvites) (*mtproto.Messages_ChatAdminsWithInvites, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
	)

	if !peer.IsChat() {
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.getAdminsWithInvites - error: ", err)
		return nil, err
	}

	rAdmins, err := c.svcCtx.Dao.ChatClient.ChatGetAdminsWithInvites(c.ctx, &chatpb.TLChatGetAdminsWithInvites{
		SelfId: c.MD.UserId,
		ChatId: peer.PeerId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getAdminsWithInvites - error: %v", err)
		return nil, err
	}

	rValues := mtproto.MakeTLMessagesChatAdminsWithInvites(&mtproto.Messages_ChatAdminsWithInvites{
		Admins: rAdmins.GetDatas(),
		Users:  []*mtproto.User{},
	}).To_Messages_ChatAdminsWithInvites()

	if len(rValues.Admins) == 0 {
		return rValues, nil
	}

	idHelper := mtproto.NewIDListHelper(c.MD.UserId)
	for _, a := range rValues.Admins {
		idHelper.AppendUsers(a.AdminId)
	}

	idHelper.Visit(func(userIdList []int64) {
		users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: userIdList,
		})
		rValues.Users = users.GetUserListByIdList(c.MD.UserId, userIdList...)
	}, nil, nil)

	return rValues, nil
}
