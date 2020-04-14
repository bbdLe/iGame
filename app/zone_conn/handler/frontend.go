package handler

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/processor"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"github.com/bbdLe/iGame/proto"
)

var (
	frontEnd comm.Peer
)

func StartFroneEnd() {
	q := comm.NewEventQueue()
	frontEnd = peer.NewGenericPeer("tcp.Acceptor", "zone_conn", "localhost:10086", q)
	processor.BindProcessorHandler(frontEnd, "tcp.ltv", func(ev processor.Event) {
		switch ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			log.Logger.Debug("New Session Conn")
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Close")
		case *proto.VerifyReq:
			ZoneMsgVerify(ev)
		default:
			MsgDispatcher.OnEvent(ev)
		}
	})
	q.StartLoop()
	frontEnd.Start()
}
