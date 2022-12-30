package main

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/bizraw/internal/server"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"
)

func main() {
	commands.Run(server.New())
}
