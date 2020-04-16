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

	_ "github.com/bbdLe/iGame/comm/peer/tcp"
	_ "github.com/bbdLe/iGame/comm/processor/tcp"
)

var (
	frontEndMsgDispatcher *processor.MessageDispatcher
)

type FrontEndManager struct {
	sessionMap map[int64]comm.Session

	queue comm.EventQueue
	acceptor comm.Peer
}

func (self *FrontEndManager) Start() {
	self.queue = comm.NewEventQueue()
	self.acceptor = peer.NewGenericPeer("tcp.Acceptor", "zone_conn", "localhost:10086", self.queue)
	processor.BindProcessorHandler(self.acceptor, "tcp.ltv", func(ev processor.Event) {
		switch ev.Message().(type) {
		case *sysmsg.SessionAccepted:
			ZoneMsgNewConn(ev)

		case *sysmsg.SessionClose:
			ZoneMsgConnClose(ev)

		case *proto.VerifyReq:
			ev.Session().(comm.ContextSet).SetContext(HeartBeatKey, time.Now().Unix())
			ZoneMsgVerify(ev)

		default:
			ev.Session().(comm.ContextSet).SetContext(HeartBeatKey, time.Now().Unix())

			// 没验证的连接要踢掉
			v, ok := ev.Session().(comm.ContextSet).GetContext(AuthKey)
			if !ok || v.(bool) != true {
				log.Logger.Error(fmt.Sprintf("session[%d] not auth, kick it out", ev.Session().ID()))
				ev.Session().Close()
				return
			}

			// 分发
			frontEndMsgDispatcher.OnEvent(ev)
		}
	})
	self.queue.StartLoop()
	self.acceptor.Start()

	// kick time out
	go self.tick()
}

func (self *FrontEndManager) AddSession(ses comm.Session) {
	self.sessionMap[ses.ID()] = ses
}

func (self *FrontEndManager) DelSession(ses comm.Session) {
	delete(self.sessionMap, ses.ID())
}

func (self *FrontEndManager) GetSession(sessionID int64) (comm.Session, bool) {
	sess, ok := self.sessionMap[sessionID]
	return sess, ok
}

func (self *FrontEndManager) Visit(f func(ses comm.Session)) {
	for _, ses := range self.sessionMap {
		f(ses)
	}
}

func (self *FrontEndManager) Post(f func()) {
	self.queue.Post(f)
}

func (self *FrontEndManager) Kick(sessionID int64) {
	self.Post(func() {
		ses, ok := self.GetSession(sessionID)
		if !ok {
			return
		}

		log.Logger.Info(fmt.Sprintf("FrontEnd kick conn[%d]", sessionID))
		ses.Close()
		self.DelSession(ses)
	})
}

func (self *FrontEndManager) tick() {
	for {
		self.Post(func() {
			self.tickBase()
		})

		time.Sleep(time.Second)
	}
}

func (self *FrontEndManager) tickBase() {
	now := time.Now()

	self.Visit(func(ses comm.Session){
		// 踢掉不活跃连接
		v1, ok := ses.(comm.ContextSet).GetContext(HeartBeatKey)
		if !ok {
			log.Logger.Error(fmt.Sprintf("session[%d] without heartbeat", ses.ID()))
			self.Kick(ses.ID())
			return
		}
		lastHeartBeatTime := v1.(int64)
		lastHeartBeat := time.Unix(lastHeartBeatTime, 0)
		if lastHeartBeat.Add(time.Second * 10).Before(now) {
			log.Logger.Info(fmt.Sprintf("session[%d] heartbeat timeout", ses.ID()))
			self.Kick(ses.ID())
			return
		}

		// 踢掉验证超时
		v2, ok := ses.(comm.ContextSet).GetContext(AuthKey)
		if !ok {
			log.Logger.Error(fmt.Sprintf("session[%d] without autukey, kick it out", ses.ID()))
			self.Kick(ses.ID())
			return
		}
		// 验证过的跳过
		auth := v2.(bool)
		if auth {
			return
		}

		v3, ok := ses.(comm.ContextSet).GetContext(AuthTimeKey)
		if !ok {
			log.Logger.Error(fmt.Sprintf("session[%d] without authtime key, kick it out", ses.ID()))
			self.Kick(ses.ID())
			return
		}

		authTime := v3.(int64)
		if time.Unix(authTime, 0).Add(time.Second * 3).Before(now) {
			log.Logger.Info(fmt.Sprintf("session[%d] auth out of limit time, kick it out", ses.ID()))
			self.Kick(ses.ID())
			return
		}
	})
}

func NewFrontEventManager() *FrontEndManager {
	return &FrontEndManager{
		sessionMap: make(map[int64]comm.Session),
	}
}

func init() {
	frontEndMsgDispatcher = processor.NewMessageDispatcher()
	frontEndMsgDispatcher.SetDefaultCallback(ZoneDefaultHanlder)
	frontEndMsgDispatcher.RegisterMessage("VerifyReq", ZoneMsgVerify)

	FrontEndMgr = NewFrontEventManager()
}