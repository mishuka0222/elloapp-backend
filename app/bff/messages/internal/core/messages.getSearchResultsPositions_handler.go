// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
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
