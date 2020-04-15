package handler

import (
	"fmt"
	"time"

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
			ev.Session().(comm.ContextSet).SetContext("auth", false)
			ev.Session().(comm.ContextSet).SetContext("lastHeartBeatTime", time.Now().Unix())
		case *sysmsg.SessionClose:
			log.Logger.Debug("Session Close")
		case *proto.VerifyReq:
			ev.Session().(comm.ContextSet).SetContext("auth", true)
			ev.Session().(comm.ContextSet).SetContext("lastHeartBeatTime", time.Now().Unix())
			ZoneMsgVerify(ev)
		default:
			ev.Session().(comm.ContextSet).SetContext("lastHeartBeatTime", time.Now().Unix())
			v, ok := ev.Session().(comm.ContextSet).GetContext("auth")
			log.Logger.Debug(fmt.Sprintf("%v, %v", v, ok))
			if !ok || v.(bool) != true {
				log.Logger.Error("session not auth, kick it out")
				ev.Session().Close()
				return
			}
			MsgDispatcher.OnEvent(ev)
		}
	})
	q.StartLoop()
	frontEnd.Start()

	// kick time out
	go frontendTick(q, frontEnd)
}

func frontendTick(q comm.EventQueue, peer comm.Peer) {
	for {
		q.Post(func() {
			frontEnd.(comm.SessionAccessor).VisitSession(func(ses comm.Session) bool {
				v, ok := ses.(comm.ContextSet).GetContext("lastHeartBeatTime")
				// 没有设置heartbeat
				if !ok {
					log.Logger.Debug("session without heartbeat")
					q.Post(func() {
						ses.Close()
					})
				}
				lastHeartBeatTime := v.(int64)
				lastHeartBeat := time.Unix(lastHeartBeatTime, 0)
				if lastHeartBeat.Add(time.Second * 3).Before(time.Now()) {
					q.Post(func() {
						log.Logger.Debug("session heartbeat timeout")
						ses.Close()
					})
				}

				return true
			})
		})

		time.Sleep(time.Second)
	}
}