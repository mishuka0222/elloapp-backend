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

// HelpGetAppConfig
// help.getAppConfig#98914110 = JSONValue;
func (c *ConfigurationCore) HelpGetAppConfig(in *mtproto.TLHelpGetAppConfig) (*mtproto.JSONValue, error) {
	// TODO: not impl
	c.Logger.Errorf("help.getAppConfig blocked, License key from https://teamgram.net required to unlock enterprise features.")

	return mtproto.MakeTLJsonObject(&mtproto.JSONValue{
		Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{},
	}).To_JSONValue(), nil
}
