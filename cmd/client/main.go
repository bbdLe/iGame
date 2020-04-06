package main

import (
	"fmt"
	"github.com/bbdLe/iGame/cmd"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/codec"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"log"
	"reflect"
	"time"

	_ "github.com/bbdLe/iGame/comm/codec/gogopb"
	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

func init() {
	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: 1,
		Type: reflect.TypeOf((*cmd.SearchRequest)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
		})
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	queue := comm.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "clientAsynCallback", "localhost:4444", queue)
	p.(peer.TCPSocketOption).SetMaxPacketSize(1000)
	processor.BindProcessorHandler(p, "tcp.ltv", func(ev processor.Event) {
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionConnected:
			fmt.Println("connected", msg)
			ev.Session().Send(&cmd.SearchRequest{PageNumber: 1})
		}
	})
	p.Start()
	queue.StartLoop()

	time.Sleep(time.Second * 10)
}