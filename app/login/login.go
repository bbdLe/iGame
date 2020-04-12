package login

import (
	"github.com/bbdLe/iGame/app/login/handler"
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
	p := peer.NewGenericPeer("tcp.Acceptor", "login", "localhost:12315", q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("Session Connect")
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Disconnected")
		default:
			handler.MsgDispatcher.OnEvent(ev)
		}
	})
	q.StartLoop()
	p.Start()

	select {
	}
}
