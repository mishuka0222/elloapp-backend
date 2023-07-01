package service

import (
	"container/list"
	"time"
)

type httpReqItem struct {
	expireTime int64
	resChan    chan interface{}
}

type httpRequestQueue struct {
	q *list.List
}

func newHttpRequestQueue() *httpRequestQueue {
	return &httpRequestQueue{
		q: list.New(),
	}
}

func (q *httpRequestQueue) Push(resChan chan interface{}) {
	q.q.PushBack(&httpReqItem{
		expireTime: time.Now().Unix() + 25,
		resChan:    resChan,
	})
}

func (q *httpRequestQueue) Pop() chan interface{} {
	e := q.q.Front()
	if e != nil {
		q.q.Remove(e)
		return e.Value.(*httpReqItem).resChan
	}

	return nil
}

func (q *httpRequestQueue) PopTimeoutList() []chan interface{} {
	var rList []chan interface{}
	date := time.Now().Unix()
	for e := q.q.Front(); e != nil; e = e.Next() {
		if date >= e.Value.(*httpReqItem).expireTime {
			rList = append(rList, e.Value.(*httpReqItem).resChan)
			q.q.Remove(e)
		} else {
			break
		}
	}
	return rList
}

func (q *httpRequestQueue) Clear() {
	for e := q.q.Front(); e != nil; e = e.Next() {
		close(e.Value.(*httpReqItem).resChan)
	}
}
