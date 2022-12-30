/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package main

import (
	dfs_helper "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"
)

func main() {
	commands.Run(dfs_helper.New())
}
