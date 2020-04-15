package client

import (
	"sync"

	"github.com/bbdLe/iGame/app/client/handler"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/proto"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func connectConn(addr string, token string) error {
	var wg sync.WaitGroup
	wg.Add(1)

	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "client", addr, q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("session connect")
			ev.Session().Send(&proto.VerifyReq{
				Token: token,
				Server: "1",
			})
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Close")
			wg.Done()
		default:
			handler.MsgDispatcher.OnEvent(ev)
		}
	})
	q.StartLoop()
	p.Start()

	wg.Wait()

	return nil
}

func Run() {
	connectConn("localhost:10086", "token")
}