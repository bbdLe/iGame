package handler

import (
	"fmt"

	"github.com/bbdLe/iGame/app/zone_svr/internal/logic"
	"github.com/bbdLe/iGame/app/zone_svr/internal/model"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

var (
	MsgDispather  *processor.MessageDispatcher
)

func init() {
	MsgDispather = processor.NewMessageDispatcher()
	MsgDispather.RegisterMessage("TransmitReq", ZoneMsgTransmit)
}

func ZoneMsgTransmit(ev processor.Event) {
	msg := ev.Message().(*proto.TransmitReq)

	// 获取玩家
	p, _ := model.GetPlayer(msg.GetClientId())

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

	log.Logger.Info(fmt.Sprintf("%v", obj))
}