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

package plugin

import (
	"context"
	"github.com/teamgram/proto/mtproto"
)

type MessagesPlugin interface {
	GetWebpagePreview(ctx context.Context, url string) (*mtproto.WebPage, error)
	GetMessageMedia(ctx context.Context, ownerId int64, media *mtproto.InputMedia) (*mtproto.MessageMedia, error)
	// RebuildMessageEntities(ctx context.Context, fromId int64, peer *mtproto.PeerUtil, noWebpage bool, message *mtproto.Message, hasBot bool) (*mtproto.Message, error)
}
