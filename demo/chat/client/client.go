package main

import (
	"bufio"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/demo/chat/proto"
	"log"
	"os"

	_ "github.com/bbdLe/iGame/comm/processor/tcp"
	_ "github.com/bbdLe/iGame/comm/peer/tcp"
)

func ReadConsoule(cb func(msg string)) {
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		cb(msg)
	}
}

func main() {
	q := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "tcp.chatclient", "localhost:14444", q)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event){
		switch msg := ev.Message().(type) {
		case *proto.ChatRes:
			log.Printf("msg content : %s, msg session : %d", msg.GetMsg(), msg.GetSessionId())
		}
	})

	p.Start()
	q.StartLoop()

	ReadConsoule(func(msg string) {
		var chatReq proto.ChatReq
		chatReq.Msg = msg

		p.(interface{ Session() comm.Session}).Session().Send(&chatReq)
	})
}