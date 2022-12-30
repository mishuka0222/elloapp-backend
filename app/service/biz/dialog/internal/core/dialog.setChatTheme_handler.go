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
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// DialogSetChatTheme
// dialog.setChatTheme user_id:long peer_type:int peer_id:long theme_emoticon:string = Bool;
func (c *DialogCore) DialogSetChatTheme(in *dialog.TLDialogSetChatTheme) (*mtproto.Bool, error) {
	sqlx.TxWrapper(c.ctx, c.svcCtx.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		c.svcCtx.Dao.DialogsDAO.UpdateCustomMapTx(
			tx,
			map[string]interface{}{
				"theme_emoticon": in.ThemeEmoticon,
			},
			in.UserId,
			in.PeerType,
			in.PeerId)
		c.svcCtx.Dao.DialogsDAO.UpdateCustomMapTx(
			tx,
			map[string]interface{}{
				"theme_emoticon": in.ThemeEmoticon,
			},
			in.PeerId,
			in.PeerType,
			in.UserId)
	})

	return mtproto.BoolTrue, nil
}
