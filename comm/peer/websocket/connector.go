package websocket

import (
	"fmt"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/gorilla/websocket"
)

const (
	maxReportFailTime = 3
)

type wsConnector  struct {
	peer.CoreSessionManager

	peer.CorePeerProperty
	peer.CoreContextSet
	peer.CoreRunningTag
	peer.CoreProcBundle

	ses *wsSession

	tryConnTimes int
	sesEndSignal sync.WaitGroup

	reconDur time.Duration
}

func (self *wsConnector) Start() {
	self.WaitStopFinish()

	if self.IsRunning() {
		return
	}

	go self.connect()
}

func (self *wsConnector) Stop() {
	if !self.IsRunning() {
		return
	}

	if self.IsStopping() {
		return
	}

	self.StartStopping()
	self.ses.Close()
	self.WaitStopFinish()
}

func (self *wsConnector) TypeName() string {
	return "ws.Connector"
}

func (self *wsConnector) SetReconnectDuration(t time.Duration) {
	self.reconDur = t
}

func (self *wsConnector) ReconnectDuration() time.Duration {
	return self.reconDur
}

func (self *wsConnector) Session() comm.Session {
	return self.ses
}

func (self *wsConnector) SetSessionManager(raw interface{}) {
	self.CoreSessionManager = raw.(peer.CoreSessionManager)
}

func (self *wsConnector) Port() int {
	if self.ses.conn == nil {
		return 0
	}

	return self.ses.conn.LocalAddr().(*net.TCPAddr).Port
}

func (self *wsConnector) connect() {
	self.SetRunning(true)

	for {
		self.tryConnTimes++

		dialer := websocket.Dialer{}
		dialer.Proxy = http.ProxyFromEnvironment
		dialer.HandshakeTimeout = time.Second * 20

		addr := self.Address()
		if !strings.HasPrefix(addr, "ws") && !strings.HasPrefix(addr, "wss")  {
			addr = "ws://" + addr
		}

		conn, _, err := dialer.Dial(addr, nil)
		self.ses.conn = conn
		if err != nil {
			if self.tryConnTimes <= maxReportFailTime {
				log.Logger.Error(fmt.Sprintf("#ws.Connector connect fail : %s", err))

				if self.tryConnTimes == maxReportFailTime {
					log.Logger.Error("#ws.Connector connect fail, but reach limit time, now mute")
				}
			}

			if self.ReconnectDuration() == 0 || !self.IsRunning() {
				self.ProcessEvent(&event.RecvMsgEvent{self.ses, &sysmsg.SessionConnectError{}})
				break
			}

			time.Sleep(self.ReconnectDuration())
			continue
		}

		self.tryConnTimes = 0
		self.sesEndSignal.Add(1)
		self.ses.Start()

		// 连接事件
		self.ProcessEvent(&event.RecvMsgEvent{self.ses, &sysmsg.SessionConnected{}})
		self.sesEndSignal.Wait()

		if self.ReconnectDuration() == 0 || self.IsRunning() {
			break
		}

		time.Sleep(self.ReconnectDuration())
	}

	self.SetRunning(false)
	self.EndStopping()
}

func init() {
	peer.RegPeerCreateor(func() comm.Peer {
		self := &wsConnector{}

		self.ses = newSession(nil, self, func() {
			self.sesEndSignal.Done()
		})

		return self
	})
}