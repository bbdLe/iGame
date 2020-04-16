package logic

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal/ZoneMsgDispatcher"
	"github.com/bbdLe/iGame/proto"
)

var (
	MsgDispatcher *ZoneMsgDispatcher.ZoneMsgDispatcher
)

func init() {
	MsgDispatcher = ZoneMsgDispatcher.NewZoneMsgDispather()
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_LOGIN_REQ), ZoneMsgLogin)
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_HEART_BETA_REQ), ZoneMsgHeartBeat)
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_CREATE_ROOM_REQ), ZoneMsgCreateRoom)
}
