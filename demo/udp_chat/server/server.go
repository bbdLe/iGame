package main

import (
	"fmt"
	"reflect"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/demo/chat/proto"
	"go.uber.org/zap"

	_ "github.com/bbdLe/iGame/comm/peer/udp"
	_ "github.com/bbdLe/iGame/comm/processor/udp"
)

func main() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("udp.Acceptor", "server.chat", "localhost:14444", q)
	processor.BindProcessorHandler(p, "udp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *proto.ChatReq:
			var rsp proto.ChatRes
			rsp.Msg = msg.GetMsg()
			rsp.SessionId = ev.Session().ID()

			log.Logger.Debug(fmt.Sprintf("Count : %d", p.(peer.SessionManager).Count()))
			p.(peer.SessionManager).VisitSession(func(sess comm.Session) bool {
				log.Logger.Debug("=====")
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
