package main

import (
	dfs_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/commands"
)

func main() {
	commands.Run(dfs_helper.New())
}
