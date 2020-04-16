package logic

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal/zonemsgdispatcher"
	"github.com/bbdLe/iGame/proto"
)

var (
	MsgDispatcher *zonemsgdispatcher.ZoneMsgDispatcher
)

func init() {
	MsgDispatcher = zonemsgdispatcher.NewZoneMsgDispather()
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_LOGIN_REQ), ZoneMsgLogin)
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_HEART_BETA_REQ), ZoneMsgHeartBeat)
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_CREATE_ROOM_REQ), ZoneMsgCreateRoom)
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_ENTER_ROOM_REQ), ZoneMsgEnterRoom)
}
