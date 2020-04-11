
package main

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/demo/ws_echo/proto"

	_ "github.com/bbdLe/iGame/comm/peer/websocket"
	_ "github.com/bbdLe/iGame/comm/processor/websocket"
	_"github.com/bbdLe/iGame/demo/ws_echo/proto"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("ws.Connector", "ws.chatclient", "localhost:14444", q)
	processor.BindProcessorHandler(p, "ws.ltv", func(ev processor.Event){
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionConnected:
			ev.Session().Send(&proto.EchoReq{
				Msg: "Hello, World",
			})
		case *sysmsg.SessionClose:
			wg.Done()
		case *proto.EchoRes:
			log.Logger.Debug(fmt.Sprintf("Recv Msg : %v", msg.Msg))
		default:
			log.Logger.Error(fmt.Sprintf("Recv wrong type : %v", reflect.TypeOf(msg)))
		}
	})

	p.Start()
	q.StartLoop()
	wg.Wait()
}
