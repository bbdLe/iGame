package logic

import (
	"fmt"
	"github.com/bbdLe/iGame/app/zone_conn/logic"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

// 后端转发前端
func ZoneMsgTransmit(ev processor.Event) {
	msg := ev.Message().(*proto.TransmitRes)

	clientID := msg.GetClientId()
	ses, ok := logic.FrontEndMgr.GetSession(clientID)
	if !ok {
		log.Logger.Error(fmt.Sprintf("session[%d] not found", clientID))
		return
	}

	msgID := msg.GetMsgId()
	meta := comm.MessageMetaByID(int(msgID))
	if meta == nil {
		log.Logger.Error(fmt.Sprintf("msg[%d] meta not found", msgID))
		return
	}

	obj := meta.NewType()
	err := meta.Codec.Decode(msg.GetMsgData(), obj)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("msg[%d] decode fail : %v", msg.GetMsgId(), err))
		return
	}

	ses.Send(obj)
}

func ZoneMsgConnDisconnectRes(ev processor.Event) {
	log.Logger.Debug(fmt.Sprintf("conn disconnect : %v", ev.Message()))
}

func ZoneMsgKickConnReq(ev processor.Event) {
	msg := ev.Message().(*proto.KickConnReq)

	log.Logger.Debug(fmt.Sprintf("recv zone_svr kick conn[%d] request", msg.GetClientId()))
	logic.FrontEndMgr.Kick(msg.GetClientId())
	ev.Session().Send(&proto.KickConnRes{
		ClientId: msg.GetClientId(),
		RetCode: 0,
	})
}