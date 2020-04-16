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
	MsgDispatcher.RegisterMessage("HeartBeatRes", ZoneHeartBeatRes)
	MsgDispatcher.RegisterMessage("CreateRoomRes", ZoneCreateRoomRes)
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

	// 进入房间
	ev.Session().Send(&proto.CreateRoomReq{})
}

func ZoneCreateRoomRes(ev processor.Event) {
	log.Logger.Debug(fmt.Sprint("player enter room[%d]", ev.Message().(*proto.CreateRoomRes).GetRoomId()))
}

func ZoneHeartBeatRes(ev processor.Event) {

}