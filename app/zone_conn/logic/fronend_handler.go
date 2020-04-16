package logic

import (
	"fmt"
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

	ev.Session().(comm.ContextSet).SetContext(SessionTypeKey, SESSION_CLIENT)
	ev.Session().(comm.ContextSet).SetContext(HeartBeatKey, now)
	ev.Session().(comm.ContextSet).SetContext(AuthTimeKey, now)
	ev.Session().(comm.ContextSet).SetContext(AuthKey, false)

	// 新增连接
	FrontEndMgr.AddSession(ev.Session())
}

func ZoneMsgConnClose(ev processor.Event) {
	v, ok := ev.Session().(comm.ContextSet).GetContext(SessionTypeKey)
	if !ok {
		log.Logger.Error(fmt.Sprintf("session[%d] without session type", v))
		return
	}

	if v.(int) == SESSION_SERVER {
		return
	}

	// 通知后端
	BackEndMgr.Send(
		&proto.ConnDisconnectReq{
			ClientId: ev.Session().ID(),
	})

	// 移除连接
	FrontEndMgr.DelSession(ev.Session())
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

	BackEndMgr.Send(msg)
	//log.Logger.Debug(fmt.Sprintf("%v", msg))
}