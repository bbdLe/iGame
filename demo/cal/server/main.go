package main

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/demo/cal/proto"
	"log"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	queue := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "server.acceptor", "localhost:4444", queue)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			log.Println("client connect")
		case *sysmsg.SessionClose:
			log.Println("client disconnect")
		case *proto.CalReq:
			var reply proto.CalRes
			reply.Result = msg.GetA() + msg.GetB()
			ev.Session().Send(&reply)
		}
	})
	p.Start()
	queue.StartLoop()

	select {

	}
}
