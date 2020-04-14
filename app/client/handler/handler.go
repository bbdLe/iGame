package handler

import (
	"fmt"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

var (
	MsgDispatcher *processor.MessageDispatcher
)

func init() {
	MsgDispatcher = processor.NewMessageDispatcher()
	MsgDispatcher.RegisterMessage("VerifyRes", ZoneVerifyRes)
	MsgDispatcher.RegisterMessage("LoginRes", ZoneLoginRes)
}

func ZoneVerifyRes(ev processor.Event) {
	msg := ev.Message().(*proto.VerifyRes)
	if msg.GetRetCode() == 0 {
		log.Logger.Debug("Verify succ")
	} else {
		log.Logger.Debug("Verify fail")
	}
	log.Logger.Debug(fmt.Sprintf("Recv VerifyRes : %v", msg.GetRetCode()))

	ev.Session().Send(&proto.LoginReq{
		Version: "12345",
	})
}

func ZoneLoginRes(ev processor.Event) {
	msg := ev.Message().(*proto.LoginRes)
	log.Logger.Debug(fmt.Sprintf("%v", msg))
}
