package logic

import (
	"github.com/bbdLe/iGame/app/zone_conn/logic"
	"sync"
	"time"

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
	msgDispatcher.RegisterMessage("KickConnReq", ZoneMsgKickConnReq)

	logic.BackEndMgr = NewBackEndManager()
}

type BackendManagerImpl struct {
	queue comm.EventQueue
	connector comm.Peer
}

func (self *BackendManagerImpl) Start() {
	var wg sync.WaitGroup
	wg.Add(1)

	self.queue = comm.NewEventQueue()
	self.connector = peer.NewGenericPeer("tcp.Connector", "zone_svr", "localhost:10010", self.queue)
	self.connector.(peer.TCPConnector).SetReconnectDuration(time.Second * 3)

	processor.BindProcessorHandler(self.connector, "tcp.ltv", func(ev processor.Event) {
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
	self.queue.StartLoop()
	self.connector.Start()

	wg.Wait()
}

func (self *BackendManagerImpl) Send(msg interface{}) {
	self.connector.(peer.TCPConnector).Session().Send(msg)
}

func NewBackEndManager() *BackendManagerImpl {
	return &BackendManagerImpl{}
}