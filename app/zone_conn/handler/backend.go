package handler

import (
	"sync"
	"time"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/proto"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

var (
	ZoneSvrConn comm.Peer
)

func ConnectBackend() {
	var wg sync.WaitGroup
	wg.Add(1)

	q := comm.NewEventQueue()
	ZoneSvrConn = peer.NewGenericPeer("tcp.Connector", "zone_svr", "localhost:10010", q)
	ZoneSvrConn.(peer.TCPConnector).SetReconnectDuration(time.Second * 3)
	processor.BindProcessorHandler(ZoneSvrConn, "tcp.ltv", func(ev processor.Event) {
		switch ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("connect")
			wg.Done()
		case *sysmsg.SessionClose:
			log.Logger.Debug("disconnect")
			wg.Add(1)
		case *proto.TransmitRes:
			log.Logger.Debug("=============")
		default:
		}
	})
	q.StartLoop()
	ZoneSvrConn.Start()

	wg.Wait()
}
