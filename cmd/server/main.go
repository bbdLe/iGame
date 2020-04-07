package main

import (
	"github.com/bbdLe/iGame/cmd"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/codec"
	_ "github.com/bbdLe/iGame/comm/codec/gogopb"
	"github.com/bbdLe/iGame/comm/peer"
	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	"github.com/bbdLe/iGame/comm/processor"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"log"
	"reflect"
)

func init() {
	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: 1,
		Type: reflect.TypeOf((*cmd.CalReq)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: 2,
		Type: reflect.TypeOf((*cmd.CalRes)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	queue := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Acceptor", "server.acceptor", "localhost:4444", queue)
	p.(peer.TCPSocketOption).SetMaxPacketSize(1000)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			log.Println("client connect")
			log.Println("Session Cnt:", p.(peer.SessionManager).SessionCount())
		case *sysmsg.SessionClose:
			log.Println("client disconnect")
			log.Println("Session Cnt:", p.(peer.SessionManager).SessionCount())
		case *cmd.CalReq:
			var reply cmd.CalRes
			reply.Result = msg.GetA() + msg.GetB()
			ev.Session().Send(&reply)
		}
	})
	p.Start()
	queue.StartLoop()

	select {

	}
}
