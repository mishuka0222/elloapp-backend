package mtproto

import (
	"github.com/gogo/protobuf/types"
	"time"
)

func (m *Document) GetFixedSize() int64 {
	if m.Size2_INT64 != 0 {
		return m.Size2_INT64
	}
	if m.Size2_INT32 != 0 {
		return int64(m.Size2_INT32)
	}
	if m.Size2 != 0 {
		return m.Size2
	}

	return 0
}

func (m *Document) FixData() *Document {
	if m.Size2_INT64 == 0 {
		m.Size2_INT64 = m.GetFixedSize()
	}
	if m.Size2_INT32 == 0 {
		m.Size2_INT32 = int32(m.Size2_INT64)
	}

	return m
}

func (m *PollResults) FixData() *PollResults {
	if m.Solution != nil {
		if m.SolutionEntities == nil {
			m.SolutionEntities = []*MessageEntity{}
		}
	}
	return m
}

func (m *Update) FixData() *Update {
	if m.GetMedia() != nil {
		m.Media = m.Media.FixData()
	}
	if m.GetMessage_MESSAGE() != nil {
		m.Message_MESSAGE = m.Message_MESSAGE.FixData()
	}

	return m
}

func (m *MessageMedia) FixData() *MessageMedia {
	if m.GetDocument() != nil {
		m.Document = m.Document.FixData()
	}

	if m.GetResults() != nil {
		m.Results = m.Results.FixData()
	}

	return m
}

func (m *Message) FixData() *Message {
	if m.GetMedia() != nil {
		m.Media = m.Media.FixData()
	}

	if m.GetReactions() != nil {
		m.Reactions = m.Reactions.FixData()
	}

	if m.GetTtlPeriod() != nil {
		ttlPeriod := m.GetTtlPeriod().GetValue() + m.GetDate() - int32(time.Now().Unix())
		if ttlPeriod < 0 {
			ttlPeriod = 0
		}
		m.TtlPeriod = &types.Int32Value{Value: ttlPeriod}
	}

	return m
}

func (m *EncryptedFile) GetFixedSize() int64 {
	if m.Size2_INT64 != 0 {
		return m.Size2_INT64
	}
	if m.Size2_INT32 != 0 {
		return int64(m.Size2_INT32)
	}

	return 0
}

func (m *ReactionCount) FixData() *ReactionCount {
	if m.Reaction != "" {
		if m.Reaction_STRING == "" {
			m.Reaction_STRING = m.Reaction
		}
		if m.Reaction_REACTION == nil {
			m.Reaction_REACTION = MakeTLReactionEmoji(&Reaction{
				Emoticon: m.Reaction,
			}).To_Reaction()
		}
	}

	if m.Reaction_STRING != "" {
		if m.Reaction_REACTION == nil {
			m.Reaction_REACTION = MakeTLReactionEmoji(&Reaction{
				Emoticon: m.Reaction_STRING,
			}).To_Reaction()
		}
	}

	if m.Reaction_REACTION != nil {
		if m.Reaction_STRING == "" {
			m.Reaction_STRING = m.Reaction_REACTION.GetEmoticon()
		}
	}

	return m
}

func (m *MessagePeerReaction) FixData() *MessagePeerReaction {
	if m.Reaction != "" {
		if m.Reaction_STRING == "" {
			m.Reaction_STRING = m.Reaction
		}
		if m.Reaction_REACTION == nil {
			m.Reaction_REACTION = MakeTLReactionEmoji(&Reaction{
				Emoticon: m.Reaction,
			}).To_Reaction()
		}
	}

	if m.Reaction_STRING != "" {
		if m.Reaction_REACTION == nil {
			m.Reaction_REACTION = MakeTLReactionEmoji(&Reaction{
				Emoticon: m.Reaction_STRING,
			}).To_Reaction()
		}
	}

	if m.Reaction_REACTION != nil {
		if m.Reaction_STRING == "" {
			m.Reaction_STRING = m.Reaction_REACTION.GetEmoticon()
		}
	}

	return m
}

func (m *MessageReactions) FixData() *MessageReactions {
	for _, v := range m.Results {
		v.FixData()
	}
	for _, v := range m.RecentReactions {
		v.FixData()
	}

	return m
}
