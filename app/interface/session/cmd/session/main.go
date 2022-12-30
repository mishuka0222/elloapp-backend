package main

import (
	session_helper "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/session"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"
)

func main() {
	commands.Run(session_helper.New())
}
