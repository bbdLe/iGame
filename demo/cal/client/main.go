package main

import (
	"fmt"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/demo/cal/proto"
	"log"
	"time"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

//func init() {
//	comm.RegMessageMeta(&comm.MessageMeta{
//		MsgId: int(util.StringHash("cmd.CalReq")),
//		Type: reflect.TypeOf((*proto.CalReq)(nil)).Elem(),
//		Codec: codec.MustGetCodec("gogopb"),
//	})
//
//	comm.RegMessageMeta(&comm.MessageMeta{
//		MsgId: int(util.StringHash("cmd.CalRes")),
//		Type: reflect.TypeOf((*proto.CalRes)(nil)).Elem(),
//		Codec: codec.MustGetCodec("gogopb"),
//	})
//}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	queue := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "clientAsynCallback", "localhost:4444", queue)
	p.(peer.TCPSocketOption).SetMaxPacketSize(1000)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionConnected:
			fmt.Println("connected", msg)
			ev.Session().Send(&proto.CalReq{A : 1, B: 2})
		case *proto.CalRes:
			fmt.Printf("CalRes: %d", msg.GetResult())
		}
	})
	p.Start()
	queue.StartLoop()

	time.Sleep(time.Second * 10)
}