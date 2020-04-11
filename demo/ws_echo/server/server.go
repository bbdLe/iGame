package main

import (
	"fmt"
	"reflect"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/demo/udp_echo/proto"

	_ "github.com/bbdLe/iGame/comm/peer/websocket"
	_ "github.com/bbdLe/iGame/comm/processor/websocket"
	_ "github.com/bbdLe/iGame/demo/udp_echo/proto"
)

func main() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("ws.Acceptor", "server.echo", "localhost:14444", q)
	processor.BindProcessorHandler(p, "ws.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			log.Logger.Debug("session accept")
		case *proto.EchoReq:
			log.Logger.Debug("send echo res")
			ev.Session().Send(&proto.EchoRes{
				Msg: msg.Msg,
			})
		case *sysmsg.SessionClose:
			log.Logger.Debug("session close")
		default:
			log.Logger.Error(fmt.Sprintf("unknow msg type : %v", reflect.TypeOf(msg)))
		}
	})

	p.Start()
	q.StartLoop()

	select {

	}
}