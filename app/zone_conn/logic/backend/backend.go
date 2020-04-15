package backend

import (
	"sync"
	"time"

	"github.com/bbdLe/iGame/app/zone_conn/logic"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

var (
	msgDispatcher *processor.MessageDispatcher
)

func init() {
	msgDispatcher = processor.NewMessageDispatcher()
	msgDispatcher.RegisterMessage("TransmitRes", ZoneMsgTransmit)
	msgDispatcher.RegisterMessage("ConnDisconnectRes", ZoneMsgConnDisconnectRes)
}

func ConnectBackend() {
	var wg sync.WaitGroup
	wg.Add(1)

	q := comm.NewEventQueue()
	logic.BackEndConnector = peer.NewGenericPeer("tcp.Connector", "zone_svr", "localhost:10010", q)
	logic.BackEndConnector.(peer.TCPConnector).SetReconnectDuration(time.Second * 3)

	processor.BindProcessorHandler(logic.BackEndConnector, "tcp.ltv", func(ev processor.Event) {
		switch ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("connect")
			wg.Done()
		case *sysmsg.SessionClose:
			log.Logger.Debug("disconnect")
			wg.Add(1)
		default:
			msgDispatcher.OnEvent(ev)
		}
	})
	q.StartLoop()
	logic.BackEndConnector.Start()

	wg.Wait()
}
