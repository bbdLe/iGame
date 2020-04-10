
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

	_ "github.com/bbdLe/iGame/comm/peer/udp"
	_ "github.com/bbdLe/iGame/comm/processor/udp"
)

func main() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("udp.Connector", "server.echo", "localhost:14444", q)
	processor.BindProcessorHandler(p, "udp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("Session Accept")
			ev.Session().Send(&proto.EchoReq{
				Msg: "Hello,World",
			})
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Close")
		case *proto.EchoRes:
			log.Logger.Debug(fmt.Sprintf("Recv : %s", msg.Msg))
		default:
			log.Logger.Error(fmt.Sprintf("unknow msg type : %v", reflect.TypeOf(msg)))
		}
	})

	p.Start()
	q.StartLoop()

	select {

	}
}
