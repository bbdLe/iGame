package logic

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/proto"
	"time"

	"github.com/bbdLe/iGame/app/zone_svr/internal/model"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
)

func ZoneMsgLogin(player internal.Player, ev processor.Event) {
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
	internal.GameMgr.SetPlayer(clientID, player)

	// 回调
	player.OnLogin()
	internal.GameMgr.OnPlayerLogin(player)

	internal.Send2Player(player, &proto.LoginRes{
		RetCode: 0,
		RetMsg: "",
	})
}

func ZoneMsgHeartBeat(player internal.Player, ev processor.Event) {
	if player == nil {
		// todo 踢掉session
		log.Logger.Error("player not exist")
		return
	}

	player.SetHeartBeat(time.Now())
	// 回包
	internal.Send2Player(player, &proto.HeartBeatRes{})
}

