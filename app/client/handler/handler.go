package handler

import (
	"fmt"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
	"os"
	"strconv"
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
	MsgDispatcher.RegisterMessage("BroadcastMsgRes", ZoneBroadcastMsgRes)
	MsgDispatcher.RegisterMessage("EnterRoomRes", ZoneMsgEnterRoom)
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
	if len(os.Args) > 1 {
		roomid, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Logger.Debug("parse int fail")
			return
		}

		ev.Session().Send(&proto.EnterRoomReq{
			RoomId: roomid,
		})
	} else {
		ev.Session().Send(&proto.CreateRoomReq{})
	}
}

func ZoneCreateRoomRes(ev processor.Event) {
	log.Logger.Debug(fmt.Sprintf("你已进入房间[%d]", ev.Message().(*proto.CreateRoomRes).GetRoomId()))
}

func ZoneBroadcastMsgRes(ev processor.Event) {
	msg := ev.Message().(*proto.BroadcastMsgRes)

	content := ""
	if msg.GetType() == proto.MSG_TYPE_SYSTEM {
		content += "[系统]"
	} else {
		content += "[世界]"
	}

	content += msg.GetMsg()
	log.Logger.Info(content)
}

func ZoneMsgEnterRoom(ev processor.Event) {
	msg := ev.Message().(*proto.EnterRoomRes)

	if msg.GetRetCode() != 0 {
		log.Logger.Debug(fmt.Sprintf("进入房间失败: %v", msg.GetRetMsg()))
	} else {
		log.Logger.Debug(fmt.Sprintf("你已经进入%d号房间", msg.GetRoomId()))
	}
}

func ZoneHeartBeatRes(ev processor.Event) {
	// pass
}