package logic

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/proto"

	"github.com/bbdLe/iGame/app/zone_svr/model"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
)

func ZoneMsgLogin(player *model.Player, ev processor.Event) {
	v, ok := ev.Session().(comm.ContextSet).GetContext("clientID")
	if !ok {
		log.Logger.Error("Get ClientID fail")
		return
	}
	clientID := v.(int64)

	// 登陆玩家应该是不存在的
	if player != nil {
		log.Logger.Error("First Login, but user is exist, kick it out")
		ev.Session().Send(&proto.KickConnReq{
			ClientId: clientID,
		})
		return
	}

	// 新建玩家
	player = model.NewPlayer(clientID, ev.Session())
	model.SetPlayer(clientID, player)

	Send2Player(player, &proto.LoginRes{
		RetCode: 0,
		RetMsg: "",
	})
}

