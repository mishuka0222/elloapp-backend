package service

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

func (c *session) onSyncData(obj mtproto.TLObject) {
	// for android, obj maybe is nil
	if obj != nil {
		logx.Infof("genericSession]]>> - session: %s, syncData: %s", c, obj.DebugString())
	} else {
		logx.Infof("genericSession]]>> - session: %s, syncData: nil", c)
	}

	gatewayId := c.getGatewayId()

	if c.isAndroidPush {
		pusMsgId := c.authSessions.getNextNotifyId()
		c.sendPushToQueue(gatewayId, pusMsgId, androidPushTooLong)
	} else {
		pusMsgId := c.authSessions.getNextPushId()
		c.sendPushToQueue(gatewayId, pusMsgId, obj)
	}

	if c.sessionOnline() {
		if gatewayId == "" {
			logx.Errorf("gatewayId is empty, send delay...")
		} else {
			c.sendQueueToGateway(gatewayId)
		}

		//	syncMessage := &pendingMessage{
		//		messageId: nextMessageId(false),
		//		confirm:   true,
		//		// tl:        obj,
		//	}
		//	if c.isAndroidPush {
		//		syncMessage.tl = androidPushTooLong
		//	} else {
		//		syncMessage.tl = obj
		//	}
		//	c.syncMessages = append(c.syncMessages, syncMessage)
		//
		//	log.Debugf("genericSession]]>> - sendPending {sess: {%s}, pushObj: {%s}", c, reflect.TypeOf(obj))
		//	c.sendPendingMessagesToClient(c.getConnId(), c.syncMessages)
		//	c.syncMessages = []*pendingMessage{}
	}
}

func (c *session) onSyncRpcResultData(reqMsgId int64, data []byte) {
	// TODO(@benqi):
	// log.Debugf("genericSession]]>> - %v", cntl)
	c.pendingQueue.Remove(reqMsgId)
	gatewayId := c.getGatewayId()
	c.sendPushRpcResultToQueue(gatewayId, reqMsgId, data)
}

func (c *session) onSyncSessionData(obj mtproto.TLObject) {
	// TODO(@benqi):
	gatewayId := c.getGatewayId()
	pusMsgId := c.authSessions.getNextPushId()

	c.sendPushToQueue(gatewayId, pusMsgId, obj)
	c.sendQueueToGateway(gatewayId)
}
