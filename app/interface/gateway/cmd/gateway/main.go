package main

import (
	gateway_helper "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/gateway"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/commands"
)

func main() {
	commands.Run(new(gateway_helper.Server))
}
