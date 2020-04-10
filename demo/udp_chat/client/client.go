package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/demo/udp_chat/proto"

	_ "github.com/bbdLe/iGame/comm/peer/udp"
	_ "github.com/bbdLe/iGame/comm/processor/udp"
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
	p := peer.NewGenericPeer("udp.Connector", "udp.chatclient", "localhost:14444", q)
	processor.BindProcessorHandler(p, "udp.ltv", func(ev processor.Event){
		switch msg := ev.Message().(type) {
		case *proto.ChatRes:
			fmt.Println("recv:", msg.Msg)
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
