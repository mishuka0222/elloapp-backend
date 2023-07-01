package dialog

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type DialogsDataHelper struct {
	Dialogs  []*mtproto.Dialog
	Messages []*mtproto.Message
	Chats    []*mtproto.Chat
	Users    []*mtproto.User
}

func (m *DialogsDataHelper) ToMessagesDialogs(count int32) *mtproto.Messages_Dialogs {
	return mtproto.MakeTLMessagesDialogsSlice(&mtproto.Messages_Dialogs{
		Dialogs:  m.Dialogs,
		Messages: m.Messages,
		Chats:    m.Chats,
		Users:    m.Users,
		Count:    count,
	}).To_Messages_Dialogs()
}

func (m *DialogsDataHelper) ToMessagesPeerDialogs(state *mtproto.Updates_State) *mtproto.Messages_PeerDialogs {
	return mtproto.MakeTLMessagesPeerDialogs(&mtproto.Messages_PeerDialogs{
		Dialogs:  m.Dialogs,
		Messages: m.Messages,
		Users:    m.Users,
		Chats:    m.Chats,
		State:    state,
	}).To_Messages_PeerDialogs()
}
