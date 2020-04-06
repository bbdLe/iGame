package tcp

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"log"
	"net"
	"sync"
	"time"
)

type tcpConnector struct {
	peer.SessionManager

	peer.CorePeerProperty
	peer.CoreContextSet
	peer.CoreRunningTag
	peer.CoreProcBundle
	peer.CoreTcpSocketOption

	ses *tcpSession
	tryConnTimes int

	sesEndSignal sync.WaitGroup
	reconDuration time.Duration
}

const (
	reportConnectFailedLimitTimes = 3
)

func (self *tcpConnector) ReconnectDuration() time.Duration {
	return self.reconDuration
}

func (self *tcpConnector) Session() comm.Session {
	return self.ses
}

func (self *tcpConnector) SetReconnectDuration(v time.Duration) {
	self.reconDuration = v
}

func (self *tcpConnector) Port() int {
	conn := self.ses.Conn()

	return conn.LocalAddr().(*net.TCPAddr).Port
}

func (self *tcpConnector) SetSessionManager(raw interface{}) {
	self.SessionManager = raw.(peer.SessionManager)
}

func (self *tcpConnector) Start() {
	self.WaitStopFinish()

	if self.IsRunning() {
		return
	}

	go self.connect(self.Address())
}


func (self *tcpConnector) Stop() {
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

func (self *tcpConnector) TypeName() string {
	return "tcp.Connector"
}

func (self *tcpConnector) connect(address string) {
	self.SetRunning(true)

	for {
		self.tryConnTimes++

		conn, err := net.Dial("tcp", address)
		self.ses.SetConn(conn)

		if err != nil {
			if self.tryConnTimes <= reportConnectFailedLimitTimes {
				log.Printf("connect %v fail: %v", self.Address(), err)

				if self.tryConnTimes == reportConnectFailedLimitTimes {
					log.Printf("connect %v report limit, mute", self.Name())
				}
			}

			if self.ReconnectDuration() == 0 || self.IsStopping() {
				self.ProcessEvent(&event.RecvMsgEvent{
					Ses: self.ses,
					Msg: &sysmsg.SessionConnectError{},
				})
				break
			}

			time.Sleep(self.ReconnectDuration())
			continue
		}

		self.sesEndSignal.Add(1)

		self.ApplySocketOption(conn)
		self.ses.Start()
		self.tryConnTimes = 0
		self.ProcessEvent(&event.RecvMsgEvent{Ses : self.ses, Msg : &sysmsg.SessionConnected{}})
		self.sesEndSignal.Wait()

		self.ses.SetConn(nil)

		if self.IsStopping() || self.ReconnectDuration() == 0 {
			break
		}

		time.Sleep(self.ReconnectDuration())
		continue
	}

	self.SetRunning(false)
	self.EndStopping()
}

func (self *tcpConnector) IsReady() bool {
	return self.SessionCount() != 0
}

func init() {
	peer.RegPeerCreateor(func() comm.Peer {
		self := &tcpConnector{
			SessionManager: new(peer.CoreSessionManager),
		}

		self.ses = newSession(nil, self, func(){
			self.sesEndSignal.Done()
		})

		self.CoreTcpSocketOption.Init()

		return self
	})
}