package logic

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

func ZoneMsgCreateRoom(player internal.Player, ev processor.Event) {
	if player == nil {
		log.Logger.Error("player not exist")
		return
	}

	r := internal.RoomMgr.NewRoom()
	r.SetCanFree(true)
	r.AddPlayer(player)

	reply := &proto.CreateRoomRes{}
	reply.RoomId = r.ID()
	player.Send(reply)
}

func ZoneMsgEnterRoom(player internal.Player, ev processor.Event) {
	if player == nil {
		log.Logger.Error("player not exist")
		return
	}

	if player.Room() != nil {

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
	reply.RoomId = msg.GetRoomId()

	room.AddPlayer(player)
	player.Send(reply)
}

func ZoneMsgMove(player internal.Player, ev processor.Event) {
	if player == nil {
		log.Logger.Error("player not exist")
		return
	}

	msg := ev.Message().(*proto.MovePosReq)
	player.Room().OnPlayerMove(player, msg.GetPos().GetX(), msg.GetPos().GetY())

	reply := &proto.MovePosRes{}
	player.Send(reply)
}