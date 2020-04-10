package udp

import (
	"fmt"
	"net"
	"time"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
)

type udpAcceptor struct {
	peer.CoreSessionManager
	peer.CorePeerProperty
	peer.CoreContextSet
	peer.CoreRunningTag
	peer.CoreProcBundle

	conn *net.UDPConn

	sesTimeout   	 time.Duration
	sesCleanTimeout	 time.Duration
	sesCleanLastTime time.Time

	sesByConnTrack map[connTrackKey]*udpSession
}

func (self *udpAcceptor) IsReady() bool {
	return self.IsRunning()
}

func (self *udpAcceptor) Port() int {
	if self.conn == nil {
		return 0
	}

	return self.conn.LocalAddr().(*net.UDPAddr).Port
}

func (self *udpAcceptor) Start() {
	addr, err := net.ResolveUDPAddr("udp", self.Address())
	if err != nil {
		log.Logger.Error(fmt.Sprintf("ResolveUDPAddr: %v", err))
		return
	}

	self.conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("ListenUDP: %v", err))
		return
	}

	go self.accept()
	return
}

func (self *udpAcceptor) Stop() {
	if self.conn != nil {
		self.conn.Close()
	}

	self.SetRunning(false)
}

func (self *udpAcceptor) accept() {
	self.SetRunning(true)

	recvBuff := make([]byte, MaxUDPRecvBuffer)

	for {
		n, remoteAddr, err := self.conn.ReadFromUDP(recvBuff)
		if err != nil {
			break
		}

		self.CheckTimeoutSession()
		if n > 0 {
			ses := self.getSession(remoteAddr)
			ses.Recv(recvBuff[:n])
		}
	}
}

func (self *udpAcceptor) TypeName() string {
	return "udp.Acceptor"
}

func (self *udpAcceptor) CheckTimeoutSession() {
	now := time.Now()

	if now.After(self.sesCleanLastTime.Add(self.sesCleanTimeout)) {
		sesToDeleteList := make([]*udpSession, 0, 10)
		for _, ses := range self.sesByConnTrack {
			if !ses.IsAlive() {
				sesToDeleteList = append(sesToDeleteList, ses)
			}
		}

		for _, ses := range sesToDeleteList {
			delete(self.sesByConnTrack, *ses.key)
		}

		self.sesCleanLastTime = now
	}
}

func (self *udpAcceptor) getSession(addr *net.UDPAddr) *udpSession {
	key := newConnTrackKey(addr)
	ses, ok := self.sesByConnTrack[*key]
	if !ok {
		ses = &udpSession{}
		ses.conn = self.conn
		ses.remote = addr
		ses.pInterface = self
		ses.CoreProcBundle = &self.CoreProcBundle
		ses.key = key
		self.sesByConnTrack[*key] = ses
	}

	ses.timeOutTick = time.Now().Add(self.sesTimeout)
	return ses
}

func (self *udpAcceptor) SetSessionTTL(dur time.Duration) {
	self.sesTimeout = dur
}

func (self *udpAcceptor) SetSessionCleanTimeout(dur time.Duration) {
	self.sesCleanTimeout = dur
}

func init() {
	peer.RegPeerCreateor(func() comm.Peer {
		p := &udpAcceptor{
			sesTimeout: time.Minute,
			sesCleanTimeout: time.Minute,
			sesCleanLastTime: time.Now(),
			sesByConnTrack: make(map[connTrackKey]*udpSession),
		}
		return p
	})
}