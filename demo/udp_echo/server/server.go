package main

import (
	"fmt"
	"github.com/bbdLe/iGame/demo/udp_echo/proto"
	"reflect"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"

	_ "github.com/bbdLe/iGame/comm/peer/udp"
	_ "github.com/bbdLe/iGame/comm/processor/udp"
	_ "github.com/bbdLe/iGame/demo/udp_echo/proto"
)

func main() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("udp.Acceptor", "server.echo", "localhost:14444", q)
	processor.BindProcessorHandler(p, "udp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *proto.EchoReq:
			ev.Session().Send(&proto.EchoRes{
				Msg: msg.Msg,
			})
		default:
			log.Logger.Error(fmt.Sprintf("unknow msg type : %v", reflect.TypeOf(msg)))
		}
	})

	p.Start()
	q.StartLoop()

	select {

	}
}
