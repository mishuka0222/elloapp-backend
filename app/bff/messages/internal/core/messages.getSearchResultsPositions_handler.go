package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetSearchResultsPositions
// messages.getSearchResultsPositions#6e9583a3 peer:InputPeer filter:MessagesFilter offset_id:int limit:int = messages.SearchResultsPositions;
func (c *MessagesCore) MessagesGetSearchResultsPositions(in *mtproto.TLMessagesGetSearchResultsPositions) (*mtproto.Messages_SearchResultsPositions, error) {
	// TODO: not impl

	return mtproto.MakeTLMessagesSearchResultsPositions(&mtproto.Messages_SearchResultsPositions{
		Count:     0,
		Positions: []*mtproto.SearchResultsPosition{},
	}).To_Messages_SearchResultsPositions(), nil
}
