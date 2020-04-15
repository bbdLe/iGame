package zone_svr

import (
	"fmt"
	"github.com/bbdLe/iGame/app/zone_svr/logic"
	"github.com/bbdLe/iGame/app/zone_svr/model"
	"github.com/bbdLe/iGame/comm/event"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/proto"

	_ "github.com/bbdLe/iGame/app/zone_svr/logic"
	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
	_ "github.com/bbdLe/iGame/proto"
)

func Run() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "zone_svr", "localhost:10010", q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event){
		switch msg :=  ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			log.Logger.Debug("zone_conn connect")
		case *sysmsg.SessionClose:
			log.Logger.Debug("zone_conn close")
		case *proto.TransmitReq:
			// 获取玩家
			p, _ := model.GetPlayer(msg.GetClientId())

			// 获取meta
			meta := comm.MessageMetaByID(int(msg.GetMsgId()))
			if meta == nil {
				log.Logger.Debug(fmt.Sprintf("msgid(%d) meta not found", msg.GetMsgId()))
				break
			}

			// byte -> obj
			obj := meta.NewType()
			err := meta.Codec.Decode(msg.GetMsgData(), obj)
			if err != nil {
				return
			}

			// 设置上下文
			ev.Session().(comm.ContextSet).SetContext("clientID", msg.GetClientId())

			logic.MsgDispatcher.OnEvent(p, &event.RecvMsgEvent{
				Ses: ev.Session(),
				Msg: obj,
			})
			log.Logger.Info(fmt.Sprintf("%v", msg))
	}
	})
	q.StartLoop()
	p.Start()

	select {

	}
}
