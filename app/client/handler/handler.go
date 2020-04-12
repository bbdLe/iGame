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
}

func ZoneVerifyRes(ev processor.Event) {
	msg := ev.Message().(*proto.VerifyRes)
	if msg.GetRetCode() == 0 {
		log.Logger.Debug("Verify succ")
	} else {
		log.Logger.Debug("Verify fail")
	}
	log.Logger.Debug(fmt.Sprintf("Recv VerifyRes : %v", msg.GetRetCode()))

	ev.Session().Send(&proto.ChatReq{
		Content: "Hello, World",
	})
}
