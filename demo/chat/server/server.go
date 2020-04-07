package main

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/demo/chat/proto"
	"go.uber.org/zap"
	"reflect"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func main() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "server.chat", "localhost:14444", q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			log.Logger.Debug("Session Accept")
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Close")
		case *proto.ChatReq:
			var rsp proto.ChatRes
			rsp.Msg = msg.GetMsg()
			rsp.SessionId = ev.Session().ID()

			p.(comm.SessionAccessor).VisitSession(func(sess comm.Session) bool {
				sess.Send(&rsp)
				return true
			})
		default:
			log.Logger.Debug("Session Close")
			log.Logger.Error("unknow msg type", zap.String("msg type", reflect.TypeOf(msg).Elem().String()))
		}
	})

	p.Start()
	q.StartLoop()
	select {

	}
}