package logic

import (
	"fmt"

	"github.com/bbdLe/iGame/app/zone_svr/player"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
)

func ZoneMsgLogin(player *player.Player, ev processor.Event) {
	log.Logger.Debug(fmt.Sprintf("%v", ev))
}

