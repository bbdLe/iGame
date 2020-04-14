package logic

import (
	"fmt"

	"github.com/bbdLe/iGame/app/zone_svr/ZoneMsgDispatcher"
	"github.com/bbdLe/iGame/app/zone_svr/model"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/proto"
)

var (
	MsgDispatcher = ZoneMsgDispatcher.NewZoneMsgDispather()
)

func Send2Player(player *model.Player, msg interface{}) {
	meta := comm.MessageMetaByMsg(msg)
	if meta == nil {
		log.Logger.Error("get meta fail")
		return
	}

	data, err := meta.Codec.Encode(msg, nil)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Encode msg(%d) fail: %v", meta.MsgId, err))
		return
	}

	player.Ses.Send(&proto.TransmitRes{
		MsgId: int32(meta.MsgId),
		MsgData: data.([]byte),
		ClientId: player.SessionID,
	})
}

func init() {
	MsgDispatcher.Register(int16(proto.ProtoID_CS_CMD_LOGIN_REQ), ZoneMsgLogin)
}
