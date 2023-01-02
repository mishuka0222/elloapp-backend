package service

import (
	"container/list"
	"math"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

const (
	maxAckedIdListSize = 500
)

type outboxMsg struct {
	msgId int64 // 1. req_msg_id; 2. 0; 3. < math.MaxInt32
	sent  int64
	state byte
	msg   *mtproto.TLMessageRawData
}

type sessionOutgoingQueue struct {
	minMsgId    int64
	maxMsgId    int64
	oMsgs       *list.List
	ackedIdList []int64
}

func newSessionOutgoingQueue() *sessionOutgoingQueue {
	return &sessionOutgoingQueue{
		minMsgId:    0,
		maxMsgId:    math.MaxInt64,
		oMsgs:       list.New(),
		ackedIdList: make([]int64, 0),
	}
}

func (q *sessionOutgoingQueue) AddRpcResultMsg(reqMsgId int64, result *mtproto.TLMessageRawData) *outboxMsg {
	oMsg := q.Lookup(reqMsgId)
	if oMsg == nil {
		oMsg = &outboxMsg{
			msgId: reqMsgId,
			sent:  0,
			state: ACKNOWLEDGED,
			msg:   result,
		}
		q.oMsgs.PushBack(oMsg)
	}

	q.Shrink()
	return oMsg
}

func (q *sessionOutgoingQueue) AddNotifyMsg(notifyId int64, confirm bool, msg *mtproto.TLMessageRawData) *outboxMsg {
	oMsg := new(outboxMsg)
	oMsg.msgId = notifyId
	oMsg.sent = 0
	if confirm {
		oMsg.state = ACKNOWLEDGED
	} else {
		oMsg.state = NEED_NO_ACK
	}
	oMsg.msg = msg
	q.oMsgs.PushBack(oMsg)

	q.Shrink()
	return oMsg
}

func (q *sessionOutgoingQueue) AddPushUpdates(pushMsgId int64, result *mtproto.TLMessageRawData) *outboxMsg {
	oMsg := q.Lookup(pushMsgId)
	if oMsg == nil {
		oMsg = &outboxMsg{
			msgId: pushMsgId,
			sent:  0,
			state: ACKNOWLEDGED,
			msg:   result,
		}
		q.oMsgs.PushBack(oMsg)
	}

	q.Shrink()
	return oMsg
}

func (q *sessionOutgoingQueue) OnMsgsAck(ackIds []int64, cb func(inMsgId int64)) {
	var next *list.Element
	for _, id := range ackIds {
		for e := q.oMsgs.Front(); e != nil; e = next {
			next = e.Next()
			if id == e.Value.(*outboxMsg).msg.MsgId {
				iMsgId := e.Value.(*outboxMsg).msgId
				q.ackedIdList = append(q.ackedIdList, iMsgId)
				cb(iMsgId)
				q.oMsgs.Remove(e)
			}
		}
	}

	if len(q.ackedIdList) > maxAckedIdListSize {
		q.ackedIdList = q.ackedIdList[len(q.ackedIdList)-maxAckedIdListSize-1:]
	}
}

func (q *sessionOutgoingQueue) Lookup(msgId int64) (oMsg *outboxMsg) {
	for e := q.oMsgs.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*outboxMsg).msgId {
			oMsg = e.Value.(*outboxMsg)
		}
	}
	return
}

func (q *sessionOutgoingQueue) Remove(msgId int64) (oMsg *outboxMsg) {
	for e := q.oMsgs.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*outboxMsg).msgId {
			oMsg = e.Value.(*outboxMsg)
			q.oMsgs.Remove(e)
			break
		}
	}
	return
}

func (q *sessionOutgoingQueue) Shrink() {
	for q.oMsgs.Len() > maxQueueSize {
		iMsg := q.oMsgs.Remove(q.oMsgs.Front())
		q.minMsgId = iMsg.(*outboxMsg).msgId
	}
}
