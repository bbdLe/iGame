package zone_conn

import (
	"github.com/bbdLe/iGame/proto"

	"github.com/bbdLe/iGame/app/zone_conn/handler"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	"github.com/bbdLe/iGame/comm/processor"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
	"github.com/bbdLe/iGame/comm/sysmsg"
)

func Run() {
	handler.ConnectBackend()

	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "zone_conn", "localhost:10086", q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
	switch ev.Message().(type) {
	case *sysmsg.SessionAccepted:
		log.Logger.Debug("New Session Conn")
	case *sysmsg.SessionClose:
		log.Logger.Debug("Session Close")
	case *proto.VerifyReq:
		handler.ZoneMsgVerify(ev)
	default:
		handler.MsgDispatcher.OnEvent(ev)
	}
	})
	q.StartLoop()
	p.Start()

	select {

	}
}