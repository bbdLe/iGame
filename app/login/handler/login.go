package handler

import (
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/util"
	"github.com/bbdLe/iGame/proto"
)

func ZoneMsgLogin(ev processor.Event) {
	loginReq := ev.Message().(*proto.LoginReq)
	log.Logger.Debug(loginReq.GetPlatform())
	log.Logger.Debug(loginReq.GetUid())
	log.Logger.Debug(loginReq.GetVersion())

	token := util.NewToken()
	ev.Session().Send(&proto.LoginRes{
		Result: 0,
		Token: token,
		ServerAddr: "localhost:10086",
	})
}