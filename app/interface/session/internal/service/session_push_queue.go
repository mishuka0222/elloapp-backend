package service

import (
	"container/list"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"

	"github.com/zeromicro/go-zero/core/logx"
)

type pushMsg struct {
	msgId int64
	msg   mtproto.TLObject
}

type sessionPushQueue struct {
	q *list.List
}

func newSessionPushQueue() *sessionPushQueue {
	return &sessionPushQueue{
		q: list.New(),
	}
}

func (q *sessionPushQueue) Add(msgId int64, pushMsg2 mtproto.TLObject) {
	logx.Infof("add msgId: %d", msgId)
	q.q.PushBack(&pushMsg{
		msgId: msgId,
		msg:   pushMsg2,
	})
}

func (q *sessionPushQueue) Remove(msgId int64) {
	logx.Infof("remove msgId: %d", msgId)
	for e := q.q.Front(); e != nil; e = e.Next() {
		if msgId == e.Value.(*pushMsg).msgId {
			q.q.Remove(e)
			break
		}
	}
}
