package server

import (
	"help_center/spiderbycolly/spiderService/server/controller"
)

func Run(address string, release bool) error {

	return controller.Run(address, release)
}
