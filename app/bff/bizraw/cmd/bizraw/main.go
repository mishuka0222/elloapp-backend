package main

import (
	"github.com/teamgram/marmota/pkg/commands"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/bizraw/internal/server"
)

func main() {
	commands.Run(server.New())
}
