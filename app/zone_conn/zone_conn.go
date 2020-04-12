package zone_conn

import (
	"github.com/bbdLe/iGame/app/zone_conn/handler"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/proto"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func Run() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "zone_conn", "localhost:10086", q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
	switch ev.Message().(type) {
	case *proto.VerifyReq:
		log.Logger.Debug("client send verify")
		ev.Session().Send(&proto.VerifyRes{
			RetCode: 0,
			RetMsg: "verify succ",
		})
	case *sysmsg.SessionConnected:
		log.Logger.Debug("New Session Conn")
	case *sysmsg.SessionClose:
		log.Logger.Debug("Session Close")
	default:
		handler.MsgDispatcher.OnEvent(ev)
	}
	})
	q.StartLoop()
	p.Start()

	select {

	}
}