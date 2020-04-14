package handler

import (
	"fmt"
	"github.com/bbdLe/iGame/comm/peer"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/proto"
)

func ZoneMsgVerify(ev processor.Event) {
	// 回包
	ev.Session().Send(&proto.VerifyRes{
		RetCode: 0,
		RetMsg: "",
	})
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

	ZoneSvrConn.(peer.TCPConnector).Session().Send(msg)
	log.Logger.Debug(fmt.Sprintf("%v", msg))
}
