package main

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/qrcode/internal/server"
)

func main() {
	commands.Run(server.New())
}
