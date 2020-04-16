package logic

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
)

func ZoneMsgCreateRoom(player internal.CommPlayer, ev processor.Event) {
	if player == nil {
		log.Logger.Error("player not exist")
		return
	}

	r := internal.RoomMgr.NewRoom()
	r.AddPlayer(player)
}
