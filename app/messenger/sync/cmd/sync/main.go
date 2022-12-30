package main

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/internal/server"
)

func main() {
	commands.Run(server.New())
}
