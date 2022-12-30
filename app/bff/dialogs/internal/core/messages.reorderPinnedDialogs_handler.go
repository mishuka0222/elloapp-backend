package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesReorderPinnedDialogs
// messages.reorderPinnedDialogs#3b1adf37 flags:# force:flags.0?true folder_id:int order:Vector<InputDialogPeer> = Bool;
func (c *DialogsCore) MessagesReorderPinnedDialogs(in *mtproto.TLMessagesReorderPinnedDialogs) (*mtproto.Bool, error) {
	if len(in.GetOrder()) == 0 {
		c.Logger.Errorf("messages.reorderPinnedDialogs - len(order) == 0")
		return mtproto.BoolTrue, nil
	}

	var (
		peerDialogIdList []int64
	)
	for _, peer := range in.GetOrder() {
		switch peer.PredicateName {
		case mtproto.Predicate_inputDialogPeer:
			p := mtproto.FromInputPeer2(c.MD.UserId, peer.Peer)
			switch p.PeerType {
			case mtproto.PEER_SELF,
				mtproto.PEER_USER,
				mtproto.PEER_CHAT,
				mtproto.PEER_CHANNEL:
				peerDialogIdList = append(peerDialogIdList, p.PeerId)
			default:
				err := mtproto.ErrPeerIdInvalid
				c.Logger.Errorf("messages.reorderPinnedDialogs - error: %v", err)
				return nil, err
			}
		default:
			err := mtproto.ErrPeerIdInvalid
			c.Logger.Errorf("messages.reorderPinnedDialogs - error: %v", err)
			return nil, err
		}
	}

	_, err := c.svcCtx.Dao.DialogClient.DialogReorderPinnedDialogs(c.ctx, &dialog.TLDialogReorderPinnedDialogs{
		UserId:   c.MD.UserId,
		Force:    mtproto.ToBool(in.Force),
		FolderId: in.FolderId,
		IdList:   peerDialogIdList,
	})
	if err != nil {
		c.Logger.Errorf("messages.reorderPinnedDialogs - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
