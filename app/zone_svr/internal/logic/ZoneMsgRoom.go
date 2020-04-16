package logic

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

func ZoneMsgCreateRoom(player internal.CommPlayer, ev processor.Event) {
	if player == nil {
		log.Logger.Error("player not exist")
		return
	}

	r := internal.RoomMgr.NewRoom()
	r.AddPlayer(player)

	reply := &proto.CreateRoomRes{}
	reply.RoomId = r.ID()
	player.Send(reply)
}

func ZoneMsgEnterRoom(player internal.CommPlayer, ev processor.Event) {
	if player == nil {
		log.Logger.Error("player not exist")
		return
	}

	msg := ev.Message().(*proto.EnterRoomReq)
	reply := &proto.EnterRoomRes{}

	room, ok := internal.RoomMgr.GetRoom(msg.GetRoomId())
	if !ok {
		reply.RetCode = -1
		reply.RetMsg = "房间不存在"
		player.Send(reply)
		return
	}

	room.AddPlayer(player)
	player.Send(reply)
}
