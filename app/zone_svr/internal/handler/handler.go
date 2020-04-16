package handler

import (
	"fmt"
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/app/zone_svr/internal/logic"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

var (
	MsgDispather  *processor.MessageDispatcher
	BlockMsgMap map[proto.ProtoID]struct{}
)

func init() {
	MsgDispather = processor.NewMessageDispatcher()
	MsgDispather.RegisterMessage("TransmitReq", ZoneMsgTransmit)
	MsgDispather.RegisterMessage("ConnDisconnectReq", ZoneMsgConnDisconnect)
	MsgDispather.RegisterMessage("KickConnRes", ZoneMsgKickConnRes)


	BlockMsgMap = make(map[proto.ProtoID]struct{})
	BlockMsgMap[proto.ProtoID_CS_CMD_HEART_BETA_REQ] = struct{}{}
}

func ZoneMsgTransmit(ev processor.Event) {
	msg := ev.Message().(*proto.TransmitReq)

	// 获取玩家
	p, _ := internal.GameMgr.GetPlayer(msg.ClientId)

	// 获取meta
	msgID := msg.GetMsgId()
	meta := comm.MessageMetaByID(int(msgID))
	if meta == nil {
		log.Logger.Error(fmt.Sprintf("msgid[%d] meta not found", msgID))
		return
	}

	// byte -> obj
	obj := meta.NewType()
	err := meta.Codec.Decode(msg.GetMsgData(), obj)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("msgid[%d] codec fail: %v", msgID, err))
		return
	}

	// 设置上下文
	ev.Session().(comm.ContextSet).SetContext("clientID", msg.GetClientId())

	// 分发
	logic.MsgDispatcher.OnEvent(p, &event.RecvMsgEvent{
		Ses: ev.Session(),
		Msg: obj,
	})

	if _, ok := BlockMsgMap[proto.ProtoID(meta.MsgId)]; !ok {
		log.Logger.Info(fmt.Sprintf("%v", obj))
	}
}

func ZoneMsgConnDisconnect(ev processor.Event) {
	msg := ev.Message().(*proto.ConnDisconnectReq)

	clientID := msg.GetClientId()
	p, ok := internal.GameMgr.GetPlayer(clientID)
	 // 不在的话, 直接回报
	if !ok {
		log.Logger.Info(fmt.Sprintf("session[%d] disconnect: player not exist", clientID))
			ev.Session().Send(&proto.ConnDisconnectRes{
			ClientId: clientID,
			RetCode: 0,
		})
		return
	}

	p.OnLogout()
	internal.GameMgr.DelPlayer(clientID)

	ev.Session().Send(&proto.ConnDisconnectRes{
		ClientId: clientID,
		RetCode: 0,
	})

	log.Logger.Info(fmt.Sprintf("player[%d] disconnect", clientID))
}

func ZoneMsgKickConnRes(ev processor.Event) {
	return
}