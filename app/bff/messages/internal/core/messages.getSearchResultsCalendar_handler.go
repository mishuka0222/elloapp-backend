package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetSearchResultsCalendar
// messages.getSearchResultsCalendar#49f0bde9 peer:InputPeer filter:MessagesFilter offset_id:int offset_date:int = messages.SearchResultsCalendar;
func (c *MessagesCore) MessagesGetSearchResultsCalendar(in *mtproto.TLMessagesGetSearchResultsCalendar) (*mtproto.Messages_SearchResultsCalendar, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.getSearchResultsCalendar blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
