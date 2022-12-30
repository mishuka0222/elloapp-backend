package dao

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

// GetOffsetIdBackwardHistoryMessages offset
func (d *Dao) GetOffsetIdBackwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
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
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetIdForwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
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
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetDateBackwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetDate, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
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
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

func (d *Dao) GetOffsetDateForwardHistoryMessages(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetDate, minId, maxId, limit int32, hash int64) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
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
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
	}

	messages = mtproto.ToSafeMessageBoxList(messages)
	return
}

// GetOffsetIdBackwardUnreadMentions GetOffsetIdBackwardUnreadMentions
func (d *Dao) GetOffsetIdBackwardUnreadMentions(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_CHAT:
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
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
	}
	return
}

func (d *Dao) GetOffsetIdForwardUnreadMentions(ctx context.Context, userId int64, peer *mtproto.PeerUtil, offsetId, minId, maxId, limit int32) (messages []*mtproto.MessageBox) {
	switch peer.PeerType {
	case mtproto.PEER_CHAT:
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
	case mtproto.PEER_CHANNEL:
		logx.Errorf("blocked, License key from https://elloapp.com required to unlock enterprise features.")
	}
	return
}
