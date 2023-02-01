package dao

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// GetOffsetIdBackwardHistoryMessages offset
func (d *Dao) GetOffsetIdBackwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT, mtproto.PEER_CHANNEL:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectBackwardByOffsetIdLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			offsetId,
			limit,
			func(i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
		// logx.WithContext(ctx).Infof("GetOffsetIdBackwardHistoryMessages: %v", rList)
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetIdForwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT, mtproto.PEER_CHANNEL:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectForwardByOffsetIdLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			offsetId,
			limit,
			func(i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetDateBackwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetDate, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT, mtproto.PEER_CHANNEL:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectBackwardByOffsetDateLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			int64(offsetDate),
			limit,
			func(i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetDateForwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetDate, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT, mtproto.PEER_CHANNEL:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectForwardByOffsetDateLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			int64(offsetDate),
			limit,
			func(i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

// GetOffsetIdBackwardUnreadMentions GetOffsetIdBackwardUnreadMentions
func (d *Dao) GetOffsetIdBackwardUnreadMentions(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_CHAT, mtproto.PEER_CHANNEL:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectBackwardUnreadMentionsByOffsetIdLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			offsetId,
			limit,
			func(i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	}
	return
}

func (d *Dao) GetOffsetIdForwardUnreadMentions(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_CHAT, mtproto.PEER_CHANNEL:
		var (
			did = mtproto.MakeDialogId(userId, peer.PeerType, peer.PeerId)
		)

		rList, _ := d.MessagesDAO.SelectForwardUnreadMentionsByOffsetIdLimitWithCB(
			ctx,
			userId,
			did.A,
			did.B,
			offsetId,
			limit,
			func(i int, v *dataobject.MessagesDO) {
				messages = append(messages, d.MakeMessageBox(ctx, userId, v))
			})
		_ = rList
	}
	return
}
