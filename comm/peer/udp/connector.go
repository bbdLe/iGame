package udp

import (
	"fmt"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/sysmsg"
	"net"

	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
)

const (
	MaxUDPRecvBuffer = 2048
)

type udpConnector struct {
	peer.SessionManager
	peer.CorePeerProperty
	peer.CoreProcBundle
	peer.CoreContextSet
	peer.CoreRunningTag

	remoteAddr *net.UDPAddr
	ses *udpSession
}

func (self *udpConnector) Start() {
	var err error
	if self.remoteAddr, err = net.ResolveUDPAddr("udp", self.Address()); err != nil {
		log.Logger.Error(fmt.Sprintf("ResolveUDPAddr %v  fail, err : %v", self.Address(), err))
		return
	}

	go self.connect()
}

func (self *udpConnector) Stop() {
	self.SetRunning(false)

	if c := self.ses.Conn(); c != nil {
		c.Close()
	}
}

func (self *udpConnector) TypeName() string {
	return "udp.Connector"
}

func (self *udpConnector) Session() comm.Session {
	return self.ses
}

func (self *udpConnector) connect() {
	conn, err := net.DialUDP("udp", nil, self.remoteAddr)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("#udp.connect failed(%s) %v", self.Name(), err))
		return
	}

	self.ses.setConn(conn)
	self.ProcessEvent(&event.RecvMsgEvent{Ses: self.ses, Msg : &sysmsg.SessionConnected{}})

	recvBuff := make([]byte, MaxUDPRecvBuffer)
	self.SetRunning(true)

	for self.IsRunning() {
		n, _, err := conn.ReadFromUDP(recvBuff)
		if err != nil {
			break
		}

		if n > 0 {
			self.ses.Recv(recvBuff[:n])
		}
	}
}

func init() {
	peer.RegPeerCreateor(func() comm.Peer {
		p := &udpConnector{}
		p.ses = &udpSession{
			pInterface: p,
			CoreProcBundle : &p.CoreProcBundle,
		}

		return p
	})
}