package zone_svr

import (
	"fmt"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/proto"

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


			log.Logger.Debug(fmt.Sprintf("%v", ev.Message()))
		default:
			log.Logger.Debug("test")
	}
	})
	q.StartLoop()
	p.Start()

	select {

	}
}
