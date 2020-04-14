package handler

import (
	"fmt"
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
		switch msg := ev.Message().(type) {
		case *sysmsg.SessionConnected:
			log.Logger.Debug("connect")
			wg.Done()
		case *sysmsg.SessionClose:
			log.Logger.Debug("disconnect")
			wg.Add(1)
		// 转发到前端
		case *proto.TransmitRes:
			// 转发
			ses := frontEnd.(comm.SessionAccessor).GetSession(msg.ClientId)
			if ses == nil {
				log.Logger.Error(fmt.Sprintf("Can't get user %d", msg.ClientId))
				return
			}

			meta := comm.MessageMetaByID(int(msg.GetMsgId()))
			if meta == nil {
				log.Logger.Error(fmt.Sprintf("Can't get msg(%d) meta", msg.GetMsgId()))
				return
			}

			obj := meta.NewType()
			err := meta.Codec.Decode(msg.GetMsgData(), obj)
			if err != nil {
				log.Logger.Error(fmt.Sprintf("Decode msg(%d) fail : %v", msg.GetMsgId(), err))
				return
			}

			ses.Send(obj)
		default:
		}
	})
	q.StartLoop()
	ZoneSvrConn.Start()

	wg.Wait()
}
