package udp

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/log"
	"github.com/bbdLe/iGame/comm/peer"
)

type DataReader interface {
	ReadData() []byte
}

type DataWriter interface {
	WriteData(data []byte)
}

type udpSession struct {
	*peer.CoreProcBundle
	peer.CoreContextSet
	peer.CoreSessionIdentify

	pInterface comm.Peer

	pkg []byte

	remote		*net.UDPAddr
	conn		*net.UDPConn
	connGuard	sync.RWMutex
	timeOutTick time.Time
	key			*connTrackKey
}

func (self *udpSession) setConn(conn *net.UDPConn) {
	self.connGuard.Lock()
	defer self.connGuard.Unlock()
	self.conn = conn
}

func (self *udpSession) Conn() *net.UDPConn {
	self.connGuard.RLock()
	defer self.connGuard.RUnlock()

	return self.conn
}

func (self *udpSession) IsAlive() bool {
	return time.Now().Before(self.timeOutTick)
}

func (self *udpSession) ID() int64 {
	return 0
}

func (self *udpSession) LocalAddress() net.Addr {
	return self.conn.LocalAddr()
}

func (self *udpSession) Peer() comm.Peer {
	return self.pInterface
}

func (self *udpSession) Recv(data []byte) {
	self.pkg = data

	msg, err := self.ReadMessage(self)
	if msg != nil && err == nil {
		self.ProcessEvent(&event.RecvMsgEvent{self, msg})
	}
}

func (self *udpSession) Send(msg interface{}) {
	self.SendMessage(&event.SendMsgEvent{self, msg})
}

func (self *udpSession) Close() {
	return
}

func (self *udpSession) Raw() interface{} {
	return self
}

func (self *udpSession) WriteData(data []byte) {
	if self.conn == nil {
		return
	}

	if self.remote == nil {
		if _, err := self.conn.Write(data); err != nil {
			log.Logger.Error(fmt.Sprintf("udp write fail: %v", err))
		}
	} else {
		if _, err := self.conn.WriteToUDP(data, self.remote); err != nil {
			log.Logger.Error(fmt.Sprintf("udp write fail: %v", err))
		}
	}
}

func (self *udpSession) ReadData() []byte {
	return self.pkg
}