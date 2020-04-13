package handler

import (
	"fmt"
	"github.com/bbdLe/iGame/comm/peer"

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

	// 回包
	ev.Session().Send(&proto.VerifyRes{
		RetCode: 0,
		RetMsg: "",
	})
}

// 默认处理, 处理消息中转
func ZoneDefaultHanlder(ev processor.Event) {
	v, ok := ev.Session().(comm.ContextSet).GetContext("user")
	if !ok {
		log.Logger.Debug("session don't verify")
		ev.Session().Close()
	}
	user := v.(*model.User)

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
		ClientId: user.Session.ID(),
	}

	ZoneSvrConn.(peer.TCPConnector).Session().Send(msg)
	log.Logger.Debug(fmt.Sprintf("%v", msg))
}
