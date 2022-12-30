package main

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/internal/server"
)

func main() {
	commands.Run(server.New())
}
