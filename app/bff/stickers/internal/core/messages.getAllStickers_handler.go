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
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetAllStickers
// messages.getAllStickers#b8a0a1a8 hash:long = messages.AllStickers;
func (c *StickersCore) MessagesGetAllStickers(in *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.getAllStickers blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return mtproto.MakeTLMessagesAllStickers(&mtproto.Messages_AllStickers{
		Hash: in.Hash,
		Sets: []*mtproto.StickerSet{},
	}).To_Messages_AllStickers(), nil
}
