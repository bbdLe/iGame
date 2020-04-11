package main

import (
	"fmt"
	"reflect"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/demo/ws_chat/proto"

	_ "github.com/bbdLe/iGame/comm/peer/websocket"
	_ "github.com/bbdLe/iGame/comm/processor/websocket"
)

func main() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("ws.Acceptor", "ws.chat_server", "localhost:12345", q)
	processor.BindProcessorHandler(p, "ws.ltv", func(ev processor.Event){
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			log.Logger.Debug("session accept")
		case *proto.ChatReq:
			ev.Session().Peer().(comm.SessionAccessor).VisitSession(func(ses comm.Session) bool {
				ses.Send(&proto.ChatRes{
					Msg: msg.Msg,
				})
				return true
			})
		default:
			log.Logger.Debug(fmt.Sprintf("Recv Unknow Type : %v", reflect.TypeOf(msg) ))
		}
	})
	q.StartLoop()
	p.Start()

	select {

	}
}
