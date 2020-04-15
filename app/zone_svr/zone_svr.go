package zone_svr

import (
	"github.com/bbdLe/iGame/app/zone_svr/internal/handler"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func Run() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "zone_svr", "localhost:10010", q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event){
		switch ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			log.Logger.Debug("zone_conn connect")
		case *sysmsg.SessionClose:
			log.Logger.Debug("zone_conn close")
		default:
			handler.MsgDispather.OnEvent(ev)
	}
	})
	q.StartLoop()
	p.Start()

	select {

	}
}
