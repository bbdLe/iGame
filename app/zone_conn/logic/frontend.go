package logic

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

func StartFrontEnd() {
	q := comm.NewEventQueue()
	frontEnd = peer.NewGenericPeer("tcp.Acceptor", "zone_conn", "localhost:10086", q)
	processor.BindProcessorHandler(frontEnd, "tcp.ltv", func(ev processor.Event) {
		switch ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			ZoneMsgNewConn(ev)
		case *sysmsg.SessionClose:
			ZoneMsgConnClose(ev)
		case *proto.VerifyReq:
			ev.Session().(comm.ContextSet).SetContext(heartBeatKey, time.Now().Unix())
			ZoneMsgVerify(ev)
		default:
			ev.Session().(comm.ContextSet).SetContext(heartBeatKey, time.Now().Unix())

			// 没验证的连接要踢掉
			v, ok := ev.Session().(comm.ContextSet).GetContext(authKey)
			if !ok || v.(bool) != true {
				log.Logger.Error(fmt.Sprintf("session[%d] not auth, kick it out", ev.Session().ID()))
				ev.Session().Close()
				return
			}

			// 分发
			FrontMsgDispatcher.OnEvent(ev)
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
			KickIllegalConn(q, peer)
		})

		time.Sleep(time.Second)
	}
}

func KickIllegalConn(q comm.EventQueue, peer comm.Peer) {
	frontEnd.(comm.SessionAccessor).VisitSession(func(ses comm.Session) bool {
		now := time.Now()

		// 踢掉不活跃连接
		v1, ok := ses.(comm.ContextSet).GetContext(heartBeatKey)
		if !ok {
			log.Logger.Error(fmt.Sprintf("session[%d] without heartbeat", ses.ID()))
			q.Post(func() {
				ses.Close()
			})
			return true
		}
		lastHeartBeatTime := v1.(int64)
		lastHeartBeat := time.Unix(lastHeartBeatTime, 0)
		if lastHeartBeat.Add(time.Second * 3).Before(now) {
			q.Post(func() {
				log.Logger.Info(fmt.Sprintf("session[%d] heartbeat timeout", ses.ID()))
				ses.Close()
			})
		}

		// 踢掉验证超时
		v2, ok := ses.(comm.ContextSet).GetContext(authKey)
		if !ok {
			log.Logger.Error(fmt.Sprintf("session[%d] without autukey, kick it out", ses.ID()))
			q.Post(func() {
				ses.Close()
			})
			return true
		}
		// 验证过的跳过
		auth := v2.(bool)
		if auth {
			return true
		}

		v3, ok := ses.(comm.ContextSet).GetContext(authTimeKey)
		if !ok {
			log.Logger.Error(fmt.Sprintf("session[%d] without authtime key, kick it out", ses.ID()))
			q.Post(func() {
				ses.Close()
			})
			return true
		}

		authTime := v3.(int64)
		if time.Unix(authTime, 0).Add(time.Second * 3).Before(now) {
			log.Logger.Info(fmt.Sprintf("session[%d] auth out of limit time, kick it out", ses.ID()))
			q.Post(func() {
				ses.Close()
			})
			return true
		}

		return true
	})
}