package ethereum

import (
	"help_center/config"
	"help_center/spiderbycolly/spiderService/ethereum/schedule"
)

func Run(conf *config.ETHConfig) error {

	return schedule.Run(conf)
}
