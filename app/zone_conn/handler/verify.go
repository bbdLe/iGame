package handler

import (
	"fmt"

	"github.com/bbdLe/iGame/app/zone_conn/model"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

func ZoneMsgVerify(ev processor.Event) {
	// todo check token

	msg := ev.Message().(*proto.VerifyReq)
	user := model.NewUser(ev.Session(), msg.GetServer())
	ev.Session().(comm.ContextSet).SetContext("user", user)

	// 回报
	ev.Session().Send(&proto.VerifyRes{
		RetCode: 0,
		RetMsg: "",
	})
}

func ZoneMsgChat(ev processor.Event) {
	user, ok := ev.Session().(comm.ContextSet).GetContext("user")
	if !ok {
		log.Logger.Debug("session don't verify")
		ev.Session().Close()
	} else {
		log.Logger.Debug(fmt.Sprintf("Recv Chat : %s, Server : %s", ev.Message().(*proto.ChatReq).GetContent(), user.(*model.User).ServerId))
	}

	ev.Session().Send(&proto.ChatRes{
		RetCode: 0,
		RetMsg: "",
	})
}
