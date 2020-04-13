package logic

import (
	"github.com/bbdLe/iGame/app/zone_svr/ZoneMsgDispatcher"
	"github.com/bbdLe/iGame/proto"
)

var (
	MsgDispatcher = ZoneMsgDispatcher.NewZoneMsgDispather()
)

func init() {
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_LOGIN_REQ), ZoneMsgLogin)
}
