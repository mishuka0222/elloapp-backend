package main

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/commands"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/internal/server"
)

func main() {
	commands.Run(server.New())
}
