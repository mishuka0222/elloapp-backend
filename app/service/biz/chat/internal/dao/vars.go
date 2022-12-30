package dao

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
)

type idxId struct {
	idx int
	id  int64
}

type kv struct {
	k string
	v interface{}
}

func removeAllNil(participants []*chatpb.ImmutableChatParticipant) []*chatpb.ImmutableChatParticipant {
	for i := 0; i < len(participants); {
		if participants[i] != nil {
			i++
			continue
		}

		if i < len(participants)-1 {
			copy(participants[i:], participants[i+1:])
		}

		participants[len(participants)-1] = nil
		participants = participants[:len(participants)-1]
	}

	return participants
}
