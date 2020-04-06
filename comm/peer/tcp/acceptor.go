package tcp

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/peer"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"log"
	"net"
	"time"
)

type tcpAcceptor struct {
	peer.SessionManager
	peer.CorePeerProperty
	peer.CoreContextSet
	peer.CoreRunningTag
	peer.CoreProcBundle
	peer.CoreTcpSocketOption

	listener net.Listener
}

func (self *tcpAcceptor) Port() int {
	if self.listener == nil {
		return 0
	}

	return self.listener.Addr().(*net.TCPAddr).Port
}

func (self *tcpAcceptor) IsReady() bool {
	return self.IsRunning()
}

func (self *tcpAcceptor) Start() {
	self.WaitStopFinish()

	if self.IsRunning() {
		return
	}

	ln, err := net.Listen("tcp", self.Address())
	if err != nil {
		log.Println("#tcp.listen fail(%s): %v", self.Address(), err)
		self.SetRunning(false)
		return
	}

	self.listener = ln

	go self.accept()
}

func (self *tcpAcceptor) accept() {
	self.SetRunning(true)

	for {
		conn, err := self.listener.Accept()

		if self.IsStopping() {
			break
		}

		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				time.Sleep(time.Millisecond)
				continue
			}

			log.Printf("#tcp.accept failed(%s) %v", self.Name, err.Error())
			break
		} else {
			go self.onNewSession(conn)
		}
	}

	self.SetRunning(false)
	self.EndStopping()
}

func (self *tcpAcceptor) Stop() {
	if !self.IsRunning() {
		return
	}

	if self.IsStopping() {
		return
	}

	self.StartStopping()
	self.listener.Close()
	self.CloseAllSession()
	self.WaitStopFinish()
}

func (self *tcpAcceptor) TypeName() string {
	return "tcp.Acceptor"
}

func (self *tcpAcceptor) onNewSession(conn net.Conn) {
	self.ApplySocketOption(conn)

	ses := newSession(conn, self, nil)
	ses.Start()

	self.ProcessEvent(&event.RecvMsgEvent{
		Ses: ses,
		Msg: &sysmsg.SessionAccepted{},
	})
}

func init() {
	peer.RegPeerCreateor(func() comm.Peer {
		p := &tcpAcceptor{
			SessionManager : new(peer.CoreSessionManager),
		}
		p.CoreTcpSocketOption.Init()
		return p
	})
}