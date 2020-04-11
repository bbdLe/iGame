package main

import (
	"bufio"
	"fmt"
	"github.com/bbdLe/iGame/demo/ws_chat/proto"
	"os"
	"reflect"
	"strings"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"

	_ "github.com/bbdLe/iGame/comm/peer/websocket"
	_ "github.com/bbdLe/iGame/comm/processor/websocket"
)

func ReadFromConsole(cb func(msg string)) {
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msg = strings.TrimRight(msg, "\n")
		cb(msg)
	}
}

func main() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("ws.Connector", "ws.chat_client", "localhost:12345", q)
	processor.BindProcessorHandler(p, "ws.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("session connect")
		case *sysmsg.SessionClose:
			log.Logger.Debug("session close")
		case *proto.ChatRes:
			log.Logger.Debug(fmt.Sprintf("recv msg : %v", msg.Msg))
		default:
			log.Logger.Debug(fmt.Sprintf("Recv Wrong Msg : %v", reflect.TypeOf(msg)))
		}
	})
	q.StartLoop()
	p.Start()

	ReadFromConsole(func(msg string) {
		p.(interface{ Session() comm.Session}).Session().Send(&proto.ChatReq{
			Msg: msg,
		})
	})
}
