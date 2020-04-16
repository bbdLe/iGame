package frontend

import (
	"fmt"
	"github.com/bbdLe/iGame/app/zone_conn/logic"
	"time"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

const (
	SESSION_CLIENT = iota
	SESSION_SERVER
)

func ZoneMsgVerify(ev processor.Event) {
	// 回包
	ev.Session().(comm.ContextSet).SetContext("auth", true)
	ev.Session().Send(&proto.VerifyRes{
		RetCode: 0,
		RetMsg: "",
	})
}

// conn连接上来
func ZoneMsgNewConn(ev processor.Event) {
	log.Logger.Debug("new session connect")
	now := time.Now().Unix()

	ev.Session().(comm.ContextSet).SetContext(logic.SessionTypeKey, SESSION_CLIENT)
	ev.Session().(comm.ContextSet).SetContext(logic.HeartBeatKey, now)
	ev.Session().(comm.ContextSet).SetContext(logic.AuthTimeKey, now)
	ev.Session().(comm.ContextSet).SetContext(logic.AuthKey, false)

	// 新增连接
	logic.FrontEndMgr.AddSession(ev.Session())
}

func ZoneMsgConnClose(ev processor.Event) {
	v, ok := ev.Session().(comm.ContextSet).GetContext(logic.SessionTypeKey)
	if !ok {
		log.Logger.Error(fmt.Sprintf("session[%d] without session type", v))
		return
	}

	if v.(int) == SESSION_SERVER {
		return
	}

	// 通知后端
	logic.BackEndMgr.Send(
		&proto.ConnDisconnectReq{
			ClientId: ev.Session().ID(),
	})

	// 移除连接
	logic.FrontEndMgr.DelSession(ev.Session())
	log.Logger.Debug("conn close")
}

// 默认处理, 处理消息中转
func ZoneDefaultHanlder(ev processor.Event) {
	meta := comm.MessageMetaByMsg(ev.Message())
	if meta == nil {
		log.Logger.Error(fmt.Sprintf("get meta fail : %v", ev.Message()))
		return
	}

	data, err := meta.Codec.Encode(ev.Message(), nil)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Codec faild : %v", err))
		return
	}

	msg := &proto.TransmitReq{
		MsgId:  int32(meta.MsgId),
		MsgData: data.([]byte),
		ClientId: ev.Session().ID(),
	}

	logic.BackEndMgr.Send(msg)
}