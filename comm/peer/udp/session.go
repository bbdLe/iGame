package udp

import (
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/event"
	"github.com/bbdLe/iGame/comm/peer"
	"net"
	"sync"
	"time"
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
	if msg != nil && err != nil {
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